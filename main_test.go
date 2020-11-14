package tdbackend

import (
	"testing"
)

func Test_StartServer(t *testing.T) {
	tests := []struct {
		name string
		want interface{}
	}{
		{
			name: "sucess",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := startServer()
			if ans != nil {
				t.Error("Expected", nil, "Got", ans)
			}

		})
	}
}
