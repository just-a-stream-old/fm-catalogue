package main

import (
	"gitlab.com/open-source-keir/financial-modelling/fm-catalogue/api"
	"gitlab.com/open-source-keir/financial-modelling/fm-catalogue/config"
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

	server, err := api.NewServer(*cfg, logger)
	if err != nil {
		// Todo: log here?
		return err
	}

	server.Run()
	return nil
}
