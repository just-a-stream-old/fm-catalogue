package repository

import "go.uber.org/zap"

type FMRepository interface {
	Do()
}

type fMRepository struct {
	logger *zap.Logger
	db string
}

func NewFMRepository(logger *zap.Logger) (*fMRepository, error) {
	fMRepository := &fMRepository{
		logger: logger,
		db:     "db connection / cursor would go here",
	}
	return fMRepository, nil
}

func (*fMRepository) Do() {

}
