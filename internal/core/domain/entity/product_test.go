package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	name := "Product Name"
	description := "Product Description"
	categoryId := 1
	amount := 99.99

	product, err := NewProduct(name, description, categoryId, amount)
	assert.NoError(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.Equal(t, name, product.Name)
	assert.Equal(t, description, product.Description)
	assert.Equal(t, categoryId, product.CategoryId)
	assert.Equal(t, amount, product.Amount)
	assert.WithinDuration(t, time.Now().UTC(), product.CreatedAt, time.Second)
	assert.WithinDuration(t, time.Now().UTC(), product.UpdatedAt, time.Second)
}

func TestNewProduct_DifferentValues(t *testing.T) {
	name := "Another Product"
	description := "Another Description"
	categoryId := 5
	amount := 123.45

	product, err := NewProduct(name, description, categoryId, amount)
	assert.NoError(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, name, product.Name)
	assert.Equal(t, description, product.Description)
	assert.Equal(t, categoryId, product.CategoryId)
	assert.Equal(t, amount, product.Amount)
}
