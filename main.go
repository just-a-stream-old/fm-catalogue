package main

import (
	"gitlab.com/open-source-keir/financial-modelling/fm-catalogue/api"
	"gitlab.com/open-source-keir/financial-modelling/fm-catalogue/config"
	"gitlab.com/open-source-keir/financial-modelling/fm-catalogue/repository"
	"gitlab.com/open-source-keir/financial-modelling/fm-catalogue/service"
	"go.uber.org/zap"
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

	//db, err := setupDB()
	//if err != nil {
	//	logger.Error(err.Error())
	//	return err
	//}
	//defer closeDB()

	fMRepository, err := repository.NewFMRepository(&cfg.Repository, logger)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	fMService := service.NewFMService(&cfg.Service, logger, fMRepository)

	server := api.NewServer(&cfg.Server, logger, fMService)

	server.Run()

	return nil
}
