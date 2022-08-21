package db

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/IvanKyrylov/news-api/internal/news"
)

// Query builder
// Why don't I use "?"? Because I have the same parameter types and I don't trust reflection

func insertNewsQuery(news news.News) string {
	q := "INSERT INTO posts (title, content) VALUE "

	var sb strings.Builder
	sb.WriteString(q)
	sb.WriteString("('")
	sb.WriteString(news.Title)
	sb.WriteString("','")
	sb.WriteString(news.Content)
	sb.WriteString("')")

	// INSERT INTO posts (title, content) VALUE ('news[i].Title', 'news[i].Content')
	return sb.String()
}

func bulkInsertNewsQuery(news []news.News) string {
	q := "INSERT INTO posts (title, content) VALUE "

	var sb strings.Builder
	sb.WriteString(q)
	for i := 0; i < len(news); i++ {
		sb.WriteString("('")
		sb.WriteString(news[i].Title)
		sb.WriteString("','")
		sb.WriteString(news[i].Content)
		sb.WriteString("'),")
	}

	// Why don't I use "?"? Because I have the same parameter types and I don't trust reflection
	// INSERT INTO posts (title, content) VALUE ('news[i].Title', 'news[i].Content'), ('news[i].Title', 'news[i].Content')
	return fmt.Sprintf(strings.TrimRight(sb.String(), ","))
}

func selectNewsByIDQuery(id string) string {
	return fmt.Sprintf("SELECT id, title, content, created_at, updated_at FROM posts WHERE id = '%s'", id)
}

func selectAllNewsQueryWithPagination(limit uint64, lastID string) string {
	q := "SELECT id, title, content, created_at, updated_at FROM posts"

	var sb strings.Builder
	sb.WriteString(q)

	if lastID != "" {
		sb.WriteString(" WHERE id < '")
		sb.WriteString(lastID)
		sb.WriteString("'")
	}

	sb.WriteString(" ORDER BY id DESC")
	if limit != 0 {
		sb.WriteString(" LIMIT ")
		sb.WriteString(strconv.FormatUint(limit, 10))
	}

	// SELECT id, title, content, created_at, updated_at FROM posts WHERE id < 'lastID' ORDER BY id DESC LIMIT limit
	return sb.String()
}

func selectNewsByIDForUpdateQuery(id string) string {
	return fmt.Sprintf("SELECT created_at FROM posts WHERE id = '%s' FOR UPDATE", id)
}

func updateNewsByID(n news.News) string {
	return fmt.Sprintf("UPDATE posts SET title = '%s', content = '%s', updated_at = now() AT TIME ZONE 'utc' WHERE id = '%s'", n.Title, n.Content, n.ID)
}

func deleteNewsByID(id string) string {
	return fmt.Sprintf("DELETE FROM posts WHERE id = '%s'", id)
}
