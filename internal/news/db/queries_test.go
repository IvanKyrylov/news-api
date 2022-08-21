package db

import (
	"testing"

	"github.com/IvanKyrylov/news-api/internal/news"
)

func Test_insertNewsQuery(t *testing.T) {
	type args struct {
		news news.News
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1",
			args: args{
				news.News{Title: "1", Content: "1"},
			},
			want: `INSERT INTO posts (title, content) VALUE ('1','1')`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertNewsQuery(tt.args.news); got != tt.want {
				t.Errorf("insertNewsQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bulkInsertNewsQuery(t *testing.T) {
	type args struct {
		param []news.News
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1",
			args: args{
				[]news.News{{Title: "1", Content: "1"}, {Title: "2", Content: "2"}, {Title: "3", Content: "3"}},
			},
			want: `INSERT INTO posts (title, content) VALUE ('1','1'),('2','2'),('3','3')`,
		},
		{name: "2",
			args: args{
				[]news.News{{Title: "1", Content: ""}, {Title: "2", Content: ""}, {Title: "3", Content: ""}},
			},
			want: `INSERT INTO posts (title, content) VALUE ('1',''),('2',''),('3','')`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bulkInsertNewsQuery(tt.args.param); got != tt.want {
				t.Errorf("insertNewsQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_selectNewsByIDQuery(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				id: "test",
			},
			want: "SELECT id, title, content, created_at, updated_at FROM posts WHERE id = 'test'",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := selectNewsByIDQuery(tt.args.id); got != tt.want {
				t.Errorf("selectNewsByIDQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_selectAllNewsQueryWithPagination(t *testing.T) {
	type args struct {
		limit  uint64
		lastID string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				limit:  100,
				lastID: "lastID",
			},
			want: `SELECT id, title, content, created_at, updated_at FROM posts WHERE id < 'lastID' ORDER BY id DESC LIMIT 100`,
		},
		{
			name: "2",
			args: args{
				limit:  0,
				lastID: "lastID",
			},
			want: `SELECT id, title, content, created_at, updated_at FROM posts WHERE id < 'lastID' ORDER BY id DESC`,
		},
		{
			name: "4",
			args: args{
				limit:  100,
				lastID: "",
			},
			want: `SELECT id, title, content, created_at, updated_at FROM posts ORDER BY id DESC LIMIT 100`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := selectAllNewsQueryWithPagination(tt.args.limit, tt.args.lastID); got != tt.want {
				t.Errorf("selectAllNewsQueryWithPagination() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_selectNewsByIDForUpdateQuery(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				id: "test",
			},
			want: "SELECT created_at FROM posts WHERE id = 'test' FOR UPDATE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := selectNewsByIDForUpdateQuery(tt.args.id); got != tt.want {
				t.Errorf("selectNewsByIDForUpdateQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_updateNewsByID(t *testing.T) {
	type args struct {
		n news.News
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1",
			args: args{
				news.News{
					ID:      "test",
					Title:   "test title",
					Content: "test content",
				},
			},
			want: "UPDATE posts SET title = 'test title', content = 'test content', updated_at = now() AT TIME ZONE 'utc' WHERE id = 'test'",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := updateNewsByID(tt.args.n); got != tt.want {
				t.Errorf("updateNewsByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deleteNewsByID(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				id: "test",
			},
			want: "DELETE FROM posts WHERE id = 'test'",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteNewsByID(tt.args.id); got != tt.want {
				t.Errorf("deleteNewsByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
