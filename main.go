package main

import (
	"log"
	"os"

	"github.com/brianstarke/dogfort/domain"
	"github.com/brianstarke/dogfort/routes"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	_ "github.com/joho/godotenv/autoload" // load all .env variables
)

func main() {
	m := martini.Classic()

	// JSON rendering middleware
	m.Use(render.Renderer(render.Options{IndentJSON: true}))

	// puts references all the initialized domain objects in the middleware layer
	m.Use(domain.DomainMiddleware())

	// user routes
	m.Post("/api/v1/users", routes.CreateUser)

	// start server
	log.Printf("dogfort starting on %s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	m.Run()
}
