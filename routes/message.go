package routes

import (
	"github.com/brianstarke/dogfort/domain"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func CreateMessage(messageDomain *domain.MessageDomain, message domain.Message, userUid domain.UserUid, r render.Render) {
	message.UserId = userUid
	id, err := messageDomain.CreateMessage(&message)

	if err != nil {
		r.JSON(400, err.Error())
	} else {
		r.JSON(200, map[string]interface{}{"message": id})
	}
}

func MessagesByChannel(md *domain.MessageDomain, cd *domain.ChannelDomain, userUid domain.UserUid, params martini.Params, r render.Render) {
	ok, err := cd.UserInChannel(&userUid, params["channelId"])

	if err != nil {
		r.JSON(400, err.Error())
	}

	if ok {
		messages, err := md.MessagesByChannel(&userUid, params["channelId"])

		if err != nil {
			r.JSON(400, err.Error())
		} else {
			r.JSON(200, messages)
		}
	} else {
		r.JSON(200, "User is not a member of this channel")
	}
}
