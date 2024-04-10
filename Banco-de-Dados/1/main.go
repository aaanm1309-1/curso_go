package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	uuid "github.com/satori/go.uuid"
)

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.NewV4().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	product := NewProduct("Notebook", 10500.99)
	err = InsertProduct(db, product)
	if err != nil {
		panic(err)
	}

}

func InsertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("insert into products (id, name, price) values(?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		panic(err)
	}
	return nil
}
