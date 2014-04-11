package main

import (
	"log"
	"os"

	"github.com/brianstarke/dogfort/domain"
	"github.com/brianstarke/dogfort/routes"
	"github.com/go-martini/martini"
	_ "github.com/joho/godotenv/autoload" // load all .env variables
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

func main() {
	m := martini.Classic()

	// JSON rendering middleware
	m.Use(render.Renderer(render.Options{IndentJSON: true}))

	// puts references all the initialized domain objects in the middleware layer
	m.Use(domain.DomainMiddleware())

	// user routes
	m.Post("/api/v1/users", binding.Json(domain.NewUser{}), binding.ErrorHandler, routes.CreateUser)

	// start server
	log.Printf("dogfort starting on %s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	m.Run()
}
