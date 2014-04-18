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
	usersCollection    = "users"
	channelsCollection = "channels"
)

/*
Make all domain structs available as martini middleware
*/
func DomainMiddleware() martini.Handler {
	db := intializeMongo()

	userDomain := &UserDomain{db.C(usersCollection)}
	channelDomain := &ChannelDomain{db.C(channelsCollection)}

	return func(context martini.Context) {
		context.Map(userDomain)
		context.Map(channelDomain)

		context.Next()
	}
}

/*
Check Authorization token
*/
func AuthenticationMiddleware(req *http.Request, context martini.Context, r render.Render) {
	token := req.Header.Get("Authorization")

	if token == "" {
		r.Error(401)
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
