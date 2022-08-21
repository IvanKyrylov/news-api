package news

type NewsDTO struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"update_at"`
}

func (dto *NewsDTO) IsEmpty() bool {
	switch {
	case dto.Title == "":
		return false
	case dto.Content == "":
		return false
	}

	return true
}

func (dto *NewsDTO) Map(n *News) {
	dto.ID = n.ID
	dto.Title = n.Title
	dto.Content = n.Content
	dto.CreatedAt = n.CreatedAt.UTC().Format("2006-01-02T15:04:05")
	if !n.UpdatedAt.Time.IsZero() {
		dto.UpdateAt = n.UpdatedAt.Time.UTC().Format("2006-01-02T15:04:05")
	}
}
