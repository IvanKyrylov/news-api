package news

import (
	"database/sql"
	"time"
)

type News struct {
	ID        string
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
