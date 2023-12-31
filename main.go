package main

import (
	"food-comming-api/config"
	"food-comming-api/database"
	"food-comming-api/server"
)

func main() {
	cfg := config.GetConfig()

	db := database.NewPostgresDatabase(&cfg)

	server.NewEchoServer(&cfg, db.GetDb()).Start()
}
