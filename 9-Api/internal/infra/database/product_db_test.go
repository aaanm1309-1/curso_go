package database

import (
	"fmt"
	"math/rand"
	"testing"

	"br.com.adrianomenezes/cursogo/9-Api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateOneProduct(t *testing.T) (*entity.Product, *gorm.DB, *Product) {
	db, productDb := OpenConnection(t)
	product, _ := entity.NewProduct("Notebook 1", 1500.10)

	err := productDb.Create(product)
	assert.Nil(t, err)
	return product, db, productDb
}

func OpenConnection(t *testing.T) (*gorm.DB, *Product) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	productDb := NewProduct(db)
	return db, productDb
}

func TestCreateNewProduct(t *testing.T) {
	product, _, productDb := CreateOneProduct(t)

	productFound, err := productDb.FindByID(product.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
	assert.NotNil(t, product.ID)
	assert.NotNil(t, product.Name)
	assert.NotNil(t, product.Price)
	assert.Greater(t, product.Price, 0.0)
}

func generateAndSaveProducts(t *testing.T, db *gorm.DB) {
	for i := 1; i < 45; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.NoError(t, err)
		db.Create(product)
	}
}

func TestFindAllProducts(t *testing.T) {
	db, productDb := OpenConnection(t)

	generateAndSaveProducts(t, db)

	products, err := productDb.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Equal(t, products[0].Name, "Product 1")
	assert.Equal(t, products[9].Name, "Product 10")

	products, err = productDb.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Equal(t, products[0].Name, "Product 11")
	assert.Equal(t, products[9].Name, "Product 20")

	products, err = productDb.FindAll(0, 0, "desc")
	assert.NoError(t, err)
	assert.Equal(t, products[0].Name, "Product 44")
	assert.Equal(t, products[9].Name, "Product 35")

	products, err = productDb.FindAll(2, 10, "")
	assert.NoError(t, err)
	assert.Equal(t, products[0].Name, "Product 11")
	assert.Equal(t, products[9].Name, "Product 20")
}

func TestFindById(t *testing.T) {
	product, _, productDb := CreateOneProduct(t)

	productFound, err := productDb.FindByID(product.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
	assert.NotNil(t, product.ID)
	assert.NotNil(t, product.Name)
	assert.NotNil(t, product.Price)
	assert.Greater(t, product.Price, 0.0)
}
