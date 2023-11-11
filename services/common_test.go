package commands

import (
	"TDBackend/logger"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Test_loadFileToMemory(t *testing.T) {
	logger.Init("debug", time.RFC3339, "TDBackend")

	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loadFileToMemory()
		})
	}
}

func Test_getData(t *testing.T) {

	logger.Init("debug", time.RFC3339, "TDBackend")

	_ = StartServer()
	partner, _ := http.NewRequest("GET", "http://rehabaam.net:8080/TriDubai/v1/sessions/getPartners", nil)
	session, _ := http.NewRequest("POST", "http://rehabaam.net:8080/TriDubai/v1/sessions/getSessions", nil)
	rw := httptest.NewRecorder()

	type args struct {
		endPoint string
		w        http.ResponseWriter
		r        *http.Request
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "getData_success",
			args: args{
				endPoint: "Partners",
				w:        rw,
				r:        partner,
			},
			wantErr: false,
		},
		{
			name: "getData_failure",
			args: args{
				endPoint: "partners!",
				w:        rw,
				r:        session,
			},
			wantErr: false,
		},
		{
			name: "getData_nil",
			args: args{
				endPoint: "partners",
				w:        rw,
				r:        &http.Request{},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantErr {
					t.Errorf("getData(), wantPanic = %v", tt.wantErr)
				}
			}()
			_, _ = getData(tt.args.endPoint, tt.args.w, tt.args.r)
		})
	}
}

func Test_readFile(t *testing.T) {
	logger.Init("debug", time.RFC3339, "TDBackend")

	type args struct {
		endPoint string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readFile(tt.args.endPoint)
			fmt.Println(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("readFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("readFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getImage(t *testing.T) {
	logger.Init("debug", time.RFC3339, "TDBackend")

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getImage(tt.args.w, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("getImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getImage() = %v, want %v", got, tt.want)
			}
		})
	}
}
