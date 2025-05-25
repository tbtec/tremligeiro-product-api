package gateway

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/database/model"
)

// Mock for IProductRepository
type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) Create(ctx context.Context, product *model.Product) error {
	args := m.Called(ctx, product)
	return args.Error(0)
}

func (m *MockProductRepository) FindByCategory(ctx context.Context, id int) (*[]model.Product, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*[]model.Product), args.Error(1)
}

func (m *MockProductRepository) DeleteById(ctx context.Context, id string) (*model.Product, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*model.Product), args.Error(1)
}

func (m *MockProductRepository) UpdateById(ctx context.Context, product *model.Product) error {
	args := m.Called(ctx, product)
	return args.Error(0)
}

func (m *MockProductRepository) FindOne(ctx context.Context, id string) (*model.Product, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*model.Product), args.Error(1)
}

func TestProductGateway_Create(t *testing.T) {
	mockRepo := new(MockProductRepository)
	gateway := NewProductGateway(mockRepo)
	ctx := context.Background()
	now := time.Now()
	product := &entity.Product{
		ID:          "1",
		Name:        "Test",
		Description: "Desc",
		CategoryId:  2,
		Amount:      10,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	mockRepo.On("Create", ctx, mock.AnythingOfType("*model.Product")).Return(nil)

	err := gateway.Create(ctx, product)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductGateway_FindByCategory(t *testing.T) {
	mockRepo := new(MockProductRepository)
	gateway := NewProductGateway(mockRepo)
	ctx := context.Background()
	now := time.Now()
	models := []model.Product{
		{
			ID:          "1",
			Name:        "Test",
			Description: "Desc",
			CategoryId:  2,
			Amount:      10,
			CreatedAt:   now,
			UpdatedAt:   now,
		},
	}
	mockRepo.On("FindByCategory", ctx, 2).Return(&models, nil)

	products, err := gateway.FindByCategory(ctx, 2)
	assert.NoError(t, err)
	assert.Len(t, products, 1)
	assert.Equal(t, "Test", products[0].Name)
	mockRepo.AssertExpectations(t)
}

func TestProductGateway_DeleteById(t *testing.T) {
	mockRepo := new(MockProductRepository)
	gateway := NewProductGateway(mockRepo)
	ctx := context.Background()
	now := time.Now()
	deletedProduct := &model.Product{
		ID:          "1",
		Name:        "Test",
		Description: "Desc",
		CategoryId:  2,
		Amount:      10,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	mockRepo.On("DeleteById", ctx, "1").Return(deletedProduct, nil)

	id, err := gateway.DeleteById(ctx, "1")
	assert.NoError(t, err)
	assert.Equal(t, "1", id)
	mockRepo.AssertExpectations(t)
}

// func TestProductGateway_UpdateById(t *testing.T) {
// 	mockRepo := new(MockProductRepository)
// 	gateway := NewProductGateway(mockRepo)
// 	ctx := context.Background()
// 	now := time.Now()
// 	oldProduct := &model.Product{
// 		ID:          "1",
// 		Name:        "Old",
// 		Description: "OldDesc",
// 		CategoryId:  2,
// 		Amount:      5,
// 		CreatedAt:   now,
// 		UpdatedAt:   now,
// 	}
// 	command := dto.UpdateProduct{
// 		ProductId:   "1",
// 		Name:        "New",
// 		Description: "",
// 		CategoryId:  0,
// 		Amount:      0,
// 	}
// 	mockRepo.On("FindOne", ctx, "1").Return(oldProduct, nil)
// 	mockRepo.On("UpdateById", ctx, mock.AnythingOfType("*model.Product")).Return(nil)

// 	product, err := gateway.UpdateById(ctx, command)
// 	assert.NoError(t, err)
// 	assert.Equal(t, "New", product.Name)
// 	assert.Equal(t, "OldDesc", product.Description)
// 	assert.Equal(t, 2, product.CategoryId)
// 	assert.Equal(t, 5, product.Amount)
// 	mockRepo.AssertExpectations(t)
// }

func TestProductGateway_UpdateById_NotFound(t *testing.T) {
	mockRepo := new(MockProductRepository)
	gateway := NewProductGateway(mockRepo)
	ctx := context.Background()
	command := dto.UpdateProduct{ProductId: "1"}
	mockRepo.On("FindOne", ctx, "1").Return((*model.Product)(nil), errors.New("not found"))

	product, err := gateway.UpdateById(ctx, command)
	assert.Error(t, err)
	assert.Equal(t, entity.Product{}, product)
	mockRepo.AssertExpectations(t)
}

func TestProductGateway_FindOne(t *testing.T) {
	mockRepo := new(MockProductRepository)
	gateway := NewProductGateway(mockRepo)
	ctx := context.Background()
	now := time.Now()
	productModel := &model.Product{
		ID:          "1",
		Name:        "Test",
		Description: "Desc",
		CategoryId:  2,
		Amount:      10,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	mockRepo.On("FindOne", ctx, "1").Return(productModel, nil)

	product, err := gateway.FindOne(ctx, "1")
	assert.NoError(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, "Test", product.Name)
	mockRepo.AssertExpectations(t)
}

func TestProductGateway_FindOne_NotFound(t *testing.T) {
	mockRepo := new(MockProductRepository)
	gateway := NewProductGateway(mockRepo)
	ctx := context.Background()
	mockRepo.On("FindOne", ctx, "1").Return((*model.Product)(nil), errors.New("not found"))

	product, err := gateway.FindOne(ctx, "1")
	assert.Error(t, err)
	assert.Nil(t, product)
	mockRepo.AssertExpectations(t)
}
