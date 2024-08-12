package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	bd "github.com/aaanm1309-1/curso_go/banco-de-dados/1/repositorio"
	"github.com/aaanm1309-1/curso_go/banco-de-dados/1/struct_product"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("______Inicio________________________")
	fmt.Println("______Limpando Produtos_____________")

	bd.DeleteAllProduct(ctx, db)

	fmt.Println("______Lista se há registros_________")
	bd.PrintAllProd(ctx, db)

	product := struct_product.NewProduct("Caderno", 120.59)
	bd.InsertProduct(db, product)

	product = struct_product.NewProduct("Lapis", 1.35)
	bd.InsertProduct(db, product)

	product = struct_product.NewProduct("Notebook", 10500.99)
	bd.InsertProduct(db, product)

	fmt.Println("_________Após Insercao______________")
	bd.PrintAllProd(ctx, db)

	product.Price = 4200.50
	err = bd.UpdateProduct(db, product)
	if err != nil {
		panic(err)
	}
	fmt.Println("________Após Update_________________")
	bd.PrintAllProd(ctx, db)

	p, err := bd.SelectProduct(ctx, db, product.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println("________Select Unico________________")
	fmt.Printf("Product: %v, possui o preço de %.2f\n", p.Name, p.Price)

	fmt.Println("____________________________________")
	bd.PrintAllProd(ctx, db)

	err = bd.DeleteProduct(ctx, db, product.ID)
	if err != nil {
		panic(err)
	}

	fmt.Println("_______Após Deleção______________")
	bd.PrintAllProducts(ctx, db)

}
