package news

import (
	"context"
	"errors"
	"fmt"
	"log"
	"regexp"
)

type Service interface {
	CreateNews(ctx context.Context, news []NewsDTO) ([]string, error)
	GetNewsByID(ctx context.Context, id string) (NewsDTO, error)
	GetAllNewsWithPagination(ctx context.Context, limit uint64, pToken string) ([]NewsDTO, error)
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
		return nil, ErrDataNotValid{Param: "news", Message: "length news is zero (0)"}
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

	if len(param) == 0 {
		return nil, ErrDataNotValid{Param: "news", Message: "news is not valid"}
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

func (s *service) GetNewsByID(ctx context.Context, id string) (NewsDTO, error) {
	if !isValidID(id) {
		return NewsDTO{}, ErrDataNotValid{Param: "id", Message: "id is not valid"}
	}

	news, err := s.repository.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return NewsDTO{}, err
		}

		return NewsDTO{}, fmt.Errorf("couldn't get news w error %w", err)
	}

	var result NewsDTO
	result.Map(&news)

	return result, nil
}

func (s *service) GetAllNewsWithPagination(ctx context.Context, limit uint64, pToken string) ([]NewsDTO, error) {
	if pToken != "" {
		if !isValidID(pToken) {
			return nil, ErrDataNotValid{Param: "id", Message: "pToken is not valid"}
		}
	}

	news, err := s.repository.GetAllWithPagination(ctx, limit, pToken)
	if err != nil {
		return nil, fmt.Errorf("couldn't get all news w error %w", err)
	}

	if len(news) == 0 {
		return nil, ErrNotFound
	}

	result := make([]NewsDTO, 0)
	for i := 0; i < len(news); i++ {
		var dto NewsDTO
		dto.Map(&news[i])
		result = append(result, dto)
	}

	return result, nil
}

func (s *service) UpdateNews(ctx context.Context, id string, n NewsDTO) error {
	if !isValidID(id) {
		return ErrDataNotValid{Param: "id", Message: "id is not valid"}
	}

	if !n.IsEmpty() {
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
	if !isValidID(id) {
		return ErrDataNotValid{Param: "id", Message: "id is not valid"}
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

func isValidID(id string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(id)
}
