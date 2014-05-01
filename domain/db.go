package domain

import (
	"log"
	"net/http"
	"os"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"labix.org/v2/mgo"
)

var (
	db            = intializeMongo()
	UserDomain    = &userDomain{db.C("users")}
	ChannelDomain = &channelDomain{db.C("channels")}
	MessageDomain = &messageDomain{db.C("messages")}
)

/*
Check Authorization token
*/
func AuthenticationMiddleware(req *http.Request, context martini.Context, r render.Render) {
	token := req.Header.Get("Authorization")

	if token == "" {
		// failover to using cookie (mostly for socket auth TODO unfuck that)
		t, err := req.Cookie("dogfort_token")

		if err != nil {
			r.Error(401)
		} else {
			token = t.Value
		}
	}

	uid, err := getUserUidFromToken(token)

	if err != nil {
		r.Error(401)
	} else {
		context.Map(*uid)
		context.Next()
	}
}

/*
Initialize the mongo connection
*/
func intializeMongo() *mgo.Database {
	session, err := mgo.Dial(os.Getenv("MONGO_HOST"))

	if err != nil {
		log.Fatalf("Could not initialize Mongo connection, : %s", err.Error())

		return nil
	} else {
		return session.DB(os.Getenv("MONGO_DB"))
	}
}
