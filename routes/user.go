package routes

import (
	"net/http"

	"github.com/brianstarke/dogfort/domain"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

// TODO unified error handling
func CreateUser(userDomain *domain.UserDomain, newUser domain.NewUser, req *http.Request, r render.Render) {
	id, err := userDomain.CreateUser(&newUser)

	if err != nil {
		r.JSON(400, err.Error())
	} else {
		r.JSON(200, map[string]interface{}{"id": id})
	}

	return
}

func AuthenticateUser(userDomain *domain.UserDomain, ar domain.AuthenticationRequest, req *http.Request, r render.Render) {
	jwt, err := userDomain.Authenticate(&ar)

	if err != nil {
		r.JSON(400, err.Error())
	} else {
		r.JSON(200, map[string]interface{}{"token": jwt})
	}
}

func VerifyUser(userDomain *domain.UserDomain, params martini.Params, r render.Render) {
	u, err := userDomain.Verify(params["token"])

	if err != nil {
		r.JSON(400, err.Error())
	} else {
		r.JSON(200, map[string]interface{}{"user": u})
	}
}
