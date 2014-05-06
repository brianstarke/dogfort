package hub

import (
	"log"
	"net/http"

	"github.com/brianstarke/dogfort/domain"
	"github.com/gorilla/websocket"
)

type Connection struct {
	ws   *websocket.Conn
	send chan map[string]interface{}
}

func (c *Connection) Reader() {
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			break
		}
		log.Print(message)
	}
	c.ws.Close()
}

func (c *Connection) Writer() {
	for message := range c.send {
		err := c.ws.WriteJSON(message)

		if err != nil {
			log.Printf("Error publishing message: %s", err.Error())
			break
		} else {
			log.Printf("Published message: ", message)
		}

	}
	c.ws.Close()
}

func WsHandler(w http.ResponseWriter, u domain.UserUid, r *http.Request) {
	ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		return
	}
	c := &Connection{send: make(chan map[string]interface{}), ws: ws}
	H.Register <- struct {
		domain.UserUid
		*Connection
	}{u, c}
	defer func() { H.Unregister <- u }()
	go c.Writer()
	c.Reader()
}
