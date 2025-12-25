package main

import (
	"log"
	"medovukha/api/rest/middlewares"
	v1 "medovukha/api/rest/v1"
	"net"
	"net/rpc"

	"github.com/gin-gonic/gin"
)

func main() {
	const socketPath = "/tmp/medovukha-core.sock"
	const serverName = "MedovukhaCore"
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		log.Fatalf("cannot connect to medovukha-core: %s", err.Error())
	}

	rpcClient := rpc.NewClient(conn)
	defer rpcClient.Close()

	api := v1.NewAPI(rpcClient, serverName)

	log.Printf("%s connected to %s\n", serverName, socketPath)

	router := gin.Default()

	// wsHub := websockets.NewHub()
	// go wsHub.Run()

	router.Static("/_app/immutable/", "./build/_app/immutable/")
	router.NoRoute(func(c *gin.Context) {
		c.File("./build/index.html")
	})

	//! dev headers
	router.Use(middlewares.CORSMiddleware())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// ws := router.Group("/ws")
	// {
	// 	ws.GET("/containerEvents", websockets.WsHandler(wsHub))
	// }

	rest := router.Group("/rest")
	{
		v1 := rest.Group("/v1")
		{
			//Containers
			v1.GET("/getContainerList", api.GetContainerList)
			v1.POST("/pauseContainerByid", api.PauseContainerByID)
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
