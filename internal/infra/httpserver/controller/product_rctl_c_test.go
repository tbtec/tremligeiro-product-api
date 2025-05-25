package controller

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/container"
	"github.com/tbtec/tremligeiro/internal/infra/database/model"
	"github.com/tbtec/tremligeiro/internal/infra/httpserver"
)

// Mock compatível com a interface IProductRepository
type mockProductRepo struct{}

func (m *mockProductRepo) Create(ctx context.Context, p *model.Product) error { return nil }
func (m *mockProductRepo) FindOne(ctx context.Context, id string) (*model.Product, error) {
	return nil, nil
}
func (m *mockProductRepo) FindByCategory(ctx context.Context, categoryId int) (*[]model.Product, error) {
	return nil, nil
}
func (m *mockProductRepo) UpdateById(ctx context.Context, p *model.Product) error { return nil }
func (m *mockProductRepo) DeleteById(ctx context.Context, id string) (*model.Product, error) {
	return nil, nil
}

// Mock compatível com a interface ICategoryRepository
type mockCategoryRepo struct{}

func (m *mockCategoryRepo) FindById(id int) *entity.Category {
	return &entity.Category{ID: id, Name: "Mock"}
}

func newMockContainer() *container.Container {
	return &container.Container{
		ProductRepository:  &mockProductRepo{},
		CategoryRepository: &mockCategoryRepo{},
	}
}

func TestNewProductCreateRestController(t *testing.T) {
	container := newMockContainer()
	ctrl := NewProductCreateRestController(container)
	assert.NotNil(t, ctrl)
}

func TestProductCreateRestController_Handle(t *testing.T) {
	container := newMockContainer()
	ctrl := NewProductCreateRestController(container)

	input := dto.CreateProduct{
		Name:        "Test Product",
		Description: "A sample",
		CategoryId:  1,
		Amount:      100.0,
	}

	inputBytes, _ := json.Marshal(input)
	req := httpserver.Request{Body: inputBytes}

	resp := ctrl.Handle(context.Background(), req)

	assert.Equal(t, 200, resp.Code)
	var output struct {
		Status string `json:"status"`
	}
	bodyBytes, err := json.Marshal(resp.Body)
	assert.NoError(t, err)
	err = json.Unmarshal(bodyBytes, &output)
	assert.NoError(t, err)
	assert.Equal(t, "", output.Status)
}
