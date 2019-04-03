package main

import (
	"flag"
	"log"

	"github.com/binkkatal/go-coffee-pods/dao"
	"github.com/binkkatal/go-coffee-pods/handler"
)

func main() {
	// Get path to assets from flag.
	assetsDir := flag.String("assetsDir", "./assets", "defines the path to asset directory")
	port := flag.String("port", ":8080", "sets the port to serve")
	flag.Parse()
	
	// I've done some abstractions here using interfaces
	// and custom structs (rather than using database/sql itself) in order to separate the db
	// from the methods for easy synchronous unit testing
	
	db, err := dao.NewDB()
	if err != nil {
		log.Fatalf("unable to connect to db, error: %v", err)
	}

	env := &handler.Env{DS: db, AssetsDir: assetsDir}
	// Initialize routes.
	r := handler.InitRoutes(env)

	r.Logger.Fatal(r.Start(*port))
}
