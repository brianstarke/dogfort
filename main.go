package main

import (
	"log"
	"os"

	"github.com/codegangsta/martini"
	_ "github.com/joho/godotenv/autoload" // load all .env variables
)

func main() {
	log.Printf("dogfort starting on %s:%s", os.Getenv("HOST"), os.Getenv("PORT"))

	m := martini.Classic()

	m.Run()
}
