package main

import (
	"golang/api-go-rest/database"
	"golang/api-go-rest/routes"
)

func main() {
	database.ConnectDatabse()
	routes.HandleRequests()
}
