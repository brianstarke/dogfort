package routes

import (
	"net/http"

	"github.com/brianstarke/dogfort/domain"
	"github.com/martini-contrib/render"
)

/*
  Attempts to authenticate a user, on success issues a JWT
*/
func UserAuthenticate(userDomain *domain.UserDomain, ar domain.AuthenticationRequest, req *http.Request, r render.Render) {
	jwt, err := userDomain.Authenticate(&ar)

	if err != nil {
		r.JSON(400, err.Error())
	} else {
		r.JSON(200, map[string]interface{}{"token": jwt})
	}
}

/*
  Uses the JWT in the Authorization header to look up and return the authenticated user
*/
func UserByToken(userDomain *domain.UserDomain, userUid domain.UserUid, r render.Render) {
	u, err := userDomain.UserByUid(userUid)

	if err != nil {
		r.JSON(400, err.Error())
	} else {
		r.JSON(200, map[string]interface{}{"user": u})
	}
}
