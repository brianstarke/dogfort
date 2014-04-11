package routes

import (
	"net/http"

	"github.com/brianstarke/dogfort/domain"
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
