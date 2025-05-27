package repository

import (
	"context"
	"errors"

	"github.com/stretchr/testify/assert"
	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/infra/database/model"
)

type MockProductRepo struct {
	CreateFunc         func(ctx context.Context, product *model.Product) error
	DeleteByIdFunc     func(ctx context.Context, id string) (*model.Product, error)
	FindByCategoryFunc func(ctx context.Context, categoryId int) (*[]model.Product, error)
	FindOneFunc        func(ctx context.Context, id string) (*model.Product, error)
	UpdateByIdFunc     func(ctx context.Context, product *model.Product) error
	ExecuteFunc        func(ctx context.Context, productId string) (string, error)
}

type MockCategoryRepo struct {
	FindByIdFunc func(id int) *entity.Category
}

func (m *MockProductRepo) Execute(ctx context.Context, productId string) (string, error) {
	return m.ExecuteFunc(ctx, productId)
}

func (m *MockCategoryRepo) FindById(id int) *entity.Category {
	return m.FindByIdFunc(id)
}

func (m *MockProductRepo) Create(ctx context.Context, product *model.Product) error {
	return m.CreateFunc(ctx, product)
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

func (m *MockProductRepo) UpdateById(ctx context.Context, product *model.Product) error {
	if m.UpdateByIdFunc != nil {
		return m.UpdateByIdFunc(ctx, product)
	}
	return nil
}

// Mock compatível com a interface IProductRepository
type MockProductRepoInterface struct{}

func (m *MockProductRepoInterface) Create(ctx context.Context, p *model.Product) error { return nil }
func (m *MockProductRepoInterface) FindOne(ctx context.Context, id string) (*model.Product, error) {
	return nil, nil
}
func (m *MockProductRepoInterface) FindByCategory(ctx context.Context, categoryId int) (*[]model.Product, error) {
	return nil, nil
}
func (m *MockProductRepoInterface) UpdateById(ctx context.Context, p *model.Product) error {
	return nil
}
func (m *MockProductRepoInterface) DeleteById(ctx context.Context, id string) (*model.Product, error) {
	return nil, nil
}

// Mock compatível com a interface ICategoryRepository
type MockCategoryRepoInterface struct{}

func (m *MockCategoryRepoInterface) FindById(id int) *entity.Category {
	return &entity.Category{ID: id, Name: "Mock"}
}

// Mock category repo that returns nil
type MockCategoryRepoNotFound struct{}

func (m *MockCategoryRepoNotFound) FindById(id int) *entity.Category { return nil }

// Mock product repo that returns error
type MockProductRepoError struct{}

func (m *MockProductRepoError) Create(ctx context.Context, p *model.Product) error {
	return assert.AnError
}

func (m *MockProductRepoError) DeleteById(ctx context.Context, id string) (*model.Product, error) {
	return nil, errors.New("erro ao deletar produto")
}

func (m *MockProductRepoError) FindByCategory(ctx context.Context, categoryId int) (*[]model.Product, error) {
	return nil, nil
}

func (m *MockProductRepoError) FindOne(ctx context.Context, id string) (*model.Product, error) {
	return nil, errors.New("not implemented")
}

func (m *MockProductRepoError) FindById(ctx context.Context, id int64) (*entity.Product, error) {
	return nil, errors.New("erro ao buscar produto")
}

func (m *MockProductRepoError) FindAll(ctx context.Context) ([]*entity.Product, error) {
	return nil, errors.New("erro ao buscar produtos")
}

// Add the missing UpdateById method to satisfy the IProductRepository interface
func (m *MockProductRepoError) UpdateById(ctx context.Context, product *model.Product) error {
	return errors.New("erro ao atualizar produto")
}
