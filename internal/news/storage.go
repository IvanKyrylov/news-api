package news

import (
	"errors"
)

type Repository interface {
}

var (
	ErrTitleEmpty = errors.New("fail title is empty")
	ErrNotFound   = errors.New("news not found")
)
