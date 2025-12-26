package websockets

import (
	"context"
	"fmt"
	"log"
	"net/http"

	dockerpb "github.com/Szent7/medovukha-web/api/docker/v1"
	"github.com/Szent7/medovukha-web/api/rest/v1"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/encoding/protojson"
)

// !Dev Upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type wsHub struct {
	conns      map[*websocket.Conn]bool
	broadcast  chan []byte
	register   chan *websocket.Conn
	unregister chan *websocket.Conn

	// for containerEventStreamer
	workCtx    context.Context
	workCancel context.CancelFunc
	api        *rest.API
}

func (h *wsHub) Run() {
	for {
		select {
		case conn := <-h.register:
			{
				h.conns[conn] = true
				if len(h.conns) == 1 {
					h.workCtx, h.workCancel = context.WithCancel(context.Background())
					go h.containerEventStreamer()
				}
			}
		case conn := <-h.unregister:
			{
				if _, ok := h.conns[conn]; ok {
					delete(h.conns, conn)
					conn.Close()
				}
				if len(h.conns) == 0 {
					if h.workCancel != nil {
						h.workCancel()
					}
				}
			}
		case msg := <-h.broadcast:
			{
				for conn := range h.conns {
					if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
						conn.Close()
						delete(h.conns, conn)
					}
				}
			}
		}
	}
}

func (h *wsHub) containerEventStreamer() {
	eventCh := make(chan *dockerpb.ContainerState)
	errCh := make(chan error)
	go h.api.GetContainerState(h.workCtx, eventCh, errCh)

	fmt.Println("Start containerEventStreamer")

	for {
		select {
		case <-h.workCtx.Done():
			{
				fmt.Println("Stop containerEventStreamer")
				return
			}
		case event := <-eventCh:
			{
				jsonBytes, err := protojson.Marshal(event)
				if err != nil {
					fmt.Printf("protobuf marshall error: %s\n", err.Error())
					continue
				}

				h.broadcast <- jsonBytes
			}
		case err := <-errCh:
			{
				if err != nil {
					log.Printf("containerEventStreamer error: %s\n", err.Error())
					h.workCancel()
				}
			}
		}
	}
}

func NewHub(api *rest.API) *wsHub {
	return &wsHub{
		conns:      make(map[*websocket.Conn]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *websocket.Conn),
		unregister: make(chan *websocket.Conn),
		api:        api,
	}
}

func WsHandler(h *wsHub) gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Printf("WsHandler error: %s\n", err.Error())
			return
		}

		h.register <- conn
		fmt.Println("New WS client")

		for {
			if _, _, err := conn.ReadMessage(); err != nil {
				log.Printf("Read error (closing): %v", err)
				break
			}
		}

		conn.Close()
		h.unregister <- conn
		fmt.Println("Close WS client")
	}
}
