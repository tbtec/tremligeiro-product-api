package controller

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tbtec/tremligeiro/internal/core/domain/usecase"
)

func TestBuildProductUpdateResponse(t *testing.T) {
	now := time.Now()
	output := usecase.CreateProductOutput{
		ProductId:    "prod-123",
		Name:         "Test Product",
		Description:  "A product for testing",
		Amount:       99.99,
		CategoryID:   42,
		CategoryName: "Test Category",
		CreatedAt:    now,
		UpdatedAt:    now.Add(time.Hour),
	}

	resp := buildProductUpdateResponse(output)

	assert.Equal(t, output.ProductId, resp.ProductId)
	assert.Equal(t, output.Name, resp.Name)
	assert.Equal(t, output.Description, resp.Description)
	assert.Equal(t, output.Amount, resp.Amount)
	assert.Equal(t, output.CategoryID, resp.Category.CategoryID)
	assert.Equal(t, output.CategoryName, resp.Category.CategoryName)
	assert.Equal(t, output.CreatedAt, resp.CreatedAt)
	assert.Equal(t, output.UpdatedAt, resp.UpdatedAt)
}
