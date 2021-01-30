package repository

import (
	"context"
	"fmt"
	"gitlab.com/open-source-keir/financial-modelling/fm-catalogue/config"
	"gitlab.com/open-source-keir/financial-modelling/fm-catalogue/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"time"
)

type FMRepository interface {
	GetExchanges(ctx context.Context) ([]model.Exchange, int, error)
}

// fMRepository is a repository for handling financial-modelling data.
type fMRepository struct {
	logger *zap.Logger
	db *mongo.Database
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
		db: dbClient.Database(cfg.DBName),
		dbName: cfg.DBName,
	}
	return fMRepository, nil
}

// Todo: Define general find method that handles cursor creation etc

func (fmr *fMRepository) GetExchanges(ctx context.Context) ([]model.Exchange, int, error) {
	c := fmr.db.Collection("exchanges")
	dbCtx, _ := context.WithTimeout(context.Background(), 5*time.Second) // Todo: Clean up with cancel

	filter := bson.M{}
	cursor, err := c.Find(ctx, filter)
	if err != nil {
		fmr.logger.Error(err.Error())
		return nil, 0, err
	}
	defer cursor.Close(ctx) // Todo: Clean up here with error handling

	var result []model.Exchange
	for cursor.Next(dbCtx) {
		exchange := &model.Exchange{}
		if err := cursor.Decode(exchange); err != nil {
			fmr.logger.Error(err.Error())
			return nil, 0, err
		}
		result = append(result, *exchange)
	}

	return result, len(result), nil
}
