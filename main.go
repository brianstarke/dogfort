package main

import (
	"log"
	"os"
	"runtime"

	"github.com/brianstarke/dogfort/domain"
	"github.com/brianstarke/dogfort/routes"
	"github.com/go-martini/martini"
	_ "github.com/joho/godotenv/autoload" // load all .env variables
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	m := martini.Classic()

	// JSON rendering middleware
	m.Use(render.Renderer(render.Options{IndentJSON: true}))

	// puts references all the initialized domain objects in the middleware layer
	m.Use(domain.DomainMiddleware())

	// user routes
	m.Post("/api/v1/users", binding.Json(domain.NewUser{}), binding.ErrorHandler, routes.CreateUser)
	m.Post("/api/v1/authenticate", binding.Json(domain.AuthenticationRequest{}), binding.ErrorHandler, routes.AuthenticateUser)
	m.Get("/api/v1/user", domain.AuthenticationMiddleware, routes.GetAuthenticatedUser) // who am I?!
	m.Get("/api/v1/users/:userId", domain.AuthenticationMiddleware, routes.GetUserById)

	// channel routes
	m.Post("/api/v1/channels", domain.AuthenticationMiddleware, binding.Json(domain.Channel{}), binding.ErrorHandler, routes.CreateChannel)
	m.Get("/api/v1/channels", domain.AuthenticationMiddleware, routes.ListChannels)
	m.Get("/api/v1/channels/user", domain.AuthenticationMiddleware, routes.GetUserChannels)
	m.Get("/api/v1/channels/join/:channelId", domain.AuthenticationMiddleware, routes.JoinChannel)
	m.Get("/api/v1/channels/leave/:channelId", domain.AuthenticationMiddleware, routes.LeaveChannel)

	// message routes
	m.Get("/api/v1/messages/:channelId", domain.AuthenticationMiddleware, routes.MessagesByChannel)
	m.Post("/api/v1/messages", domain.AuthenticationMiddleware, binding.Json(domain.Message{}), binding.ErrorHandler, routes.CreateMessage)

	// start server
	log.Printf("dogfort starting on %s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	m.Run()
}
