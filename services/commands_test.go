package commands

import (
	"TDBackend/logger"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Test_getSessions(t *testing.T) {

	logger.Init("debug", time.RFC3339, "TDBackend")

	StartServer()
	data, _ := http.NewRequest("GET", "http://localhost:8080/api/v1/sessions", nil)
	rw := httptest.NewRecorder()

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name      string
		args      args
		wantPanic bool
	}{
		{
			name: "getSessions_Success",
			args: args{
				w: rw,
				r: data,
			},
			wantPanic: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("getSessions(), wantPanic = %v", tt.wantPanic)
				}
			}()
			getSessions(tt.args.w, tt.args.r)
		})
	}
	StopServer()
}

func Test_getDeals(t *testing.T) {

	logger.Init("debug", time.RFC3339, "TDBackend")

	_ = StartServer()
	data, _ := http.NewRequest("GET", "http://localhost:8080/api/v1/deals", nil)
	rw := httptest.NewRecorder()

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name      string
		args      args
		wantPanic bool
	}{
		{
			name: "getDeals_success",
			args: args{
				w: rw,
				r: data,
			},
			wantPanic: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("getDeals(), wantPanic = %v", tt.wantPanic)
				}
			}()
			getDeals(tt.args.w, tt.args.r)
		})
	}
	StopServer()
}

func Test_getPartners(t *testing.T) {

	logger.Init("debug", time.RFC3339, "TDBackend")

	_ = StartServer()
	data, _ := http.NewRequest("GET", "http://localhost:8080/api/v1/partners", nil)
	rw := httptest.NewRecorder()

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name      string
		args      args
		wantPanic bool
	}{
		{
			name: "getPartner_success",
			args: args{
				w: rw,
				r: data,
			},
			wantPanic: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("getPartner(), wantPanic = %v", tt.wantPanic)
				}
			}()
			getPartners(tt.args.w, tt.args.r)
		})
	}
	StopServer()
}

func Test_getKit(t *testing.T) {

	logger.Init("debug", time.RFC3339, "TDBackend")

	_ = StartServer()
	data, _ := http.NewRequest("GET", "http://localhost:8080/api/v1/kit", nil)
	rw := httptest.NewRecorder()

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name      string
		args      args
		wantPanic bool
	}{
		{
			name: "getKit_success",
			args: args{
				w: rw,
				r: data,
			},
			wantPanic: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("getKit(), wantPanic = %v", tt.wantPanic)
				}
			}()
			getKit(tt.args.w, tt.args.r)
		})
	}
	StopServer()
}

func Test_getFAQs(t *testing.T) {

	logger.Init("debug", time.RFC3339, "TDBackend")

	_ = StartServer()
	data, _ := http.NewRequest("GET", "http://localhost:8080/api/v1/faqs", nil)
	rw := httptest.NewRecorder()

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name      string
		args      args
		wantPanic bool
	}{
		{
			name: "getFAQs_Success",
			args: args{
				w: rw,
				r: data,
			},
			wantPanic: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("getFAQs(), wantPanic = %v", tt.wantPanic)
				}
			}()
			getFAQs(tt.args.w, tt.args.r)
		})
	}
	StopServer()
}

func TestStartServer(t *testing.T) {
	logger.Init("debug", time.RFC3339, "TDBackend")

	tests := []struct {
		name      string
		wantPanic bool
	}{
		{
			name:      "StartServer_Success",
			wantPanic: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("getSessions(), wantPanic = %v", tt.wantPanic)
				}
			}()
			_ = StartServer()
		})
	}
	StopServer()
}

func Test_serveImage(t *testing.T) {

	logger.Init("debug", time.RFC3339, "TDBackend")

	_ = StartServer()
	data, _ := http.NewRequest("GET", "http://localhost:8080/api/v1/img/FFMC.jpg", nil)
	rw := httptest.NewRecorder()

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name      string
		args      args
		wantPanic bool
	}{
		{
			name: "getFFMC_Success",
			args: args{
				w: rw,
				r: data,
			},
			wantPanic: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("serveImage(), wantPanic = %v", tt.wantPanic)
				}
			}()
			serveImage(tt.args.w, tt.args.r)
		})
	}
	StopServer()
}
