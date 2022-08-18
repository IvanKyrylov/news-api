package news

import "log"

type Service interface {
}

var _ Service = &service{}

type service struct {
	storage Storage
	logger  *log.Logger
}

func NewService(userStorage Storage, logger *log.Logger) (Service, error) {
	return &service{
		storage: userStorage,
		logger:  logger,
	}, nil
}
