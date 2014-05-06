package routes

import (
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
		r.JSON(400, err.Error())
	} else {
		m, _ := domain.MessageDomain.MessageById(*id)
		hub.H.MessagePublish(message.ChannelId, m)

		r.JSON(200, map[string]interface{}{"message": id})
	}
}

func MessagesByChannel(userUid domain.UserUid, params martini.Params, r render.Render) {
	ok, err := domain.ChannelDomain.UserInChannel(&userUid, params["id"])

	if err != nil {
		r.JSON(400, err.Error())
	}

	if ok {
		messages, err := domain.MessageDomain.MessagesByChannel(&userUid, params["id"])

		if err != nil {
			r.JSON(400, err.Error())
		} else {
			r.JSON(200, messages)
		}
	} else {
		r.JSON(200, "User is not a member of this channel")
	}
}
