package hub

import "log"

type Hub struct {
	Connections map[*Connection]bool // registered connections
	Broadcast   chan []byte          // inbound messages from connections
	Register    chan *Connection     // register requests
	Unregister  chan *Connection     // unregister requests
}

var H = Hub{
	Broadcast:   make(chan []byte),
	Register:    make(chan *Connection),
	Unregister:  make(chan *Connection),
	Connections: make(map[*Connection]bool),
}

func (h *Hub) Run() {
	for {
		select {
		case c := <-h.Register:
			h.Connections[c] = true
			log.Printf("new connection registered, %d currently active", len(h.Connections))
		case c := <-h.Unregister:
			delete(h.Connections, c)
			close(c.send)
			log.Printf("connection unregistered, %d currently active", len(h.Connections))
		case m := <-h.Broadcast:
			for c := range h.Connections {
				select {
				case c.send <- m:
				default:
					delete(h.Connections, c)
					close(c.send)
					go c.ws.Close()
				}
			}
		}
	}
}
