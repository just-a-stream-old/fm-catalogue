package service

import "go.uber.org/zap"

type FMService interface {
	Do()
}

type fMService struct {
	logger *zap.Logger
}

func (fms *fMService) Do() {
}

func NewFMService(logger *zap.Logger) *fMService {
	return &fMService{
		logger: logger,
	}
}
