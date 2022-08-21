package news

import (
	"context"
	"errors"
	"log"
	"reflect"
	"testing"
	"time"
)

type pgMock struct {
}

func NewPGMock() Repository {
	return &pgMock{}
}

func (s *pgMock) CreateOne(ctx context.Context, n News) (string, error) {
	return "fb377dbd-27b0-436a-8c05-b48276d279f1", nil
}

func (s *pgMock) CreateMany(ctx context.Context, n []News) ([]string, error) {
	return []string{"fb377dbd-27b0-436a-8c05-b48276d279f1"}, nil
}

func (s *pgMock) GetByID(ctx context.Context, id string) (News, error) {

	return News{
		ID:        "fb377dbd-27b0-436a-8c05-b48276d279f1",
		Title:     "test",
		Content:   "test",
		CreatedAt: time.Time{},
	}, nil
}

func (s *pgMock) GetAllWithPagination(ctx context.Context, limit uint64, lastID string) ([]News, error) {
	return []News{{ID: "fb377dbd-27b0-436a-8c05-b48276d279f1",
		Title:     "test",
		Content:   "test",
		CreatedAt: time.Time{}}}, nil
}

func (s *pgMock) Update(ctx context.Context, n News) error {
	return nil
}

func (s *pgMock) DeleteByID(ctx context.Context, id string) error {
	return nil
}

func validation(n News) error {
	switch {
	case n.Title == "":
		return errors.New("title is empty")
	case n.Content == "":
		return errors.New("content is empty")
	}
	return nil
}

