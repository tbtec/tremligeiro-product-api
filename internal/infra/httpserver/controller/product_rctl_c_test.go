package controller

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/container"
	"github.com/tbtec/tremligeiro/internal/infra/httpserver"
	"github.com/tbtec/tremligeiro/test/repository"
)

func newMockContainer() *container.Container {
	return &container.Container{
		ProductRepository:  &repository.MockProductRepoInterface{},
		CategoryRepository: &repository.MockCategoryRepoInterface{},
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

func TestProductCreateRestController_Handle_InvalidJSON(t *testing.T) {
	container := newMockContainer()
	ctrl := NewProductCreateRestController(container)

	// Invalid JSON (not a valid dto.CreateProduct)
	req := httpserver.Request{Body: []byte(`{invalid json}`)}

	resp := ctrl.Handle(context.Background(), req)

	assert.NotEqual(t, 200, resp.Code)
}

func TestProductCreateRestController_Handle_MissingFields(t *testing.T) {
	container := newMockContainer()
	ctrl := NewProductCreateRestController(container)

	// Missing required fields
	input := map[string]interface{}{
		"Name": "", // Name is empty
	}
	inputBytes, _ := json.Marshal(input)
	req := httpserver.Request{Body: inputBytes}

	resp := ctrl.Handle(context.Background(), req)

	assert.NotEqual(t, 200, resp.Code)
}

func TestProductCreateRestController_Handle_CategoryNotFound(t *testing.T) {
	container := &container.Container{
		ProductRepository:  &repository.MockProductRepoInterface{},
		CategoryRepository: &repository.MockCategoryRepoNotFound{},
	}
	ctrl := NewProductCreateRestController(container)

	input := dto.CreateProduct{
		Name:        "Test Product",
		Description: "A sample",
		CategoryId:  999, // Non-existent category
		Amount:      100.0,
	}
	inputBytes, _ := json.Marshal(input)
	req := httpserver.Request{Body: inputBytes}

	resp := ctrl.Handle(context.Background(), req)
	assert.NotEqual(t, 200, resp.Code)
}
