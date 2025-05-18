package gateway

import (
	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/infra/database/repository"
)

type CategoryGateway struct {
	categoryRepository repository.ICategoryRepository
}

func NewCategoryGateway(categoryRepository repository.ICategoryRepository) *CategoryGateway {

	return &CategoryGateway{
		categoryRepository: categoryRepository,
	}
}

func (gw *CategoryGateway) FindById(id int) *entity.Category {

	return gw.categoryRepository.FindById(id)
}
