package routes

import (
	"encoding/json"
	"net/http"

	"github.com/brianstarke/dogfort/domain"
	"github.com/codegangsta/martini-contrib/render"
)

// TODO unified error handling
func CreateUser(userDomain *domain.UserDomain, req *http.Request, r render.Render) {
	var newUser domain.NewUser

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&newUser)

	if err != nil {
		r.JSON(400, err.Error())
		return
	}

	// TODO validation

	id, err := userDomain.CreateUser(&newUser)

	if err != nil {
		r.JSON(400, err.Error())
	} else {
		r.JSON(200, map[string]interface{}{"id": id})
	}

	return
}
