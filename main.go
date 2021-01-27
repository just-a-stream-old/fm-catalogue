package main

import (
	"gitlab.com/open-source-keir/financial-modelling/fm-catalogue/api"
	"gitlab.com/open-source-keir/financial-modelling/fm-catalogue/config"
	"gitlab.com/open-source-keir/financial-modelling/fm-catalogue/service"
	"go.uber.org/zap"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	cfg, err := config.GetConfig()
	if err != nil {
		return err
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		return err
	}
	defer logger.Sync()

	fmService, err := service.NewFmService()

	server, err := api.NewServer(&api.Config{

	})
	if err != nil {
		// Todo: log here?
		return err
	}

	server.Run()
	return nil
}
