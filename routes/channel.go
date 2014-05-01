package routes

import (
	"net/http"

	"github.com/brianstarke/dogfort/domain"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func ChannelCreate(userUid domain.UserUid, channel domain.Channel, req *http.Request, r render.Render) {
	id, err := domain.ChannelDomain.CreateChannel(&channel, &userUid)

	if err != nil {
		r.JSON(400, err.Error())
	} else {
		r.JSON(200, map[string]interface{}{"id": id})
	}

	return
}

func ChannelList(r render.Render) {
	c, err := domain.ChannelDomain.ListChannels()

	if err != nil {
		r.JSON(400, err.Error())
	} else {
		r.JSON(200, map[string]interface{}{"channels": c})
	}
}

func ChannelsByUser(userUid domain.UserUid, req *http.Request, r render.Render) {
	c, err := domain.ChannelDomain.GetUserChannels(&userUid)

	if err != nil {
		r.JSON(400, err.Error())
	} else {
		r.JSON(200, map[string]interface{}{"channels": c})
	}

	return
}

func ChannelJoin(userUid domain.UserUid, params martini.Params, r render.Render) {
	err := domain.ChannelDomain.SubscribeToChannel(&userUid, params["id"])

	if err != nil {
		r.JSON(400, err.Error())
	} else {
		r.JSON(200, "ok")
	}
}

func ChannelLeave(userUid domain.UserUid, params martini.Params, r render.Render) {
	err := domain.ChannelDomain.UnsubscribeFromChannel(&userUid, params["id"])

	if err != nil {
		r.JSON(400, err.Error())
	} else {
		r.JSON(200, "ok")
	}
}
