package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/tbtec/tremligeiro/internal/infra/database/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IProductRepository interface {
	Create(ctx context.Context, product *model.Product) error
	FindOne(ctx context.Context, id string) (*model.Product, error)
	FindByCategory(ctx context.Context, id int) (*[]model.Product, error)
	DeleteById(ctx context.Context, id string) (*model.Product, error)
	UpdateById(ctx context.Context, product *model.Product) error
}

type ProductRepository struct {
	database *mongo.Collection
}

func NewProductRepository(database *mongo.Collection) IProductRepository {
	return &ProductRepository{
		database: database,
	}
}

func (repository *ProductRepository) Create(ctx context.Context, product *model.Product) error {

	//result := repository.database.DB.WithContext(ctx).Create(&product)
	result, err := repository.database.InsertOne(context.Background(), &product)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("result: ", result)

	return nil
}

func (repository *ProductRepository) FindOne(ctx context.Context, id string) (*model.Product, error) {
	product := &model.Product{}

	err := repository.database.FindOne(ctx, bson.M{"id": id}).Decode(&product)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return product, nil
}

func (repository *ProductRepository) FindByCategory(ctx context.Context, id int) (*[]model.Product, error) {
	product := []model.Product{}

	cursor, err := repository.database.Find(ctx, bson.M{"categoryid": id})
	if err != nil {
		//TODO
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &product); err != nil {
		//TODO
	}

	fmt.Printf(cursor.Current.String())

	return &product, nil
}

func (repository *ProductRepository) DeleteById(ctx context.Context, id string) (*model.Product, error) {
	product := &model.Product{
		ID: id,
	}

	err := repository.database.FindOneAndDelete(ctx, bson.M{"id": id})

	fmt.Printf(product.ID)
	if err != nil {
		return nil, err.Err()
	} /*else if err. RowsAffected < 1 {
		return product, fmt.Errorf("record not found")
	}*/

	return product, nil
}

func (repository *ProductRepository) UpdateById(ctx context.Context, product *model.Product) error {

	result := repository.database.FindOneAndUpdate(
		ctx,
		bson.M{"id": product.ID},
		bson.M{"$set": product})

	if result.Err() != nil {
		fmt.Printf(result.Err().Error())
		return result.Err()
	}

	return nil
}
