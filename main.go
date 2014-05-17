package main

import (
	"flag"
	"log"
	"net/http"
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
	appPath = flag.String("appPath", "public", "path to dogfort app")
	apiRoot = "/api/v1"
	tls     = flag.Bool("tls", false, "Enable TLS.")
)

func main() {
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	m := martini.Classic()

	// serve app html/js as well
	log.Println("Serving dogfort app from [%s]", *appPath)
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

	// start hub
	go hub.H.Run()

	// start server
	addr := os.Getenv("HOST") + ":" + os.Getenv("PORT")
	log.Println(addr)
	if *tls {
		/*
			// start http redirect
			go func() {
				if err := http.ListenAndServe(addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					http.Redirect(w, r, "https://"+addr+r.RequestURI, http.StatusMovedPermanently)
				})); err != nil {
					log.Fatal(err)
				}
			}()
		*/
		if err := http.ListenAndServeTLS(addr, os.Getenv("PEM_CERT"), os.Getenv("PEM_KEY"), m); err != nil {
			log.Fatal(err)
		}
	} else {
		// NOTE: Avoid 3000 default port in m.Run().
		if err := http.ListenAndServe(addr, m); err != nil {
			log.Fatal(err)
		}
	}
}
