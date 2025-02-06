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

	produtosDb, err := db.Query("select * from products order by id asc")

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

func GetProduct(id string) Product {
	db := database.ConnectDataBase()

	productDb, err := db.Query("select id, name, description, price, quantity from products where id = $1", id)

	defer db.Close()

	if err != nil {
		panic(err.Error())
	}

	product := Product{}

	for productDb.Next() {

		err = productDb.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Quantity)

		if err != nil {
			panic(err.Error())
		}
	}

	return product
}

func UpdateProduct(product Product) {
	db := database.ConnectDataBase()

	insertProduct, err := db.Prepare("update products set name = $1, description = $2, price = $3, quantity = $4 where id = $5")

	if err != nil {
		panic(err.Error())
	}

	insertProduct.Exec(product.Name, product.Description, product.Price, product.Quantity, product.Id)

	defer db.Close()
}
