package gateway

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
)

type MockCategoryRepository struct {
	mock.Mock
}

func (m *MockCategoryRepository) FindById(id int) *entity.Category {
	args := m.Called(id)
	if cat, ok := args.Get(0).(*entity.Category); ok {
		return cat
	}
	return nil
}

func TestNewCategoryGateway(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	gw := NewCategoryGateway(mockRepo)
	assert.NotNil(t, gw)
	assert.Equal(t, mockRepo, gw.categoryRepository)
}

func TestCategoryGateway_FindById(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	expectedCategory := &entity.Category{ID: 1, Name: "Test"}
	mockRepo.On("FindById", 1).Return(expectedCategory)

	gw := NewCategoryGateway(mockRepo)
	result := gw.FindById(1)

	assert.Equal(t, expectedCategory, result)
	mockRepo.AssertExpectations(t)
}

func TestCategoryGateway_FindById_NotFound(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	mockRepo.On("FindById", 2).Return(nil)

	gw := NewCategoryGateway(mockRepo)
	result := gw.FindById(2)

	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}
