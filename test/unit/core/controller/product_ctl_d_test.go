package controller

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tbtec/tremligeiro/internal/core/controller"
	"github.com/tbtec/tremligeiro/internal/infra/container"
	"github.com/tbtec/tremligeiro/internal/infra/database/model"
	"github.com/tbtec/tremligeiro/test/fixtures"
)

func TestDeleteProductController_Execute_Success(t *testing.T) {
	ctx := context.Background()

	productRepo := &fixtures.MockProductRepo{
		DeleteByIdFunc: func(ctx context.Context, id string) (*model.Product, error) {
			if id == "prod1" {
				return &model.Product{ID: "prod1"}, nil
			}
			return nil, errors.New("not found")
		},
	}

	testContainer := &container.Container{
		ProductRepository: productRepo,
	}

	controller := controller.NewDeleteProductController(testContainer)

	id := "prod1"
	result, err := controller.Execute(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, id, result)
}

func TestDeleteProductController_Execute_NotFound(t *testing.T) {
	ctx := context.Background()

	productRepo := &fixtures.MockProductRepo{
		DeleteByIdFunc: func(ctx context.Context, id string) (*model.Product, error) {
			return nil, errors.New("not found")
		},
	}

	testContainer := &container.Container{
		ProductRepository: productRepo,
	}

	controller := controller.NewDeleteProductController(testContainer)

	id := ""
	result, err := controller.Execute(ctx, id)
	assert.Error(t, err)
	assert.Empty(t, result)
}
