package db

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/IvanKyrylov/news-api/internal/news"
)

// Query builder
// Why don't I use "?"? Because I have the same parameter types and I don't trust reflection

func insertNewsQuery(news ...news.News) string {
	q := "INSERT INTO news (title, content) VALUES "

	var sb strings.Builder
	sb.WriteString(q)
	for i := 0; i < len(news); i++ {
		sb.WriteString("('")
		sb.WriteString(news[i].Title)
		sb.WriteString("','")
		sb.WriteString(news[i].Content)
		sb.WriteString("'),")
	}

	result := strings.TrimRight(sb.String(), ",")
	sb.Reset()
	sb.WriteString(result)
	sb.WriteString(" RETURNING id")

	// Why don't I use "?"? Because I have the same parameter types and I don't trust reflection
	// INSERT INTO news (title, content) VALUE ('news[i].Title','news[i].Content'),('news[i].Title','news[i].Content') RETURNING id
	return sb.String()
}

func selectNewsByIDQuery(id string) string {
	return fmt.Sprintf("SELECT id, title, content, created_at, updated_at FROM news WHERE id = '%s'", id)
}

func selectAllNewsQueryWithPagination(limit uint64, lastID string) string {
	q := "SELECT id, title, content, created_at, updated_at FROM news"

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

	// SELECT id, title, content, created_at, updated_at FROM news WHERE id < 'lastID' ORDER BY id DESC LIMIT limit
	return sb.String()
}

func selectNewsByIDForUpdateQuery(id string) string {
	return fmt.Sprintf("SELECT created_at FROM news WHERE id = '%s' FOR UPDATE", id)
}

func updateNewsByID(n news.News) string {
	return fmt.Sprintf("UPDATE news SET title = '%s', content = '%s', updated_at = now() AT TIME ZONE 'utc' WHERE id = '%s'", n.Title, n.Content, n.ID)
}

func deleteNewsByID(id string) string {
	return fmt.Sprintf("DELETE FROM news WHERE id = '%s'", id)
}
