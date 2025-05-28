package controller

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tbtec/tremligeiro/internal/infra/container"
	"github.com/tbtec/tremligeiro/internal/infra/database/model"
	"github.com/tbtec/tremligeiro/internal/infra/httpserver"
	"github.com/tbtec/tremligeiro/test/repository"
)

func TestProductDeleteController_Handle_NoContent(t *testing.T) {
	container := &container.Container{
		ProductRepository: &repository.MockProductRepo{
			ExecuteFunc: func(ctx context.Context, productId string) (string, error) {
				return "", nil
			},
		},
	}
	ctrl := NewProductDeleteByIdRestController(container)

	req := httpserver.Request{Params: map[string]string{"productId": "123"}}

	resp := ctrl.Handle(context.Background(), req)
	assert.Equal(t, httpserver.NoContent(), resp)
}

func TestProductDeleteController_Handle_NotFound(t *testing.T) {
	container := &container.Container{
		ProductRepository: &repository.MockProductRepo{
			ExecuteFunc: func(ctx context.Context, productId string) (string, error) {
				return "", assert.AnError
			},
		},
	}
	ctrl := NewProductDeleteByIdRestController(container)
	req := httpserver.Request{Params: map[string]string{"productId": "123"}}

	resp := ctrl.Handle(context.Background(), req)
	assert.Equal(t, 204, resp.Code)
}

func TestProductDeleteController_Handle_RecordNotFound(t *testing.T) {
	container := &container.Container{
		ProductRepository: &repository.MockProductRepo{
			ExecuteFunc: func(ctx context.Context, productId string) (string, error) {
				return "", assert.AnError
			},
			DeleteByIdFunc: func(ctx context.Context, id string) (*model.Product, error) {
				return nil, errors.New("record not found")
			},
		},
	}
	ctrl := NewProductDeleteByIdRestController(container)
	req := httpserver.Request{Params: map[string]string{"productId": ""}}

	resp := ctrl.Handle(context.Background(), req)

	assert.Equal(t, httpserver.NotFound(errors.New("record not found")), resp)
}

func TestProductDeleteController_Handle_UnprocessableEntity(t *testing.T) {
	container := &container.Container{
		ProductRepository: &repository.MockProductRepo{
			ExecuteFunc: func(ctx context.Context, productId string) (string, error) {
				return "", assert.AnError
			},
			DeleteByIdFunc: func(ctx context.Context, id string) (*model.Product, error) {
				return nil, errors.New("Unprocessable")
			},
		},
	}
	ctrl := NewProductDeleteByIdRestController(container)
	req := httpserver.Request{Params: map[string]string{"productId": ""}}

	resp := ctrl.Handle(context.Background(), req)

	assert.Equal(t, 422, resp.Code)
}
