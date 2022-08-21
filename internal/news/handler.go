package news

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/IvanKyrylov/news-api/internal/apperror"
)

const (
	newsURL     = "/posts"
	newsByIDURL = "/posts/"
)

type Handler struct {
	Logger      *log.Logger
	NewsService Service
}

func (h *Handler) Register(router *http.ServeMux) {
	router.HandleFunc(newsURL, apperror.Middleware(h.Posts))
	router.HandleFunc(newsByIDURL, apperror.Middleware(h.PostsByID))
}

func (h *Handler) Posts(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		return h.get(w, r)
	case http.MethodPost:
		return h.create(w, r)
	default:
		return apperror.BadRequestError("metod GET, POST")
	}
}

func (h *Handler) PostsByID(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		return h.getByID(w, r)
	case http.MethodPut:
		return h.put(w, r)
	case http.MethodDelete:
		return h.delete(w, r)
	default:
		return apperror.BadRequestError("metod GET, PUT, DELETE")
	}
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) error {
	limitQuery := r.URL.Query().Get("limit")
	limit, err := strconv.ParseUint(limitQuery, 10, 64)
	if err != nil {
		return apperror.BadRequestError("limit query parameter is required integers")
	}

	// pToken is last ID in prev page
	pToken := r.URL.Query().Get("ptoken")
	news, err := h.NewsService.GetAllNewsWithPagination(r.Context(), limit, pToken)
	if err != nil {
		return apperror.BadRequestError(err.Error())
	}

	newsByte, err := json.Marshal(news)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(newsByte)

	return nil

}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) error {
	news := make([]NewsDTO, 0)
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&news); err != nil {
		return apperror.BadRequestError("invalid JSON scheme. check swagger API")
	}

	ids, err := h.NewsService.CreateNews(r.Context(), news)
	if err != nil {
		return apperror.BadRequestError(err.Error())
	}

	w.Header().Set("Location", fmt.Sprintf("%s/%s", newsURL, ids[0]))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

	return nil
}

func (h *Handler) getByID(w http.ResponseWriter, r *http.Request) error {
	id := r.URL.Path[len(newsByIDURL):]
	if id == "" {
		return apperror.BadRequestError("id query parameter is required")
	}

	news, err := h.NewsService.GetNewsByID(r.Context(), id)
	if err != nil {
		return apperror.BadRequestError(err.Error())
	}

	newsByte, err := json.Marshal(news)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(newsByte)

	return nil

}

func (h *Handler) put(w http.ResponseWriter, r *http.Request) error {
	id := r.URL.Path[len(newsByIDURL):]
	if id == "" {
		return apperror.BadRequestError("id query parameter is required")
	}

	var news NewsDTO
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&news); err != nil {
		return apperror.BadRequestError("invalid JSON scheme. check swagger API")
	}

	err := h.NewsService.UpdateNews(r.Context(), id, news)
	if err != nil {
		return apperror.BadRequestError(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

	return nil

}

func (h *Handler) delete(w http.ResponseWriter, r *http.Request) error {
	id := r.URL.Path[len(newsByIDURL):]
	if id == "" {
		return apperror.BadRequestError("id query parameter is required")
	}

	err := h.NewsService.DeleteNews(r.Context(), id)
	if err != nil {
		return apperror.BadRequestError(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

	return nil
}
