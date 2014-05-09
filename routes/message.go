package routes

import (
	"log"

	"github.com/brianstarke/dogfort/domain"
	"github.com/brianstarke/dogfort/hub"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func CreateMessage(message domain.Message, userUid domain.UserUid, r render.Render) {
	message.UserId = userUid
	id, err := domain.MessageDomain.CreateMessage(&message)

	message.IsHtml = false
	message.IsAdminMsg = false

	if err != nil {
		log.Printf("Error creating message: %s", err.Error())

		r.JSON(400, err.Error())
		return
	} else {
		m, err := domain.MessageDomain.MessageById(*id)

		if err != nil {
			log.Printf("Error getting message for publish: %s", err.Error())
			return
		}

		hub.H.MessagePublish(message.ChannelId, m)

		r.JSON(200, map[string]interface{}{"message": id})
		return
	}
}

func MessagesByChannel(userUid domain.UserUid, params martini.Params, r render.Render) {
	ok, err := domain.ChannelDomain.UserInChannel(&userUid, params["id"])

	if err != nil {
		r.JSON(400, err.Error())
	}

	if err != nil {
		r.JSON(400, err.Error())
	}

	if ok {
		messages, err := domain.MessageDomain.MessagesByChannel(&userUid, params["id"], params["before"], params["num"])

		if err != nil {
			r.JSON(400, err.Error())
		} else {
			r.JSON(200, messages)
		}
	} else {
		r.JSON(200, "User is not a member of this channel")
	}
}
