package controller

import (
	"context"
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/infra/container"
	"github.com/tbtec/tremligeiro/internal/infra/database/model"
	"github.com/tbtec/tremligeiro/internal/infra/httpserver"
	"github.com/tbtec/tremligeiro/test/repository"
)

func TestHandle_Success(t *testing.T) {

	mockProductsSlice := []model.Product{
		{ID: "prod1", Name: "Product 1", CategoryId: 10},
		{ID: "prod2", Name: "Product 2", CategoryId: 10},
	}
	mockProducts := &mockProductsSlice

	container := &container.Container{
		ProductRepository: &repository.MockProductRepo{
			FindByCategoryFunc: func(ctx context.Context, categoryId int) (*[]model.Product, error) {
				if categoryId == 10 {

					return mockProducts, nil
				}

				return nil, errors.New("not found")
			},
		},

		CategoryRepository: &repository.MockCategoryRepo{
			FindByIdFunc: func(id int) *entity.Category {
				if id == 10 {
					return &entity.Category{ID: 10, Name: "Category 10"}
				}
				return nil
			},
		},
	}
	ctrl := NewProductFindByCategoryRestController(container)

	req := httpserver.Request{
		Query: map[string]string{"categoryId": "10"},
	}
	resp := ctrl.Handle(context.Background(), req)

	assert.Equal(t, 200, resp.Code)
}

func TestHandle_InvalidCategoryId(t *testing.T) {

	container := &container.Container{
		ProductRepository: &repository.MockProductRepo{
			FindByCategoryFunc: func(ctx context.Context, categoryId int) (*[]model.Product, error) {

				return nil, nil
			},
		},
	}
	ctrl := NewProductFindByCategoryRestController(container)

	req := httpserver.Request{
		Query: map[string]string{"categoryId": "abc"},
	}
	resp := ctrl.Handle(context.Background(), req)

	assert.Equal(t, httpserver.HandleError(context.Background(), &strconv.NumError{
		Func: "Atoi",
		Num:  "abc",
		Err:  errors.New("invalid syntax"),
	}).Code, resp.Code)
}

func TestHandle_ExecuteError(t *testing.T) {

	container := &container.Container{
		ProductRepository: &repository.MockProductRepo{
			FindByCategoryFunc: func(ctx context.Context, categoryId int) (*[]model.Product, error) {

				return nil, errors.New("record not found")
			},
		},

		CategoryRepository: &repository.MockCategoryRepo{
			FindByIdFunc: func(id int) *entity.Category {
				return nil
			},
		},
	}
	ctrl := NewProductFindByCategoryRestController(container)

	req := httpserver.Request{
		Query: map[string]string{"categoryId": "42"},
	}
	resp := ctrl.Handle(context.Background(), req)

	assert.Equal(t, 422, resp.Code)

}
