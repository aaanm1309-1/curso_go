package repositorio

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/aaanm1309-1/curso_go/banco-de-dados/1/struct_product"
)

func InsertProduct(db *sql.DB, product *struct_product.Product) error {
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

func UpdateProduct(db *sql.DB, product *struct_product.Product) error {
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

func SelectProduct(ctx context.Context, db *sql.DB, id string) (*struct_product.Product, error) {
	stmt, err := db.Prepare("select id, name, price from products where id = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	var p struct_product.Product
	err = stmt.QueryRowContext(ctx, id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		panic(err)
	}
	return &p, nil
}

func SelectAllProducts(ctx context.Context, db *sql.DB) ([]struct_product.Product, error) {
	rows, err := db.QueryContext(ctx, "select id, name, price from products")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var products []struct_product.Product
	for rows.Next() {
		var p struct_product.Product
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

func PrintAllProducts(ctx context.Context, db *sql.DB) {
	rows, err := db.QueryContext(ctx, "select id, name, price from products")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var products []struct_product.Product
	for rows.Next() {
		var p struct_product.Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			panic(err)
		}
		products = append(products, p)
	}
	for _, p := range products {
		fmt.Printf("ID: %v, Product: %v, possui o preço de R$ %.2f\n", p.ID, p.Name, p.Price)
	}

}

func PrintAllProd(ctx context.Context, db *sql.DB) {
	products, err := SelectAllProducts(ctx, db)
	if err != nil {
		panic(err)
	}
	for _, p := range products {
		fmt.Printf("ID: %v, Product: %v, possui o preço de R$ %.2f\n", p.ID, p.Name, p.Price)
	}

}

func DeleteAllProduct(ctx context.Context, db *sql.DB) error {
	stmt, err := db.Prepare("delete from products")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx)
	if err != nil {
		panic(err)
	}
	return nil
}
