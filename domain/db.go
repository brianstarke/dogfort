package domain

import (
	"log"
	"os"

	"github.com/codegangsta/martini"
	"labix.org/v2/mgo"
)

var (
	usersCollection = "users"
)

/*
Make all domain structs available as martini middleware
*/
func DomainMiddleware() martini.Handler {
	db := intializeMongo()

	userDomain := &UserDomain{db.C(usersCollection)}

	return func(context martini.Context) {
		context.Map(userDomain)

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
