package main

import (
	"github.com/maickmachado/upvote-api/database"
	"github.com/maickmachado/upvote-api/routes"
)

func init() {
	database.CreateMongoBD()
}

func main() {
	routes.HandleRequest()
}
