package websocket

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	// Restrict access to clients
	CheckOrigin: func(r *http.Request) bool { return true },
}

// Hub Struct
type Hub struct {
	clients    []*Client
	register   chan *Client
	unregister chan *Client
	mutex      *sync.Mutex
}

// Constructor of Hub
func NewHub() *Hub {
	return &Hub{
		clients:    make([]*Client, 0),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		mutex:      &sync.Mutex{},
	}
}

// Route of Websockets
func (hub *Hub) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade connection
	socket, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

	// Create New Client
	client := NewClient(hub, socket)

	hub.register <- client

	// GoRoutine wich send messages to registered clients
	go client.Write()
}

//
func (hub *Hub) onConnect(client *Client) {
	// Console message
	log.Println("Client connected", client.socket.RemoteAddr())

	hub.mutex.Lock()
	defer hub.mutex.Unlock()

	// Assign id for client
	client.id = client.socket.RemoteAddr().String()

	// Add client to hub
	hub.clients = append(hub.clients, client)
}

func (hub *Hub) Run() {
	for {
		select {
		case client := <-hub.register:
			hub.onConnect(client)
		case client := <-hub.unregister:
			hub.onDisconnect(client)
		}
	}
}
