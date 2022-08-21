package news

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("news not found")
var ErrNotCreate = errors.New("news not create")

type ErrDataNotValid struct {
	Message string
	Param   string
}

func (err ErrDataNotValid) Error() string {
	return fmt.Sprintf("couldn't validate data by param %v w error %v", err.Param, err.Message)
}
