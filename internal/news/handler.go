package news

import (
	"log"
	"net/http"

	"github.com/IvanKyrylov/news-api/internal/apperror"
)

const (
	newsURL = "/api/news"
)

type Handler struct {
	Logger      *log.Logger
	NewsService Service
}

func (h *Handler) Register(router *http.ServeMux) {
	router.HandleFunc(newsURL, apperror.Middleware(h.GetNews))
}

func (h *Handler) GetNews(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return apperror.BadRequestError("metod GET")
	}
	h.Logger.Println("GET News")
	w.Header().Set("Content-Type", "application/json")

	// userBytes, err := json.Marshal(user)
	// if err != nil {
	// 	return err
	// }

	// w.WriteHeader(http.StatusOK)
	w.Write([]byte("test"))
	return nil
}
