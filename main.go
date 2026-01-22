package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	dockerpb "github.com/Szent7/medovukha-web/api/docker/v1"
	"github.com/Szent7/medovukha-web/api/rest/middlewares"
	rest "github.com/Szent7/medovukha-web/api/rest/v1"
	"github.com/Szent7/medovukha-web/api/websockets"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/gin-gonic/gin"
)

const retryPolicy = `{
		"methodConfig": [{
		  "name": [{"service": "docker.v1.DockerService"}],
		  "retryPolicy": {
			  "MaxAttempts": 4,
			  "InitialBackoff": ".01s",
			  "MaxBackoff": ".01s",
			  "BackoffMultiplier": 1.0,
			  "RetryableStatusCodes": [ "UNAVAILABLE" ]
		  }
		}]}`

func main() {
	const socketPath = "unix:/tmp/medovukha-core.sock"

	grpcClient, err := grpc.NewClient(
		socketPath,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(func(ctx context.Context, addr string) (net.Conn, error) {
			return net.DialTimeout("unix", addr[len("unix:"):], 5*time.Second)
		}),
		grpc.WithDefaultServiceConfig(retryPolicy),
	)
	if err != nil {
		log.Fatalf("failed to create grpc client: %s", err.Error())
	}

	rpcClient := dockerpb.NewDockerServiceClient(grpcClient)

	api := rest.NewAPI(rpcClient)

	log.Printf("connected to %s\n", socketPath)

	router := gin.Default()

	wsHub := websockets.NewHub(api)

	router.Static("/_app/immutable/", "./build/_app/immutable/")
	router.NoRoute(func(c *gin.Context) {
		c.File("./build/index.html")
	})

	//! dev headers
	router.Use(middlewares.CORSMiddleware(), middlewares.TimeoutMiddleware(2*time.Minute))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	ws := router.Group("/ws")
	{
		ws.GET("/containerEvents", websockets.WsContainerEventsHandler(wsHub))
		ws.GET("/imageEvents", websockets.WsImageEventsHandler(wsHub))
		ws.GET("/networkEvents", websockets.WsNetworkEventsHandler(wsHub))
		ws.GET("/volumeEvents", websockets.WsVolumeEventsHandler(wsHub))
		ws.GET("/buildLogs", websockets.WsBuildLogsHandler(wsHub))
	}

	rest := router.Group("/rest")
	{
		v1 := rest.Group("/v1")
		{
			//Containers
			v1.POST("/getContainerById", api.GetContainerByID)
			v1.GET("/getContainerList", api.GetContainerList)
			v1.POST("/pauseContainerById", api.PauseContainerByID)
			v1.POST("/unpauseContainerById", api.UnpauseContainerByID)
			v1.POST("/killContainerById", api.KillContainerByID)
			v1.POST("/startContainerById", api.StartContainerByID)
			v1.POST("/stopContainerById", api.StopContainerByID)
			v1.POST("/restartContainerById", api.RestartContainerByID)
			v1.POST("/removeContainerById", api.RemoveContainerByID)
			//Images
			v1.POST("/getImageById", api.GetImageByID)
			v1.GET("/getImageList", api.GetImageList)
			v1.POST("/removeImage", api.RemoveImage)
			//Networks
			v1.POST("/getNetworkById", api.GetNetworkByID)
			v1.GET("/getNetworkList", api.GetNetworkList)
			v1.POST("/removeNetwork", api.RemoveNetwork)
			//Volumes
			v1.POST("/getVolumeById", api.GetVolumeByID)
			v1.GET("/getVolumeList", api.GetVolumeList)
			v1.POST("/removeVolume", api.RemoveVolume)
			//Deploy
			v1.POST("/createFromGit", api.CreateFromGit)
		}
	}

	srv := &http.Server{
		Addr:    ":10015",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s", err.Error())
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigCh
	log.Printf("received signal %s, shutting down...\n", sig)

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if grpcClient != nil {
		if err := grpcClient.Close(); err != nil {
			log.Printf("error while closing gRPC service (dockerService): %s\n", err.Error())
		}
	}

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("error while closing gRPC client (dockerService): %s\n", err.Error())
	}

	<-shutdownCtx.Done()
	log.Println("Shutdown complete")
}
