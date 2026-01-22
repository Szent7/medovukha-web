package websockets

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/Szent7/medovukha-web/api/rest/v1"
	"github.com/gorilla/websocket"
)

const (
	containerEventStreamGroup = "container-events"
	imageEventStreamGroup     = "image-events"
	networkEventStreamGroup   = "network-events"
	volumeEventStreamGroup    = "volume-events"
	buildLogStreamGroup       = "build-container-"
)

type Hub struct {
	api    *rest.API
	mu     sync.RWMutex
	groups map[string]*Group
}

func NewHub(api *rest.API) *Hub {
	return &Hub{
		api:    api,
		groups: make(map[string]*Group)}
}

func (h *Hub) onGroupDone(groupName string) {
	h.mu.Lock()
	g, ok := h.groups[groupName]
	if ok {
		delete(h.groups, groupName)
	}
	h.mu.Unlock()

	if ok {
		g.stop()
	}
}

func (h *Hub) Register(groupName string, conn *websocket.Conn, streamFn StreamFunc) *Group {
	h.mu.Lock()
	g, ok := h.groups[groupName]
	if !ok {
		g = newGroup(groupName, h.onGroupDone)
		h.groups[groupName] = g
		g.start(streamFn)
	}
	h.mu.Unlock()

	g.addConn(conn)
	return g
}

func (h *Hub) Unregister(groupName string, conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()

	g, ok := h.groups[groupName]
	if !ok {
		return
	}
	conn.Close()
	g.removeConn(conn)
	if g.isEmpty() {
		g.stop()
		delete(h.groups, groupName)
	}
}

type StreamFunc func(ctx context.Context, broadcast chan<- []byte, buildID string) error

type Group struct {
	name         string
	conns        map[*websocket.Conn]struct{}
	mu           sync.RWMutex
	broadcast    chan []byte
	streamCtx    context.Context
	streamCancel context.CancelFunc

	stopOnce sync.Once
	onDone   func(groupName string)
}

func newGroup(name string, onDone func(string)) *Group {
	return &Group{
		name:      name,
		conns:     make(map[*websocket.Conn]struct{}),
		broadcast: make(chan []byte, 128),
		onDone:    onDone,
	}
}

func (g *Group) getName() string {
	return g.name
}

func (g *Group) addConn(conn *websocket.Conn) {
	g.mu.Lock()
	defer g.mu.Unlock()
	//g.conns = make(map[*websocket.Conn]struct{})
	g.conns[conn] = struct{}{}
}

func (g *Group) removeConn(c *websocket.Conn) {
	g.mu.Lock()
	defer g.mu.Unlock()
	delete(g.conns, c)
}

func (g *Group) isEmpty() bool {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return len(g.conns) == 0
}

func (g *Group) start(streamFn StreamFunc) {
	g.streamCtx, g.streamCancel = context.WithCancel(context.Background())
	//go streamFn(g.streamCtx, g.broadcast, g.name)
	go g.streamLoop(streamFn)
	go g.relay()
}

func (g *Group) stop() {
	g.stopOnce.Do(func() {
		g.streamCancel()

		g.mu.Lock()
		for c := range g.conns {
			c.Close()
			delete(g.conns, c)
		}
		g.mu.Unlock()

		close(g.broadcast)
	})
}

func (g *Group) relay() {
	for {
		select {
		case <-g.streamCtx.Done():
			{
				return
			}
		case msg, ok := <-g.broadcast:
			{
				if !ok {
					return
				}

				g.mu.Lock()
				for c := range g.conns {
					c.SetWriteDeadline(time.Now().Add(5 * time.Second))
					if err := c.WriteMessage(websocket.TextMessage, msg); err != nil {
						log.Printf("hub: write error %s", err.Error())
						c.Close()
						g.removeConn(c)
					}
				}
				g.mu.Unlock()
			}
		}
	}
}

func (g *Group) streamLoop(streamFn StreamFunc) {
	err := streamFn(g.streamCtx, g.broadcast, g.name)
	if err != nil {
		log.Printf("[group=%s] streamer finished with error: %s", g.name, err.Error())
	} else {
		log.Printf("[group=%s] streamer finished without error\n", g.name)
	}

	g.onDone(g.name)
}
