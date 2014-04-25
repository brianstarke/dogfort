package hub

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type Connection struct {
	ws   *websocket.Conn
	send chan []byte
}

func (c *Connection) Reader() {
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			break
		}
		H.Broadcast <- message
	}
	c.ws.Close()
}

func (c *Connection) Writer() {
	for message := range c.send {
		err := c.ws.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			break
		}
	}
	c.ws.Close()
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		return
	}
	c := &Connection{send: make(chan []byte, 256), ws: ws}
	H.Register <- c
	defer func() { H.Unregister <- c }()
	go c.Writer()
	c.Reader()
}
