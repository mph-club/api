package main

import (
	"log"
	"mphclub-rest-server/database"
	"mphclub-rest-server/server"
)

func main() {
	log.Println("starting create schema")
	database.CreateSchema()

	server.CreateAndListen()
}
