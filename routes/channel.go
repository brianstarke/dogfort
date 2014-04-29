package routes

import (
	"net/http"

	"github.com/brianstarke/dogfort/domain"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func ChannelCreate(channelDomain *domain.ChannelDomain, userUid domain.UserUid, channel domain.Channel, req *http.Request, r render.Render) {
	id, err := channelDomain.CreateChannel(&channel, &userUid)

	if err != nil {
		r.JSON(400, err.Error())
	} else {
		r.JSON(200, map[string]interface{}{"id": id})
	}

	return
}

func ChannelList(channelDomain *domain.ChannelDomain, r render.Render) {
	c, err := channelDomain.ListChannels()

	if err != nil {
		r.JSON(400, err.Error())
	} else {
		r.JSON(200, map[string]interface{}{"channels": c})
	}
}

func ChannelsByUser(channelDomain *domain.ChannelDomain, userUid domain.UserUid, req *http.Request, r render.Render) {
	c, err := channelDomain.GetUserChannels(&userUid)

	if err != nil {
		r.JSON(400, err.Error())
	} else {
		r.JSON(200, map[string]interface{}{"channels": c})
	}

	return
}

func ChannelJoin(cd *domain.ChannelDomain, userUid domain.UserUid, params martini.Params, r render.Render) {
	err := cd.SubscribeToChannel(&userUid, params["id"])

	if err != nil {
		r.JSON(400, err.Error())
	} else {
		r.JSON(200, "ok")
	}
}

func ChannelLeave(cd *domain.ChannelDomain, userUid domain.UserUid, params martini.Params, r render.Render) {
	err := cd.UnsubscribeFromChannel(&userUid, params["id"])

	if err != nil {
		r.JSON(400, err.Error())
	} else {
		r.JSON(200, "ok")
	}
}
