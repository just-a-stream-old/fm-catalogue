package repository

import (
	"gitlab.com/open-source-keir/financial-modelling/fm-catalogue/config"
	"go.uber.org/zap"
)

type FMRepository interface {
	Do()
}

type fMRepository struct {
	logger *zap.Logger
	db string
}

func NewFMRepository(cfg *config.Repository, logger *zap.Logger) (*fMRepository, error) {
	// Setup database


	// Setup database connection
	//connectionStr := fmt.Sprintf("")

	fMRepository := &fMRepository{
		logger: logger,
		db:     "db connection / cursor would go here",
	}
	return fMRepository, nil
}

func (*fMRepository) Do() {

}
