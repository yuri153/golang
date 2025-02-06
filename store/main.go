package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

type Product struct {
	Id          int
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
	db := connectDataBase()

	produtosDb, err := db.Query("select * from products")

	defer db.Close()

	if err != nil {
		panic(err.Error())
	}

	produtos := []Product{}

	for produtosDb.Next() {

		var product Product

		err = produtosDb.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Quantity)

		if err != nil {
			panic(err.Error())
		}

		produtos = append(produtos, product)
	}

	templates.ExecuteTemplate(writer, "Index", produtos)
}

func connectDataBase() *sql.DB {
	connection := "user=postgres dbname=store password=1234 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}

	return db
}
