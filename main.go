package main

import (
	"log"
	"os"
	"runtime"

	"github.com/brianstarke/dogfort/domain"
	"github.com/brianstarke/dogfort/hub"
	"github.com/brianstarke/dogfort/routes"
	"github.com/go-martini/martini"
	_ "github.com/joho/godotenv/autoload" // load all .env variables
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

var apiRoot string = "/api/v1"

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// start hub
	go hub.H.Run()

	m := martini.Classic()

	// JSON rendering
	m.Use(render.Renderer(render.Options{IndentJSON: true}))

	// references to all the initialized domain objects
	m.Use(domain.DomainMiddleware())

	/*
	  Authentication

	  POST returns a JWT token
	  GET (with JWT in the Authorization header) returns the authenticated user
	*/
	m.Group(apiRoot+"/auth", func(r martini.Router) {
		r.Post("", binding.Json(domain.AuthenticationRequest{}), binding.ErrorHandler, routes.UserAuthenticate)
		r.Get("", domain.AuthenticationMiddleware, routes.UserByToken)
	})

	/*
	   Users
	*/
	m.Group(apiRoot+"/users", func(r martini.Router) {
		m.Post("", binding.Json(domain.NewUser{}), binding.ErrorHandler, routes.UserCreate)
		m.Get("/:id", domain.AuthenticationMiddleware, routes.UserById)
	})

	/*
	   Channels
	*/
	m.Group(apiRoot+"/channels", func(r martini.Router) {
		m.Post("", binding.Json(domain.Channel{}), binding.ErrorHandler, routes.ChannelCreate)
		m.Get("", routes.ChannelList)
		m.Get("/user/:id", routes.ChannelsByUser)
		m.Get("/join/:id", routes.ChannelJoin)
		m.Get("/leave/:id", routes.ChannelLeave)
	}, domain.AuthenticationMiddleware)

	/*
	  Messages
	*/
	m.Group(apiRoot+"/messages", func(r martini.Router) {
		m.Post("", binding.Json(domain.Message{}), binding.ErrorHandler, routes.CreateMessage)
		m.Get("/channel/:id", routes.MessagesByChannel)
	}, domain.AuthenticationMiddleware)

	// socket connector
	m.Get("/ws/connect", hub.WsHandler)

	// start server
	log.Printf("dogfort starting on %s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	m.Run()

}
