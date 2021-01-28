package service

import (
	"gitlab.com/open-source-keir/financial-modelling/fm-catalogue/repository"
	"go.uber.org/zap"
)

type FMService interface {
	Do()
}

type fMService struct {
	logger *zap.Logger
	fMRepository repository.FMRepository
}

func (fms *fMService) Do() {
}

func NewFMService(logger *zap.Logger, fMRepository repository.FMRepository) *fMService {
	return &fMService{
		logger: logger,
		fMRepository: fMRepository,
	}
}
