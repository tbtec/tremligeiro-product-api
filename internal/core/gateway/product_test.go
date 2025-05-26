package gateway

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/infra/database/model"
	"github.com/tbtec/tremligeiro/test/repository"
)

func TestProductGateway_FindByCategory_Success(t *testing.T) {
	now := time.Now()
	mockProducts := []model.Product{
		{
			ID:          "1",
			Name:        "Product 1",
			Description: "Desc 1",
			Amount:      10,
			CategoryId:  2,
			CreatedAt:   now,
			UpdatedAt:   now,
		},
		{
			ID:          "2",
			Name:        "Product 2",
			Description: "Desc 2",
			Amount:      20,
			CategoryId:  2,
			CreatedAt:   now,
			UpdatedAt:   now,
		},
	}
	mockRepo := &repository.MockProductRepo{
		FindByCategoryFunc: func(ctx context.Context, id int) (*[]model.Product, error) {
			return &mockProducts, nil
		},
	}
	gateway := &ProductGateway{productRepository: mockRepo}

	got, err := gateway.FindByCategory(context.Background(), 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []entity.Product{
		{
			ID:          "1",
			Name:        "Product 1",
			Description: "Desc 1",
			Amount:      10,
			CategoryId:  2,
			CreatedAt:   now,
			UpdatedAt:   now,
		},
		{
			ID:          "2",
			Name:        "Product 2",
			Description: "Desc 2",
			Amount:      20,
			CategoryId:  2,
			CreatedAt:   now,
			UpdatedAt:   now,
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FindByCategory() = %v, want %v", got, want)
	}
}

func TestProductGateway_FindByCategory_RepositoryError(t *testing.T) {
	mockRepo := &repository.MockProductRepo{
		FindByCategoryFunc: func(ctx context.Context, id int) (*[]model.Product, error) {
			return nil, errors.New("db error")
		},
	}
	gateway := &ProductGateway{productRepository: mockRepo}

	got, err := gateway.FindByCategory(context.Background(), 99)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if got != nil {
		t.Errorf("expected nil products, got %v", got)
	}
}

func TestProductGateway_FindByCategory_EmptyResult(t *testing.T) {
	mockProducts := []model.Product{}
	mockRepo := &repository.MockProductRepo{
		FindByCategoryFunc: func(ctx context.Context, id int) (*[]model.Product, error) {
			return &mockProducts, nil
		},
	}
	gateway := &ProductGateway{productRepository: mockRepo}

	got, err := gateway.FindByCategory(context.Background(), 123)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 0 {
		t.Errorf("expected empty slice, got %v", got)
	}
}
