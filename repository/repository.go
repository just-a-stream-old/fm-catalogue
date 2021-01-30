package repository

import (
	"fmt"
	"gitlab.com/open-source-keir/financial-modelling/fm-catalogue/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type FMRepository interface {
	Do()
}

// fMRepository is a repository for handling financial-modelling data.
type fMRepository struct {
	logger *zap.Logger
	dbClient *mongo.Client
	dbName string
}

func NewFMRepository(cfg *config.Repository, logger *zap.Logger, dbClient *mongo.Client) (*fMRepository, error) {
	if logger == nil {
		return nil, fmt.Errorf("logger is required to construct the repository, and is nil")
	}
	if dbClient == nil {
		return nil, fmt.Errorf("dbClient is required to construct the repository, and is nil")
	}
	if cfg.DBName == "" {
		return nil, fmt.Errorf("dbName is required to construct the repository, and is empty")
	}

	fMRepository := &fMRepository{
		logger: logger,
		dbClient: dbClient,
		dbName: cfg.DBName,
	}
	return fMRepository, nil
}

func (*fMRepository) Do() {

}
