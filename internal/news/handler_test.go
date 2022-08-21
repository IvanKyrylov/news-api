package news

import (
	"net/http"
	"testing"
)

func TestHandler_Register(t *testing.T) {
	type args struct {
		router *http.ServeMux
	}
	tests := []struct {
		name string
		h    *Handler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Register(tt.args.router)
		})
	}
}

func TestHandler_Posts(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		h       *Handler
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.Posts(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("Handler.Posts() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_PostsByID(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		h       *Handler
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.PostsByID(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("Handler.PostsByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_get(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		h       *Handler
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.get(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("Handler.get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_create(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		h       *Handler
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.create(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("Handler.create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_getByID(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		h       *Handler
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.getByID(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("Handler.getByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_put(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		h       *Handler
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.put(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("Handler.put() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_delete(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		h       *Handler
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.delete(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("Handler.delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
