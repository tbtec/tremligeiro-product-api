package controller

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/infra/container"
	"github.com/tbtec/tremligeiro/internal/infra/database/model"
	"github.com/tbtec/tremligeiro/internal/infra/httpserver"
	"github.com/tbtec/tremligeiro/test/repository"
)

// --- Mocks ---

type mockFindOneProductController struct {
	mock.Mock
}

func (m *mockFindOneProductController) Execute(ctx context.Context, productID string) (interface{}, error) {
	args := m.Called(ctx, productID)
	return args.Get(0), args.Error(1)
}

type mockRequest struct {
	mock.Mock
}

func (m *mockRequest) ParseParamString(key string) string {
	args := m.Called(key)
	return args.String(0)
}

// --- Test ---

func TestProductFindOneController_Handle_Success(t *testing.T) {

	mockProduct := &model.Product{
		ID:          "prod1",
		Name:        "Product 1",
		Description: "Description 1",
		Amount:      100.0,
		CategoryId:  1,
		CreatedAt:   time.Now(),
	}

	container := &container.Container{
		ProductRepository: &repository.MockProductRepo{
			FindOneFunc: func(ctx context.Context, id string) (*model.Product, error) {
				if id == "prod1" {
					return mockProduct, nil
				}
				return nil, errors.New("not found")
			},
		},

		CategoryRepository: &repository.MockCategoryRepo{
			FindByIdFunc: func(id int) *entity.Category {
				if id == 1 {
					return &entity.Category{ID: 1, Name: "Category 1"}
				}
				return nil
			},
		},
	}
	ctrl := NewProductFindOneRestController(container)

	req := httpserver.Request{
		Params: map[string]string{"productId": "prod1"},
	}
	resp := ctrl.Handle(context.Background(), req)

	assert.Equal(t, 200, resp.Code)
}
