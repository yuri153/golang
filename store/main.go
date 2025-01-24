package main

import (
	"html/template"
	"net/http"
)

type Product struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(writer http.ResponseWriter, request *http.Request) {
	produtos := []Product{
		{"Camiseta", "Azul, bem bonita", 39, 5},
		{"Tênis", "Confortável", 89, 3},
		{"Fone", "Muito bom", 59, 2},
	}

	templates.ExecuteTemplate(writer, "Index", produtos)
}
