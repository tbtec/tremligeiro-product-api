package gateway

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/dto"
)

type MockProductGateway struct {
	CreateFunc func(ctx context.Context, product *entity.Product) error
}

type MockCategoryGateway struct {
	FindByIdFunc func(id int) *entity.Category
}

type MockProductPresenter struct {
	BuildProductCreateResponseFunc func(product entity.Product, category entity.Category) dto.Product
}

func (m *MockProductGateway) Create(ctx context.Context, product *entity.Product) error {

	return m.CreateFunc(ctx, product)
}

func (m *MockCategoryGateway) FindById(id int) *entity.Category {
	return m.FindByIdFunc(id)
}

func (m *MockProductPresenter) BuildProductCreateResponse(product entity.Product, category entity.Category) dto.Product {
	return m.BuildProductCreateResponseFunc(product, category)
}
