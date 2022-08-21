package news

import "log"

type Service interface {
}

var _ Service = &service{}

type service struct {
	repository Repository
	logger     *log.Logger
}

func NewService(repository Repository, logger *log.Logger) (Service, error) {
	return &service{
		repository: repository,
		logger:     logger,
	}, nil
}
