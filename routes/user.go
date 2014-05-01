package routes

import (
	"net/http"

	"github.com/brianstarke/dogfort/domain"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func UserCreate(newUser domain.NewUser, req *http.Request, r render.Render) {
	id, err := domain.UserDomain.CreateUser(&newUser)

	if err != nil {
		r.JSON(400, err.Error())
	} else {
		r.JSON(200, map[string]interface{}{"id": id})
	}

	return
}

func UserById(params martini.Params, r render.Render) {
	u, err := domain.UserDomain.UserByUid(domain.UserUid(params["id"]))

	if err != nil {
		r.JSON(400, err.Error())
	} else {
		r.JSON(200, u)
	}
}
