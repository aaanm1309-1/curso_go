package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID            int `gorm:"primaryKey"`
	Name          string
	ProductCatOrm []ProductCatOrm
}

type ProductCatOrm struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID              int `gorm:"primaryKey"`
	Number          string
	ProductCatOrmID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&ProductCatOrm{}, &Category{}, &SerialNumber{})

	// category := Category{Name: "Eletronicos"}
	// db.Create(&category)

	// category := Category{Name: "Telefones"}
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

	// db.Create(&ProductCatOrm{
	// 	Name:       "Iphone",
	// 	Price:      5500.00,
	// 	CategoryID: 4,
	// })

	// db.Create(&ProductCatOrm{
	// 	Name:       "S23",
	// 	Price:      4700.00,
	// 	CategoryID: 4,
	// })
	// db.Create(&ProductCatOrm{
	// 	Name:       "S24",
	// 	Price:      3400.00,
	// 	CategoryID: 4,
	// })
	// db.Create(&ProductCatOrm{
	// 	Name:       "S10",
	// 	Price:      1100.00,
	// 	CategoryID: 4,
	// })

	// db.Create(&ProductCatOrm{
	// 	Name:       "Holofote",
	// 	Price:      440.00,
	// 	CategoryID: 3,
	// })

	// db.Create(&SerialNumber{
	// 	Number:          "1234567890",
	// 	ProductCatOrmID: 4,
	// })

	// product := ProductCatOrm{}

	// db.Where("name = ?", "Iphone").First(&product)
	// fmt.Println(product)
	// db.UpdateColumn("CategoryID", 4)

	// db.Model(&ProductCatOrm{}).Where("name = ?", "Iphone").Update("CategoryID", 4)

	// fmt.Println(product)

	// var products []ProductCatOrm
	// db.Preload("Category").Preload("ProductCatOrms.SerialNumber").Find(&products)
	// for _, p := range products {
	// 	fmt.Println(p)
	// }

	var categories []Category
	db.Preload("ProductCatOrm").Preload("ProductCatOrm.SerialNumber").Find(&categories)
	for _, c := range categories {
		fmt.Println("Categoria: ", c.Name)
		for _, p := range c.ProductCatOrm {
			if p.SerialNumber.Number == "" {
				fmt.Println("- Produto: ", p.Name, " -  Valor: R$", p.Price)
			} else {
				fmt.Println("- Produto: ", p.Name, " -  Valor: R$", p.Price, " -  SerialNumber", p.SerialNumber.Number)
			}
		}
	}

	// err = db.Model(&Category{}).Preload("ProductCatOrm.SerialNumber").Find(&categories).Error
	// if err != nil {
	// 	panic(err)
	// }
	// for _, c := range categories {
	// 	fmt.Println("Categoria: ", c.Name)
	// 	for _, p := range c.ProductCatOrm {
	// 		fmt.Println("- Produto: ", p.Name, " -  Valor: R$", p.Price, " -  SerialNumber", p.SerialNumber.Number)
	// 	}
	// }

}
