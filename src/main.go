package main

import (
	"fmt"

	"github.com/aliaydins/todo/config"
	"github.com/aliaydins/todo/pkg/router"
	"github.com/aliaydins/todo/pkg/server"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config := config.LoadConfig(".")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  config.GetDBURL(),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		fmt.Println("Couldn't connect to the DB.")
	}

	r := router.Setup(db, config.SecretKey)

	err = server.NewServer(r, config.AppPort).Run()
	if err != nil {
		panic("Couldn't start the HTTP server.")
	}

}
