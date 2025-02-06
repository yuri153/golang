package models

import (
	"golang/store/database"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetProducts() []Product {
	db := database.ConnectDataBase()

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

	return produtos
}

func CreateProduct(product Product) {
	db := database.ConnectDataBase()

	insertProduct, err := db.Prepare("insert into products(name, description, price, quantity) values($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insertProduct.Exec(product.Name, product.Description, product.Price, product.Quantity)

	defer db.Close()
}

func DeleteProduct(id string) {
	db := database.ConnectDataBase()

	insertProduct, err := db.Prepare("delete from products where id = $1")

	if err != nil {
		panic(err.Error())
	}

	insertProduct.Exec(id)

	defer db.Close()
}
