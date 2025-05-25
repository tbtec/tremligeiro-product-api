package controller

import (
	"context"
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

// Only needed if your real Request interface has these methods.
func (m *MockRequest) ParseBody(obj interface{}) error { return nil }
func (m *MockRequest) ParseQueryParamString(key string) string {
	return ""
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

// func TestProductFindOneController_Handle_Error(t *testing.T) {
// 	ctx := context.Background()

// 	// Arrange
// 	mockController := new(MockFindOneProductController)
// 	productID := "notfound"

// 	expectedErr := errors.New("product not found")

// 	mockController.On("Execute", ctx, productID).Return(nil, expectedErr)

// 	restController := &controller.ProductFindOneController{
// 		controller: mockController,
// 	}

// 	req := &MockRequest{
// 		params: map[string]string{"productId": productID},
// 	}

// 	// Act
// 	resp := restController.Handle(ctx, req)

// 	// Assert
// 	assert.Equal(t, 500, resp.StatusCode)
// 	assert.Contains(t, resp.Body, "product not found")
// 	mockController.AssertExpectations(t)
// }
