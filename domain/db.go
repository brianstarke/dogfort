package domain

import (
	"log"
	"os"

	"github.com/codegangsta/martini"
	"labix.org/v2/mgo"
)

/*
Make all domain structs available as martini middleware
*/
func DomainMiddleware() martini.Handler {
	db := intializeMongo()

	return func(context martini.Context) {
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
	} else {
		return session.DB(os.Getenv("MONGO_DB"))
	}
}
