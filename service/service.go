package service

type service interface {

}

type Service struct {

}

func NewFmService() (service, error) {
	return Service{}, nil
}
