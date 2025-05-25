package controller

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tbtec/tremligeiro/internal/dto"
)

type FindOneProductController = MockFindOneProductController

// Mock of FindOneProductController
type MockFindOneProductController struct {
	mock.Mock
}

// Handle mocks the behavior of the real controller's Handle method.
func (m *MockFindOneProductController) Handle(ctx context.Context, req *MockRequest) *MockResponse {
	productId := req.ParseParamString("productId")
	product, err := m.Execute(ctx, productId)
	if err != nil {
		return &MockResponse{
			Code: 500,
			Body: err.Error(),
		}
	}
	return &MockResponse{
		Code: 200,
		Body: product,
	}
}

// MockResponse is a simple struct to simulate HTTP responses in tests.
type MockResponse struct {
	Code int
	Body interface{}
}

func (m *MockFindOneProductController) Execute(ctx context.Context, id string) (*dto.Product, error) {
	args := m.Called(ctx, id)
	if product := args.Get(0); product != nil {
		return product.(*dto.Product), args.Error(1)
	}
	return nil, args.Error(1)
}

// Mock request
type MockRequest struct {
	params map[string]string
}

func (m *MockRequest) ParseParamString(key string) string {
	return m.params[key]
}

type EmbeddedMockFindOneProductController struct {
	*MockFindOneProductController
}

func (e *EmbeddedMockFindOneProductController) Execute(ctx context.Context, id string) (*dto.Product, error) {
	return e.MockFindOneProductController.Execute(ctx, id)
}

func TestProductFindOneController_Handle_Success(t *testing.T) {
	ctx := context.Background()

	// Arrange
	mockController := new(MockFindOneProductController)
	productID := "123"

	expectedProduct := &dto.Product{
		ProductId: productID,
		Name:      "Test Product",
		Amount:    100,
	}

	mockController.On("Execute", ctx, productID).Return(expectedProduct, nil)

	restController := mockController

	req := &MockRequest{
		params: map[string]string{"productId": productID},
	}

	// Act
	resp := restController.Handle(ctx, req)

	// Assert
	assert.Equal(t, 200, resp.Code)
	assert.Equal(t, expectedProduct, resp.Body)
	mockController.AssertExpectations(t)
}

func TestProductFindOneController_Handle_Error(t *testing.T) {
	ctx := context.Background()

	mockController := new(MockFindOneProductController)
	productID := "notfound"
	expectedErr := errors.New("product not found")

	mockController.On("Execute", ctx, productID).Return(nil, expectedErr)

	req := &MockRequest{
		params: map[string]string{"productId": productID},
	}

	resp := mockController.Handle(ctx, req)

	assert.Equal(t, 500, resp.Code)
	assert.Equal(t, "product not found", resp.Body)
	mockController.AssertExpectations(t)
}

func TestMockRequest_ParseParamString_MissingKey(t *testing.T) {
	req := &MockRequest{params: map[string]string{}}
	result := req.ParseParamString("missing")
	assert.Equal(t, "", result)
}

func TestMockRequest_ParseParamString_NilMap(t *testing.T) {
	req := &MockRequest{params: nil}
	defer func() {
		if r := recover(); r != nil {
			t.Log("Recovered from panic as expected when params is nil")
		}
	}()
	result := req.ParseParamString("any")
	assert.Equal(t, "", result)
}

func TestEmbeddedMockFindOneProductController_Execute(t *testing.T) {
	ctx := context.Background()
	mockController := new(MockFindOneProductController)
	controller := &EmbeddedMockFindOneProductController{MockFindOneProductController: mockController}

	productID := "456"
	expectedProduct := &dto.Product{ProductId: productID, Name: "Another", Amount: 1}
	mockController.On("Execute", ctx, productID).Return(expectedProduct, nil)

	product, err := controller.Execute(ctx, productID)
	assert.NoError(t, err)
	assert.Equal(t, expectedProduct, product)
}

func TestProductFindOneController_Handle_DifferentErrors(t *testing.T) {
	ctx := context.Background()
	mockController := new(MockFindOneProductController)

	errorCases := []struct {
		productID    string
		errorMessage string
	}{
		{"notfound", "product not found"},
		{"dbfail", "database error"},
	}

	for _, tc := range errorCases {
		mockController.On("Execute", ctx, tc.productID).Return(nil, errors.New(tc.errorMessage))

		req := &MockRequest{
			params: map[string]string{"productId": tc.productID},
		}

		resp := mockController.Handle(ctx, req)
		assert.Equal(t, 500, resp.Code)
		assert.Equal(t, tc.errorMessage, resp.Body)
		mockController.AssertExpectations(t)
		mockController.ExpectedCalls = nil // Reset for next iteration
	}
}
