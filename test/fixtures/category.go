package fixtures

import "github.com/tbtec/tremligeiro/internal/core/domain/entity"

type MockCategoryRepo struct {
	FindByIdFunc func(id int) *entity.Category
}

func (m *MockCategoryRepo) FindById(id int) *entity.Category {
	return m.FindByIdFunc(id)
}
