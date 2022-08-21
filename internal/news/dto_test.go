package news

import "testing"

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
