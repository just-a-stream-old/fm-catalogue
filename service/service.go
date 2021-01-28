package service

type FMService interface {
	Do()
}

type fMService struct {

}

func (fms *fMService) Do() {
}

func NewFMService() *fMService {
	return &fMService{

	}
}
