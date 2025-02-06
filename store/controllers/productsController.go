package controllers

import (
	"golang/store/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(writer http.ResponseWriter, request *http.Request) {
	products := models.GetProducts()

	templates.ExecuteTemplate(writer, "Index", products)
}

func New(writer http.ResponseWriter, request *http.Request) {
	templates.ExecuteTemplate(writer, "New", nil)
}

func Insert(writer http.ResponseWriter, request *http.Request) {

	if request.Method == "POST" {
		name := request.FormValue("Name")
		description := request.FormValue("Description")
		priceString := request.FormValue("Price")
		quantityString := request.FormValue("Quantity")

		price, err := strconv.ParseFloat(priceString, 64)

		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		quantity, err := strconv.Atoi(quantityString)

		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		product := models.Product{Name: name, Description: description, Price: price, Quantity: quantity}

		models.CreateProduct(product)
	}

	http.Redirect(writer, request, "/", http.StatusMovedPermanently)
}

func Delete(writer http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")

	models.DeleteProduct(id)

	http.Redirect(writer, request, "/", http.StatusMovedPermanently)
}

func Edit(writer http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")

	product := models.GetProduct(id)

	templates.ExecuteTemplate(writer, "Edit", product)
}

func Update(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		idString := request.FormValue("Id")
		name := request.FormValue("Name")
		description := request.FormValue("Description")
		priceString := request.FormValue("Price")
		quantityString := request.FormValue("Quantity")

		price, err := strconv.ParseFloat(priceString, 64)

		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		quantity, err := strconv.Atoi(quantityString)

		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		id, err := strconv.Atoi(idString)

		if err != nil {
			log.Println("Erro na conversão do id:", err)
		}

		product := models.Product{Id: id, Name: name, Description: description, Price: price, Quantity: quantity}

		models.UpdateProduct(product)
	}

	http.Redirect(writer, request, "/", http.StatusMovedPermanently)
}
