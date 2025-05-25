package fixtures

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/infra/database/model"
)

type MockProductRepo struct {
	CreateFunc         func(ctx context.Context, p *model.Product) error
	DeleteByIdFunc     func(ctx context.Context, id string) (*model.Product, error)
	FindByCategoryFunc func(ctx context.Context, categoryId int) (*[]model.Product, error)
	FindOneFunc        func(ctx context.Context, id string) (*model.Product, error)
	UpdateByIdFunc     func(ctx context.Context, p *model.Product) error
}

func (m *MockProductRepo) Create(ctx context.Context, p *model.Product) error {
	return m.CreateFunc(ctx, p)
}

func (m *MockProductRepo) DeleteById(ctx context.Context, id string) (*model.Product, error) {
	if m.DeleteByIdFunc != nil {
		return m.DeleteByIdFunc(ctx, id)
	}
	return nil, nil
}

func (m *MockProductRepo) FindByCategory(ctx context.Context, categoryId int) (*[]model.Product, error) {
	if m.FindByCategoryFunc != nil {
		return m.FindByCategoryFunc(ctx, categoryId)
	}
	return nil, nil
}

func (m *MockProductRepo) FindOne(ctx context.Context, id string) (*model.Product, error) {
	if m.FindOneFunc != nil {
		return m.FindOneFunc(ctx, id)
	}
	return nil, nil
}

func (m *MockProductRepo) UpdateById(ctx context.Context, p *model.Product) error {
	if m.UpdateByIdFunc != nil {
		return m.UpdateByIdFunc(ctx, p)
	}
	return nil
}
