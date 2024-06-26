package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type ProductCatOrm struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&ProductCatOrm{}, &Category{})

	// category := Category{Name: "Eletronicos"}
	// db.Create(&category)

	// db.Create(&ProductCatOrm{
	// 	Name:       "Notebook",
	// 	Price:      1356.00,
	// 	CategoryID: category.ID,
	// })

	// db.Create(&ProductCatOrm{
	// 	Name:       "Mouse",
	// 	Price:      56.00,
	// 	CategoryID: 1,
	// })

	var products []ProductCatOrm
	db.Preload("Category").Find(&products)
	for _, p := range products {
		fmt.Println(p)
	}

}
