package main

import (
	"ato-chat/config"
	"ato-chat/internal/user"
	"ato-chat/internal/ws"
	"ato-chat/router"
	"log"
)

func main() {
	dbConn, err := config.DatabaseConnection()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}

	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	router.InitRouter(userHandler, wsHandler)
	router.Start("0.0.0.0:8080")

}
