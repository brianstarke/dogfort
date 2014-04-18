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
	m.Post("/api/v1/authenticate", binding.Json(domain.AuthenticationRequest{}), binding.ErrorHandler, routes.AuthenticateUser)
	m.Get("/api/v1/user", domain.AuthenticationMiddleware, routes.GetAuthenticatedUser) // who am I?!

	// channel routes
	m.Post("/api/v1/channels", domain.AuthenticationMiddleware, binding.Json(domain.Channel{}), binding.ErrorHandler, routes.CreateChannel)
	m.Get("/api/v1/channels", domain.AuthenticationMiddleware, routes.ListChannels)

	// start server
	log.Printf("dogfort starting on %s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	m.Run()
}
