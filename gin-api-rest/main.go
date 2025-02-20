package main

import (
	"golang/api-go-rest/models"
	"golang/api-go-rest/routes"
)

func main() {
	models.Students = append(models.Students, models.Student{ID: 1, Name: "João da Silva", CPF: "123.456.789-00", RG: "12.345.678-9"})
	models.Students = append(models.Students, models.Student{ID: 2, Name: "Maria da Silva", CPF: "987.654.321-00", RG: "98.765.432-1"})
	models.Students = append(models.Students, models.Student{ID: 3, Name: "José da Silva", CPF: "456.789.123-00", RG: "45.678.912-3"})
	models.Students = append(models.Students, models.Student{ID: 4, Name: "Joana da Silva", CPF: "654.321.987-00", RG: "65.432.198-7"})
	routes.HandleRequests()
}