func TestNewService(t *testing.T) {
	type args struct {
		pgMock Repository
		logger *log.Logger
	}
	pg := &pgMock{}
	logger := log.Default()
	tests := []struct {
		name string
		args args
		want Service
	}{
		{
			name: "1",
			args: args{
				pgMock: pg,
				logger: logger,
			},
			want: &service{
				repository: pg,
				logger:     logger,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.pgMock, tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_CreateNews(t *testing.T) {
	type args struct {
		ctx  context.Context
		news []NewsDTO
	}
	pg := &pgMock{}
	logger := log.Default()
	s := &service{
		repository: pg,
		logger:     logger,
	}
	tests := []struct {
		name    string
		s       *service
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "1",
			s:    s,
			args: args{
				context.TODO(),
				[]NewsDTO{
					{Title: "test", Content: "test"},
				},
			},
			want:    []string{"fb377dbd-27b0-436a-8c05-b48276d279f1"},
			wantErr: false,
		},
		{
			name: "2",
			s:    s,
			args: args{
				context.TODO(),
				[]NewsDTO{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "2",
			s:    s,
			args: args{
				context.TODO(),
				[]NewsDTO{
					{Title: "", Content: ""},
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.CreateNews(tt.args.ctx, tt.args.news)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.CreateNews() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.CreateNews() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetNewsByID(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}

	pg := &pgMock{}
	logger := log.Default()
	s := &service{
		repository: pg,
		logger:     logger,
	}
	tests := []struct {
		name    string
		s       *service
		args    args
		want    NewsDTO
		wantErr bool
	}{
		{
			name: "1",
			s:    s,
			args: args{
				context.TODO(),
				"fb377dbd-27b0-436a-8c05-b48276d279f1",
			},
			want: NewsDTO{
				ID:        "fb377dbd-27b0-436a-8c05-b48276d279f1",
				Title:     "test",
				Content:   "test",
				CreatedAt: time.Time.Format(time.Time{}, "2006-01-02T15:04:05"),
				UpdateAt:  "",
			},
			wantErr: false,
		},
		{
			name: "2",
			s:    s,
			args: args{
				context.TODO(),
				"",
			},
			want:    NewsDTO{},
			wantErr: true,
		},
		{
			name: "3",
			s:    s,
			args: args{
				context.TODO(),
				"agGAwgawa",
			},
			want:    NewsDTO{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetNewsByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetNewsByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetNewsByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetAllNewsWithPagination(t *testing.T) {
	type args struct {
		ctx    context.Context
		limit  uint64
		lastID string
	}

	pg := &pgMock{}
	logger := log.Default()
	s := &service{
		repository: pg,
		logger:     logger,
	}
	tests := []struct {
		name    string
		s       *service
		args    args
		want    []NewsDTO
		wantErr bool
	}{
		{
			name: "1",
			s:    s,
			args: args{
				context.TODO(),
				100,
				"fb377dbd-27b0-436a-8c05-b48276d279f1",
			},
			want: []NewsDTO{{
				ID:        "fb377dbd-27b0-436a-8c05-b48276d279f1",
				Title:     "test",
				Content:   "test",
				CreatedAt: time.Time.Format(time.Time{}, "2006-01-02T15:04:05"),
				UpdateAt:  "",
			}},
			wantErr: false,
		},
		{
			name: "2",
			s:    s,
			args: args{
				context.TODO(),
				100,
				"",
			},
			want: []NewsDTO{{
				ID:        "fb377dbd-27b0-436a-8c05-b48276d279f1",
				Title:     "test",
				Content:   "test",
				CreatedAt: time.Time.Format(time.Time{}, "2006-01-02T15:04:05"),
				UpdateAt:  "",
			}},
			wantErr: false,
		},
		{
			name: "3",
			s:    s,
			args: args{
				context.TODO(),
				0,
				"",
			},
			want: []NewsDTO{{
				ID:        "fb377dbd-27b0-436a-8c05-b48276d279f1",
				Title:     "test",
				Content:   "test",
				CreatedAt: time.Time.Format(time.Time{}, "2006-01-02T15:04:05"),
				UpdateAt:  "",
			}},
			wantErr: false,
		},
		{
			name: "4",
			s:    s,
			args: args{
				context.TODO(),
				0,
				"wawwagw",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetAllNewsWithPagination(tt.args.ctx, tt.args.limit, tt.args.lastID)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetAllNewsWithPagination() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetAllNewsWithPagination() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_UpdateNews(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
		n   NewsDTO
	}

	pg := &pgMock{}
	logger := log.Default()
	s := &service{
		repository: pg,
		logger:     logger,
	}

	tests := []struct {
		name    string
		s       *service
		args    args
		wantErr bool
	}{
		{
			name: "1",
			s:    s,
			args: args{
				context.TODO(),
				"fb377dbd-27b0-436a-8c05-b48276d279f1",
				NewsDTO{
					ID:        "fb377dbd-27b0-436a-8c05-b48276d279f1",
					Title:     "test",
					Content:   "test",
					CreatedAt: time.Time.Format(time.Time{}, "2006-01-02T15:04:05"),
					UpdateAt:  "",
				},
			},
			wantErr: false,
		},
		{
			name: "2",
			s:    s,
			args: args{
				context.TODO(),
				"",
				NewsDTO{
					ID:        "fb377dbd-27b0-436a-8c05-b48276d279f1",
					Title:     "test",
					Content:   "test",
					CreatedAt: time.Time.Format(time.Time{}, "2006-01-02T15:04:05"),
					UpdateAt:  "",
				},
			},
			wantErr: true,
		},
		{
			name: "3",
			s:    s,
			args: args{
				context.TODO(),
				"faawsfafwfa",
				NewsDTO{
					ID:        "faaaafssfafwa",
					Title:     "test",
					Content:   "test",
					CreatedAt: time.Time.Format(time.Time{}, "2006-01-02T15:04:05"),
					UpdateAt:  "",
				},
			},
			wantErr: true,
		},
		{
			name: "4",
			s:    s,
			args: args{
				context.TODO(),
				"fb377dbd-27b0-436a-8c05-b48276d279f1",
				NewsDTO{
					ID:        "fb377dbd-27b0-436a-8c05-b48276d279f1",
					Title:     "",
					Content:   "",
					CreatedAt: time.Time.Format(time.Time{}, "2006-01-02T15:04:05"),
					UpdateAt:  "",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.UpdateNews(tt.args.ctx, tt.args.id, tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("service.UpdateNews() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_DeleteNews(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}

	pg := &pgMock{}
	logger := log.Default()
	s := &service{
		repository: pg,
		logger:     logger,
	}
	tests := []struct {
		name    string
		s       *service
		args    args
		wantErr bool
	}{
		{
			name: "1",
			s:    s,
			args: args{
				context.TODO(),
				"fb377dbd-27b0-436a-8c05-b48276d279f1",
			},
			wantErr: false,
		},
		{
			name: "2",
			s:    s,
			args: args{
				context.TODO(),
				"",
			},
			wantErr: true,
		},
		{
			name: "3",
			s:    s,
			args: args{
				context.TODO(),
				"aawfawfa",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.DeleteNews(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("service.DeleteNews() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
