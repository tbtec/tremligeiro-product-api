// go
package test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tbtec/tremligeiro/internal/infra/container"
)

// Mock ProductRepository for deletion
type mockProductRepo struct {
	DeleteByIdFunc func(ctx context.Context, id string) error
}

func (m *mockProductRepo) DeleteById(ctx context.Context, id string) error {
	return m.DeleteByIdFunc(ctx, id)
}

// Implement other methods if required by interface, but leave them empty for this test

func TestDeleteProductController_Execute_Success(t *testing.T) {
	ctx := context.Background()

	productRepo := &mockProductRepo{
		DeleteByIdFunc: func(ctx context.Context, id string) error {
			if id == "prod1" {
				return nil
			}
			return errors.New("not found")
		},
	}

	testContainer := &container.Container{
		ProductRepository: productRepo,
	}

	controller := NewDeleteProductController(testContainer)

	id := "prod1"
	result, err := controller.Execute(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, id, result)
}

func TestDeleteProductController_Execute_NotFound(t *testing.T) {
	ctx := context.Background()

	productRepo := &mockProductRepo{
		DeleteByIdFunc: func(ctx context.Context, id string) error {
			return errors.New("not found")
		},
	}

	testContainer := &container.Container{
		ProductRepository: productRepo,
	}

	controller := NewDeleteProductController(testContainer)

	id := "invalid-id"
	result, err := controller.Execute(ctx, id)
	assert.Error(t, err)
	assert.Empty(t, result)
}
