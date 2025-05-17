package repository

import (
	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
)

var (
	Lanche = entity.Category{
		ID:   1,
		Name: "Lanche",
	}
	Acompanhamento = entity.Category{
		ID:   2,
		Name: "Acompanhamento",
	}
	Bebida = entity.Category{
		ID:   3,
		Name: "Bebida",
	}
	Sobremesa = entity.Category{
		ID:   4,
		Name: "Sobremesa",
	}
)

var categoryCatalog = []entity.Category{
	Lanche,
	Acompanhamento,
	Bebida,
	Sobremesa,
}

type ICategoryRepository interface {
	FindById(id int) *entity.Category
}

type CategoryRepository struct {
}

func NewCategoryRepository() ICategoryRepository {
	return &CategoryRepository{}
}

func (repository *CategoryRepository) FindById(id int) *entity.Category {
	for _, category := range categoryCatalog {
		if category.ID == id {
			return &category
		}
	}

	return nil
}
