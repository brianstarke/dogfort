package routes

import (
	"net/http"

	"github.com/brianstarke/dogfort/domain"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func CreateChannel(channelDomain *domain.ChannelDomain, userUid domain.UserUid, channel domain.Channel, req *http.Request, r render.Render) {
	id, err := channelDomain.CreateChannel(&channel, &userUid)

	if err != nil {
		r.JSON(400, err.Error())
	} else {
		r.JSON(200, map[string]interface{}{"id": id})
	}

	return
}

func ListChannels(channelDomain *domain.ChannelDomain, r render.Render) {
	c, err := channelDomain.ListChannels()

	if err != nil {
		r.JSON(400, err.Error())
	} else {
		r.JSON(200, map[string]interface{}{"channels": c})
	}
}

func GetUserChannels(channelDomain *domain.ChannelDomain, userUid domain.UserUid, req *http.Request, r render.Render) {
	c, err := channelDomain.GetUserChannels(&userUid)

	if err != nil {
		r.JSON(400, err.Error())
	} else {
		r.JSON(200, map[string]interface{}{"channels": c})
	}

	return
}

func JoinChannel(cd *domain.ChannelDomain, userUid domain.UserUid, params martini.Params, r render.Render) {
	err := cd.SubscribeToChannel(&userUid, params["channelId"])

	if err != nil {
		r.JSON(400, err.Error())
	} else {
		r.JSON(200, "ok")
	}
}

func LeaveChannel(cd *domain.ChannelDomain, userUid domain.UserUid, params martini.Params, r render.Render) {
	err := cd.UnsubscribeFromChannel(&userUid, params["channelId"])

	if err != nil {
		r.JSON(400, err.Error())
	} else {
		r.JSON(200, "ok")
	}
}
