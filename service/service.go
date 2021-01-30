package service

import (
	"context"
	"gitlab.com/open-source-keir/financial-modelling/fm-catalogue/config"
	"gitlab.com/open-source-keir/financial-modelling/fm-catalogue/model"
	"gitlab.com/open-source-keir/financial-modelling/fm-catalogue/repository"
	"go.uber.org/zap"
)

type FMService interface {
	GetExchanges(context.Context) ([]model.Exchange, int, error)
}

type fMService struct {
	logger *zap.Logger
	fMRepository repository.FMRepository
}

func (fms *fMService) GetExchanges(ctx context.Context) ([]model.Exchange, int, error) {
	exchanges, count, err := fms.fMRepository.GetExchanges(ctx)
	if err != nil {
		return nil, count, err
	}

	return exchanges, count, err
}

func NewFMService(cfg *config.Service, logger *zap.Logger, fMRepository repository.FMRepository) *fMService {
	return &fMService{
		logger: logger,
		fMRepository: fMRepository,
	}
}
