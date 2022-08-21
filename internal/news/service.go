package news

import (
	"context"
	"errors"
	"fmt"
	"log"
)

type Service interface {
	CreateNews(ctx context.Context, news []NewsDTO) ([]string, error)
	GetNewsByID(ctx context.Context, id string) (News, error)
	GetAllNewsWithPagination(ctx context.Context, limit uint64, lastID string) ([]News, error)
	UpdateNews(ctx context.Context, id string, news NewsDTO) error
	DeleteNews(ctx context.Context, id string) error
}

var _ Service = &service{}

type service struct {
	repository Repository
	logger     *log.Logger
}

func NewService(repository Repository, logger *log.Logger) Service {
	return &service{
		repository: repository,
		logger:     logger,
	}
}

func (s *service) CreateNews(ctx context.Context, news []NewsDTO) ([]string, error) {
	if len(news) == 0 {
		return []string{}, ErrDataNotValid{Param: "news", Message: "length news is zero (0)"}
	}

	param := make([]News, 0)
	for i := 0; i < len(news); i++ {
		if news[i].IsEmpty() {
			param = append(param, News{
				Title:   news[i].Title,
				Content: news[i].Content,
			})
		}
	}

	ids, err := s.repository.CreateMany(ctx, param)
	if err != nil {
		return nil, fmt.Errorf("couldn't create news w error %w", err)
	}

	if len(ids) == 0 {
		return nil, ErrNotCreate
	}

	return ids, nil
}

func (s *service) GetNewsByID(ctx context.Context, id string) (News, error) {
	if id == "" {
		return News{}, ErrDataNotValid{Param: "id", Message: "id is empty"}
	}

	news, err := s.repository.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return News{}, err
		}

		return News{}, fmt.Errorf("couldn't get news w error %w", err)
	}

	return news, nil
}

func (s *service) GetAllNewsWithPagination(ctx context.Context, limit uint64, lastID string) ([]News, error) {
	news, err := s.repository.GetAllWithPagination(ctx, limit, lastID)
	if err != nil {
		return nil, fmt.Errorf("couldn't get all news w error %w", err)
	}

	if len(news) == 0 {
		return nil, ErrNotFound
	}

	return news, nil
}

func (s *service) UpdateNews(ctx context.Context, id string, n NewsDTO) error {
	if !n.IsEmpty() && id == "" {
		return ErrDataNotValid{Param: "news", Message: "field news is empty"}
	}

	news := News{
		ID:      id,
		Title:   n.Title,
		Content: n.Content,
	}
	err := s.repository.Update(ctx, news)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return err
		}
		return fmt.Errorf("couldn't update news w error %w", err)
	}

	return nil
}

func (s *service) DeleteNews(ctx context.Context, id string) error {
	if id == "" {
		return ErrDataNotValid{Param: "id", Message: "id is empty"}
	}

	err := s.repository.DeleteByID(ctx, id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return err
		}
		return fmt.Errorf("couldn't delete news w error %w", err)
	}

	return nil
}
