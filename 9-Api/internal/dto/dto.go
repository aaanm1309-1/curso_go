package dto

import "br.com.adrianomenezes/cursogo/9-Api/pkg/entity"

type CreateProduct struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type UpdateProduct struct {
	ID    entity.ID `json:"id"`
	Name  string    `json:"name"`
	Price float64   `json:"price"`
}
type CreateUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
