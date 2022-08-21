package news

import (
	"testing"
)

func TestNewsDTO_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		n    *NewsDTO
		want bool
	}{
		{
			name: "1",
			n: &NewsDTO{
				Title:   "test",
				Content: "test",
			},
			want: true,
		},
		{
			name: "2",
			n: &NewsDTO{
				Title:   "test",
				Content: "",
			},
			want: false,
		},
		{
			name: "3",
			n: &NewsDTO{
				Title:   "",
				Content: "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.IsEmpty(); got != tt.want {
				t.Errorf("NewsDTO.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewsDTO_Map(t *testing.T) {
	type args struct {
		n *News
	}
	tests := []struct {
		name string
		dto  *NewsDTO
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.dto.Map(tt.args.n)
		})
	}
}
