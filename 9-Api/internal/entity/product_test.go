package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Produto 1", 150.01)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.Name)
	assert.Equal(t, product.Name, "Produto 1")
	assert.NotEmpty(t, product.Price)
	assert.Equal(t, product.Price, 150.01)
	assert.NotEqual(t, product.Price, 150)
	assert.NotEmpty(t, product.CreatedAt)
	assert.NotEqual(t, product.CreatedAt, time.Now())

	product, err = NewProduct("Pr", 90.01)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, err, ErrInvalidName)

	product, err = NewProduct("Produto 2", 0)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, err, ErrPriceIsRequired)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	product, err := NewProduct("", 90.01)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, err, ErrNameIsRequired)

}

func TestProductWhenNameIsInvalid(t *testing.T) {
	product, err := NewProduct("Pr", 90.01)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, err, ErrInvalidName)

}

func TestProductValidate(t *testing.T) {
	product, err := NewProduct("Prod", 90.01)
	assert.NotNil(t, product)
	assert.Nil(t, err)
	assert.Nil(t, product.Validate())

}
func TestProductWhenPriceIsRequired(t *testing.T) {
	product, err := NewProduct("Prod", 0)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, err, ErrPriceIsRequired)

}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	product, err := NewProduct("Prod", -1.0)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, err, ErrInvalidPrice)

}
