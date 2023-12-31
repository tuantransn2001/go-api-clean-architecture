package main

import (
	"food-comming-api/config"
	"food-comming-api/database"
	"food-comming-api/user/entities"
)

func main() {
	cfg := config.GetConfig()
	db := database.NewPostgresDatabase(&cfg)
	userMigrate(db)
}

func userMigrate(db database.Database) {
	db.GetDb().Migrator().CreateTable(&entities.User{})
	db.GetDb().CreateInBatches(([]entities.User{}), 10)
}
