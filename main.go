package main

import (
	"sejutacita/config"
	"sejutacita/lib/databases"
	"sejutacita/middlewares"
	"sejutacita/routes"
)

func main() {
	config.InitDB()
	databases.MockDataAdmin()
	e := routes.New()
	middlewares.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8080"))
}
