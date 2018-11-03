package main

import (
	"mphclub-rest-server/database"
	"mphclub-rest-server/server"
)

func main() {
	database.CreateSchema()

	server.CreateAndListen()
}
