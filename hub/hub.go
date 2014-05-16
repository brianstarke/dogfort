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
			h.Broadcast("user_join", string(c.UserUid))

		case c := <-h.Unregister:
			close(h.Connections[c].send)
			delete(h.Connections, c)

			h.Broadcast("user_leave", string(c))
			log.Printf("connection unregistered, %d currently active", len(h.Connections))
		}
	}
}

/*
Returns a list of users that are currently online
*/
func (h *Hub) UsersOnline() []domain.UserUid {
	users := make([]domain.UserUid, len(h.Connections))

	for k, _ := range h.Connections {
		users = append(users, k)
	}

	return users
}

func (h *Hub) Broadcast(topic string, message string) {
	for _, v := range h.Connections {
		v.send <- map[string]interface{}{topic: message}
	}
}

func (h *Hub) MessagePublish(message *domain.Message) {
	for k, v := range h.Connections {
		sendIt, err := domain.ChannelDomain.UserInChannel(&k, message.ChannelId)

		if err != nil {
			log.Printf("Error publishing message: %s", err.Error())
		}

		if sendIt {
			v.send <- map[string]interface{}{"message": message}
		}
	}
}
