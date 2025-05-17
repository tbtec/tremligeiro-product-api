package container

import (
	"context"
	"log"

	"github.com/tbtec/tremligeiro/internal/env"
	"github.com/tbtec/tremligeiro/internal/infra/database/mongodb"
	"github.com/tbtec/tremligeiro/internal/infra/database/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type Container struct {
	Config             env.Config
	TremLigeiroDB      *mongo.Collection
	ProductRepository  repository.IProductRepository
	CategoryRepository repository.ICategoryRepository
}

func New(config env.Config) (*Container, error) {
	factory := Container{}
	factory.Config = config

	return &factory, nil
}

func (container *Container) Start() error {

	err := mongodb.Migrate(getMongoDBConf(container.Config))
	if err != nil {
		log.Fatalf("Erro ao conectar ao MongoDB: %v", err)
	}

	container.TremLigeiroDB, err = mongodb.New(getMongoDBConf(container.Config))
	if err != nil {
		log.Fatalf("Erro ao conectar ao MongoDB: %v", err)
	}

	container.ProductRepository = repository.NewProductRepository(container.TremLigeiroDB)
	container.CategoryRepository = repository.NewCategoryRepository()

	log.Println("Banco de dados:", container.TremLigeiroDB.Name())

	return nil
}

func (container *Container) Stop() error {
	db := container.TremLigeiroDB

	defer db.Database().Client().Disconnect(context.Background())
	return nil
}

func getMongoDBConf(config env.Config) mongodb.MongoConf {
	return mongodb.MongoConf{
		User:           config.DbUser,
		Pass:           config.DbPassword,
		Url:            config.DbHost,
		Port:           config.DbPort,
		DbName:         config.DbName,
		CollectionName: config.CollectionName,
	}
}
