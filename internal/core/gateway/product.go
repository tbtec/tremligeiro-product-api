package gateway

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/database/model"
	"github.com/tbtec/tremligeiro/internal/infra/database/repository"
)

type ProductGateway struct {
	productRepository  repository.IProductRepository
	categoryRepository repository.ICategoryRepository
}

func NewProductGateway(productRepository repository.IProductRepository) *ProductGateway {
	return &ProductGateway{
		productRepository: productRepository,
	}
}

func (gtw *ProductGateway) Create(ctx context.Context, product *entity.Product) error {

	productModel := model.Product{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		CategoryId:  product.CategoryId,
		Amount:      product.Amount,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}

	err := gtw.productRepository.Create(ctx, &productModel)

	if err != nil {
		return err
	}

	return nil
}

func (gtw *ProductGateway) FindByCategory(ctx context.Context, id int) ([]entity.Product, error) {
	productModels, err := gtw.productRepository.FindByCategory(ctx, id)
	if err != nil {
		return nil, err
	}

	products := []entity.Product{}

	for _, productModel := range *productModels {
		product := entity.Product{
			ID:          productModel.ID,
			Name:        productModel.Name,
			Description: productModel.Description,
			Amount:      productModel.Amount,
			CategoryId:  productModel.CategoryId,
			CreatedAt:   productModel.CreatedAt,
			UpdatedAt:   productModel.UpdatedAt,
		}
		products = append(products, product)
	}

	return products, nil
}

func (gtw *ProductGateway) DeleteById(ctx context.Context, id string) (string, error) {

	_, err := gtw.productRepository.DeleteById(ctx, id)

	return id, err
}

func (gtw *ProductGateway) UpdateById(ctx context.Context, command dto.UpdateProduct) (entity.Product, error) {

	old_product, errProduct := gtw.productRepository.FindOne(ctx, command.ProductId)

	if old_product == nil {
		return entity.Product{}, errProduct
	}

	if errProduct != nil {
		return entity.Product{}, errProduct
	}

	buildUpdateProduct(&command, old_product)

	new_product := model.Product{
		ID:          command.ProductId,
		Name:        command.Name,
		Description: command.Description,
		CategoryId:  command.CategoryId,
		Amount:      command.Amount,
		CreatedAt:   command.CreatedAt,
	}

	err := gtw.productRepository.UpdateById(ctx, &new_product)
	if err != nil {
		return entity.Product{}, err
	}

	output := entity.Product{
		ID:          new_product.ID,
		Name:        new_product.Name,
		Description: new_product.Description,
		CategoryId:  new_product.CategoryId,
		Amount:      new_product.Amount,
		CreatedAt:   new_product.CreatedAt,
		UpdatedAt:   new_product.UpdatedAt,
	}

	return output, nil
}

func buildUpdateProduct(command *dto.UpdateProduct, old_product *model.Product) error {

	if command.Name == "" {
		command.Name = old_product.Name
	}
	if command.Description == "" {
		command.Description = old_product.Description
	}
	if command.CategoryId == 0 {
		command.CategoryId = old_product.CategoryId
	}
	if command.Amount == 0 {
		command.Amount = old_product.Amount
	}
	command.CreatedAt = old_product.CreatedAt

	return nil

}

func (gtw *ProductGateway) FindOne(ctx context.Context, id string) (*entity.Product, error) {

	productModel, err := gtw.productRepository.FindOne(ctx, id)
	if productModel == nil {
		return nil, err
	}

	product := entity.Product{
		ID:          productModel.ID,
		Name:        productModel.Name,
		Description: productModel.Description,
		Amount:      productModel.Amount,
		CategoryId:  productModel.CategoryId,
		CreatedAt:   productModel.CreatedAt,
		UpdatedAt:   productModel.UpdatedAt,
	}

	return &product, nil
}
