package main

import (
	"flag"
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

var (
	appPath        = flag.String("appPath", "public", "path to dogfort app")
	apiRoot string = "/api/v1"
)

func main() {
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	// start hub
	go hub.H.Run()

	m := martini.Classic()

	// serve app html/js as well
	log.Printf("Serving dogfort app from [%s]", *appPath)
	m.Use(martini.Static(*appPath))

	// JSON rendering
	m.Use(render.Renderer(render.Options{IndentJSON: true}))

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
		m.Get("/", domain.AuthenticationMiddleware, routes.UsersOnline)
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
		m.Get("/channel/:id/before/:before/num/:num", routes.MessagesByChannel)
	}, domain.AuthenticationMiddleware)

	/*
	   Receive github notifications for dogfort updates channel
	*/
	m.Post(apiRoot+"/github/:channelId", binding.Json(routes.GithubMsg{}), binding.ErrorHandler, routes.GithubHandler)

	// socket connector
	m.Get("/ws/connect", domain.AuthenticationMiddleware, hub.WsHandler)

	// start server
	log.Printf("dogfort starting on %s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	m.Run()

}
