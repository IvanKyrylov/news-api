package news

import (
	"context"
)

type Repository interface {
	CreateOne(ctx context.Context, n News) (string, error)

	CreateMany(ctx context.Context, news []News) ([]string, error)

	GetByID(ctx context.Context, id string) (News, error)

	GetAllWithPagination(ctx context.Context, limit uint64, lastID string) ([]News, error)

	Update(ctx context.Context, n News) error

	DeleteByID(ctx context.Context, id string) error
}
