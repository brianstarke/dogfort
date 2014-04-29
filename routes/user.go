package routes

import (
	"net/http"

	"github.com/brianstarke/dogfort/domain"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func UserCreate(userDomain *domain.UserDomain, newUser domain.NewUser, req *http.Request, r render.Render) {
	id, err := userDomain.CreateUser(&newUser)

	if err != nil {
		r.JSON(400, err.Error())
	} else {
		r.JSON(200, map[string]interface{}{"id": id})
	}

	return
}

func UserById(userDomain *domain.UserDomain, params martini.Params, r render.Render) {
	u, err := userDomain.UserByUid(domain.UserUid(params["id"]))

	if err != nil {
		r.JSON(400, err.Error())
	} else {
		r.JSON(200, u)
	}
}
