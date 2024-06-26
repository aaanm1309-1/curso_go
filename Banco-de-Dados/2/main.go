package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ProductOrm struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8&parseTime=True&loc=Local"

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	// defer cancel()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&ProductOrm{})
	// // defer db.Close()

	db.Create(&ProductOrm{
		Name: "Notebook", Price: 1200.00,
	})

	// products := []ProductOrm{
	// 	{Name: "Video", Price: 1500.00},
	// 	{Name: "Mouse", Price: 70.00},
	// 	{Name: "Keyboard", Price: 110.00},
	// }

	// db.Create(&products)

	// //select one
	// var product ProductOrm
	// // db.First(&product, 2)
	// // fmt.Println(product)
	// // fmt.Println("________________________________________________________________")
	// db.First(&product, "name = ?", "Mouse")
	// fmt.Println(product)
	// fmt.Println("________________________________________________________________")

	// //select all
	var products []ProductOrm
	// // db.First(&product, 2)
	// // fmt.Println(product)
	// // fmt.Println("________________________________________________________________")
	// db.Limit(2).Offset(2).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }
	// fmt.Println("________________________________________________________________")

	// //select com where
	// db.Where("price > ?", 100).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }
	// fmt.Println("________________________________________________________________")

	//select one and update
	// var p ProductOrm
	// db.First(&p, 1)
	// fmt.Println(p)
	// fmt.Println("________________________________________________________________")
	// p.Name = "New Notebook"
	// db.Save(&p)

	var p2 ProductOrm
	db.First(&p2, 1)
	fmt.Println(p2)
	fmt.Println("________________________________________________________________")

	// db.Delete(&p2)
	db.Delete(&ProductOrm{}, p2.ID)

	//select all
	db.Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}
	fmt.Println("________________________________________________________________")

}
