package main

import (
	"fmt"
	"golang/api-go-rest/database"
	"golang/api-go-rest/routes"
)

func main() {
	database.ConnectDatabse()

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
