package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("______Inicio________________________")
	printAllProducts(ctx, db)

	product := NewProduct("Notebook", 10500.99)
	err = InsertProduct(db, product)
	if err != nil {
		panic(err)
	}

	fmt.Println("_________Após Insercao______________")
	printAllProducts(ctx, db)

	product.Price = 4200.50
	err = UpdateProduct(db, product)
	if err != nil {
		panic(err)
	}
	fmt.Println("________Após Update_________________")
	printAllProducts(ctx, db)

	p, err := SelectProduct(ctx, db, product.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println("________Select Unico________________")
	fmt.Printf("Product: %v, possui o preço de %.2f\n", p.Name, p.Price)

	fmt.Println("____________________________________")
	printAllProducts(ctx, db)

	err = DeleteProduct(ctx, db, product.ID)
	if err != nil {
		panic(err)
	}

	fmt.Println("_______Após Deleção______________")
	printAllProducts(ctx, db)

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

func UpdateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("update products set name = ?, price = ? where id = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		panic(err)
	}
	return nil
}

func SelectProduct(ctx context.Context, db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select id, name, price from products where id = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	var p Product
	err = stmt.QueryRowContext(ctx, id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		panic(err)
	}
	return &p, nil
}

func SelectAllProducts(ctx context.Context, db *sql.DB) ([]Product, error) {
	rows, err := db.QueryContext(ctx, "select id, name, price from products")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var products []Product
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			panic(err)
		}
		products = append(products, p)
	}
	return products, nil
}

func DeleteProduct(ctx context.Context, db *sql.DB, id string) error {
	stmt, err := db.Prepare("delete from products where id = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		panic(err)
	}
	return nil
}

func printAllProducts(ctx context.Context, db *sql.DB) {
	products, err := SelectAllProducts(ctx, db)
	if err != nil {
		panic(err)
	}
	for _, p := range products {
		fmt.Printf("ID: %v, Product: %v, possui o preço de R$ %.2f\n", p.ID, p.Name, p.Price)
	}

}
