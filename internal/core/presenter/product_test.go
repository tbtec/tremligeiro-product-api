package presenter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/dto"
)

func TestProductPresenter_BuildProductCreateResponse(t *testing.T) {
	p := NewProductPresenter()
	now := time.Now()

	product := entity.Product{
		ID:          "prod-123",
		Name:        "Product Name",
		Description: "Product Description",
		Amount:      42,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	category := entity.Category{
		ID:   7,
		Name: "Category Name",
	}

	result := p.BuildProductCreateResponse(product, category)

	assert.Equal(t, dto.Product{
		ProductId:   "prod-123",
		Name:        "Product Name",
		Description: "Product Description",
		Amount:      42,
		Category: dto.Category{
			ID:   7,
			Name: "Category Name",
		},
		CreatedAt: now,
		UpdatedAt: now,
	}, result)
}
