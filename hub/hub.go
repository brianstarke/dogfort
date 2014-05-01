package hub

import (
	"log"

	"github.com/brianstarke/dogfort/domain"
)

type Hub struct {
	Connections map[domain.UserUid]*Connection
	Register    chan struct {
		domain.UserUid
		*Connection
	}
	Unregister chan domain.UserUid
}

var H = Hub{
	Connections: make(map[domain.UserUid]*Connection),
	Register: make(chan struct {
		domain.UserUid
		*Connection
	}),
	Unregister: make(chan domain.UserUid),
}

func (h *Hub) Run() {
	for {
		select {

		case c := <-h.Register:
			h.Connections[c.UserUid] = c.Connection
			log.Printf("%s is active, %d currently active", c.UserUid, len(h.Connections))

		case c := <-h.Unregister:
			close(h.Connections[c].send)
			delete(h.Connections, c)

			log.Printf("connection unregistered, %d currently active", len(h.Connections))

		}
	}
}

func (h *Hub) Publish(topic string, message *string) {
	for _, v := range h.Connections {
		v.send <- map[string]interface{}{topic: message}
	}
}

func (h *Hub) MessagePublish(channelId string, message interface{}) {
	for k, v := range h.Connections {
		sendIt, err := domain.ChannelDomain.UserInChannel(&k, channelId)

		if err != nil {
			log.Println(err) // TODO handle this
		}

		if sendIt {
			v.send <- map[string]interface{}{channelId: message}
		}
	}
}
