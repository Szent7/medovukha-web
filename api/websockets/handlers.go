package websockets

import (
	"context"
	"log"
	"net/http"

	"github.com/Szent7/medovukha-web/api/rest/v1/types"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// !Dev Upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func WsContainerEventsHandler(hub *Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError,
				types.NewFailed[types.BaseMessage](types.APIError{Code: types.ErrWeb, Message: err.Error()}))
			log.Printf("WsHandler error: %s\n", err.Error())
			return
		}

		g := hub.Register(containerEventStreamGroup, conn, hub.ContainerEventStreamer)
		go handleConn(g, conn, hub)
	}
}

func WsImageEventsHandler(hub *Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError,
				types.NewFailed[types.BaseMessage](types.APIError{Code: types.ErrWeb, Message: err.Error()}))
			log.Printf("WsHandler error: %s\n", err.Error())
			return
		}

		g := hub.Register(imageEventStreamGroup, conn, hub.ImageEventStreamer)
		go handleConn(g, conn, hub)
	}
}

func WsNetworkEventsHandler(hub *Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError,
				types.NewFailed[types.BaseMessage](types.APIError{Code: types.ErrWeb, Message: err.Error()}))
			log.Printf("WsHandler error: %s\n", err.Error())
			return
		}

		g := hub.Register(networkEventStreamGroup, conn, hub.NetworkEventStreamer)
		go handleConn(g, conn, hub)
	}
}

func WsVolumeEventsHandler(hub *Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError,
				types.NewFailed[types.BaseMessage](types.APIError{Code: types.ErrWeb, Message: err.Error()}))
			log.Printf("WsHandler error: %s\n", err.Error())
			return
		}

		g := hub.Register(volumeEventStreamGroup, conn, hub.VolumeEventStreamer)
		go handleConn(g, conn, hub)
	}
}

func WsBuildLogsHandler(hub *Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		buildID := c.Query("buildID")
		if buildID == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "missing buildID"})
			return
		}

		groupName := buildLogStreamGroup + buildID

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError,
				types.NewFailed[types.BaseMessage](types.APIError{Code: types.ErrWeb, Message: err.Error()}))
			log.Printf("WsHandler error: %s\n", err.Error())
			return
		}

		streamFn := func(ctx context.Context, ch chan<- []byte, buildID string) error {
			return hub.BuildLogStreamer(ctx, ch, buildID)
		}

		g := hub.Register(groupName, conn, streamFn)
		go handleConn(g, conn, hub)
	}
}

func handleConn(g *Group, conn *websocket.Conn, hub *Hub) {
	defer func() {
		hub.Unregister(g.getName(), conn)
		conn.Close()
	}()

	for {
		if _, _, err := conn.ReadMessage(); err != nil {
			return
		}
	}
}
