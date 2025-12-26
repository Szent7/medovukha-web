package main

import (
	"context"
	"log"
	"net"
	"time"

	dockerpb "github.com/Szent7/medovukha-web/api/docker/v1"
	"github.com/Szent7/medovukha-web/api/rest/middlewares"
	rest "github.com/Szent7/medovukha-web/api/rest/v1"
	"github.com/Szent7/medovukha-web/api/websockets"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/gin-gonic/gin"
)

func main() {
	const socketPath = "unix:/tmp/medovukha-core.sock"

	//ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	//defer cancel()

	conn, err := grpc.NewClient(
		socketPath,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(func(ctx context.Context, addr string) (net.Conn, error) {
			return net.DialTimeout("unix", addr[len("unix:"):], 5*time.Second)
		}),
	)
	if err != nil {
		log.Fatalf("failed to create grpc client: %s", err.Error())
	}
	defer conn.Close()

	rpcClient := dockerpb.NewDockerServiceClient(conn)

	api := rest.NewAPI(rpcClient)

	log.Printf("connected to %s\n", socketPath)

	router := gin.Default()

	wsHub := websockets.NewHub(api)
	go wsHub.Run()

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
		ws.GET("/containerEvents", websockets.WsHandler(wsHub))
	}

	rest := router.Group("/rest")
	{
		v1 := rest.Group("/v1")
		{
			//Containers
			v1.GET("/getContainerList", api.GetContainerList)
			v1.POST("/pauseContainerById", api.PauseContainerByID)
			v1.POST("/unpauseContainerById", api.UnpauseContainerByID)
			v1.POST("/killContainerById", api.KillContainerByID)
			v1.POST("/startContainerById", api.StartContainerByID)
			v1.POST("/stopContainerById", api.StopContainerByID)
			v1.POST("/restartContainerById", api.RestartContainerByID)
			v1.POST("/removeContainerById", api.RemoveContainerByID)
			//Images
			v1.GET("/getImageList", api.GetImageList)
			//Networks
			v1.GET("/getNetworkList", api.GetNetworkList)
			//Volumes
			v1.GET("/getVolumeList", api.GetVolumeList)
			//Deploy
			v1.POST("/createFromGit", api.CreateFromGit)
		}
	}

	router.Run("0.0.0.0:10015")
}
