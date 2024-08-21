package main

import (
	"net/http"

	"br.com.adrianomenezes/cursogo/9-Api/configs"
	"br.com.adrianomenezes/cursogo/9-Api/internal/entity"
	"br.com.adrianomenezes/cursogo/9-Api/internal/infra/database"
	"br.com.adrianomenezes/cursogo/9-Api/internal/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	// println(config.DBDriver)
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/products", productHandler.GetAllProducts)
	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)

	r.Post("/users", userHandler.CreateUser)
	r.Get("/users/{id}", userHandler.ListUser)
	r.Get("/users", userHandler.GetAllUsers)

	http.ListenAndServe(":8000", r)

}
