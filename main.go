package main

import (
	"context"
	"fmt"
	"gitlab.com/open-source-keir/financial-modelling/fm-catalogue/api"
	"gitlab.com/open-source-keir/financial-modelling/fm-catalogue/config"
	"gitlab.com/open-source-keir/financial-modelling/fm-catalogue/repository"
	"gitlab.com/open-source-keir/financial-modelling/fm-catalogue/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"time"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return err
	}
	defer logger.Sync()

	cfg, err := config.GetConfig()
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	dbClient, err := setupDB(&cfg.Repository)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	defer func() {
		if err := closeDB(dbClient); err != nil {
			logger.Error(err.Error())
		}
	}()

	fMRepository, err := repository.NewFMRepository(&cfg.Repository, logger, dbClient)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	fMService := service.NewFMService(&cfg.Service, logger, fMRepository)

	server, err := api.NewServer(&cfg.Server, logger, fMService)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	server.Run()

	return nil
}

func setupDB(cfg *config.Repository) (*mongo.Client, error) {
	// Cancel if database connection is not established within time provided
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	connStr := fmt.Sprintf("%s://%s:%s", cfg.DBDriver, cfg.DBHost, cfg.DBPort)
	dbClient, err := mongo.Connect(ctx, options.Client().ApplyURI(connStr))

	return dbClient, err
}

func closeDB(dbClient *mongo.Client) error {
	return dbClient.Disconnect(context.TODO())
}


