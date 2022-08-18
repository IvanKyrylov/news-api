package db

import (
	"database/sql"
	"log"

	"github.com/IvanKyrylov/news-api/internal/news"
)

var _ news.Storage = &db{}

type db struct {
	storage *sql.DB
	logger  *log.Logger
}

func NewStorage(storage *sql.DB, logger *log.Logger) news.Storage {
	return &db{
		storage: storage,
		logger:  logger,
	}
}
