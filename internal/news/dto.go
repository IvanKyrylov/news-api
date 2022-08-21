package news

type NewsDTO struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (n *NewsDTO) IsEmpty() bool {
	switch {
	case n.Title == "":
		return false
	case n.Content == "":
		return false
	}

	return true
}
