package main

import (
	"fmt"
	"golang/api-go-rest/models"
	"golang/api-go-rest/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Id: 1, Name: "Nome 1", History: "Historia 1"},
		{Id: 2, Name: "Nome 2", History: "Historia 2"},
	}

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
