package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/IvanKyrylov/news-api/internal/news"
	"github.com/IvanKyrylov/news-api/pkg/postgres"
)

var _ news.Repository = &repository{}

type repository struct {
	db     postgres.Client
	logger *log.Logger
}

func NewRepository(db postgres.Client, logger *log.Logger) news.Repository {
	return &repository{
		db:     db,
		logger: logger,
	}
}

func (s *repository) CreateOne(ctx context.Context, n news.News) (string, error) {
	if err := validation(n); err != nil {
		return "", fmt.Errorf("failed validation w error %w", err)
	}

	var id string
	err := s.db.QueryRowContext(ctx, insertNewsQuery(n)).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("failed exec query w error %w", err)
	}

	return id, nil
}

func (s *repository) CreateMany(ctx context.Context, n []news.News) ([]string, error) {
	if len(n) == 0 {
		return nil, errors.New("length news equal zero (0)")
	}

	rows, err := s.db.QueryContext(ctx, insertNewsQuery(n...))
	if err != nil {
		return nil, fmt.Errorf("failed exec query w error %w", err)
	}

	ids := make([]string, 0)
	for rows.Next() {
		var id string
		rows.Scan(&id)
		ids = append(ids, id)
	}

	return ids, nil
}

func (s *repository) GetByID(ctx context.Context, id string) (news.News, error) {
	if id == "" {
		return news.News{}, errors.New("id is empty")
	}

	var n news.News

	err := s.db.QueryRowContext(ctx, selectNewsByIDQuery(id)).Scan(&n.ID, &n.Title, &n.Content, &n.CreatedAt, &n.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return news.News{}, news.ErrNotFound
		}

		return news.News{}, fmt.Errorf("failed to query rowx w error %w", err)
	}

	return n, nil
}

func (s *repository) GetAllWithPagination(ctx context.Context, limit uint64, lastID string) ([]news.News, error) {
	rows, err := s.db.QueryContext(ctx, selectAllNewsQueryWithPagination(limit, lastID))
	if err != nil {
		return nil, fmt.Errorf("failed to query w error %w", err)
	}

	result := make([]news.News, 0)
	for rows.Next() {

		var n news.News

		if err := rows.Scan(&n.ID, &n.Title, &n.Content, &n.CreatedAt, &n.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan w error %w", err)
		}

		result = append(result, n)

	}

	return result, nil
}

func (s *repository) Update(ctx context.Context, n news.News) error {
	if err := validation(n); err != nil {
		return fmt.Errorf("failed validation w error %w", err)
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed start transaction w error %w", err)
	}

	var createdAt time.Time
	err = s.db.QueryRowContext(ctx, selectNewsByIDForUpdateQuery(n.ID)).Scan(&createdAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			tx.Rollback()
			return news.ErrNotFound
		}

		tx.Rollback()
		return fmt.Errorf("failed to query rowx w error %w", err)
	}

	_, err = tx.ExecContext(ctx, updateNewsByID(n))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed exec query w error %w", err)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed commit transaction w error %w", err)
	}

	return nil
}

func (s *repository) DeleteByID(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id is empty")
	}

	rows, err := s.db.ExecContext(ctx, deleteNewsByID(id))
	if err != nil {
		return fmt.Errorf("failed exec query w error %w", err)
	}

	affected, err := rows.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if affected == 0 {
		return news.ErrNotFound
	}

	return nil
}

func validation(n news.News) error {
	switch {
	case n.Title == "":
		return errors.New("title is empty")
	case n.Content == "":
		return errors.New("content is empty")
	}
	return nil
}
