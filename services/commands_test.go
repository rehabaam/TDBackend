package commands

import (
	"github.com/rehabaam/TDBackend/logger"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
)

func Test_getSessions(t *testing.T) {

	logger.Init("debug", time.RFC3339, "TDBackend")
	repo["Sessions"] = `test`

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
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "https://localhost/api/v1/sessions", nil),
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
}

func Test_getDeals(t *testing.T) {

	logger.Init("debug", time.RFC3339, "TDBackend")
	repo["Deals"] = `test`

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
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "https://localhost/api/v1/deals", nil),
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
}

func Test_getPartners(t *testing.T) {

	logger.Init("debug", time.RFC3339, "TDBackend")
	repo["Partners"] = `test`

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
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "https://localhost/api/v1/partners", nil),
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
}

func Test_getKit(t *testing.T) {

	logger.Init("debug", time.RFC3339, "TDBackend")
	repo["Kit"] = `test`

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
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "https://localhost/api/v1/kit", nil),
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
}

func Test_getFAQs(t *testing.T) {

	logger.Init("debug", time.RFC3339, "TDBackend")
	repo["FAQs"] = `test`

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
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "https://localhost/api/v1/faqs", nil),
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
}

func Test_serveImage(t *testing.T) {

	logger.Init("debug", time.RFC3339, "TDBackend")

	r := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/img/FFMC.jpg", nil)
	r = mux.SetURLVars(r, map[string]string{"name": "FFMC.jpg"})

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
				w: httptest.NewRecorder(),
				r: r,
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
}

func TestStartServer(t *testing.T) {
	srv := NewServer()
	go func() {
		time.Sleep(1 * time.Second)
		srv.Shutdown(context.Background())
	}()
	err := srv.Start()
	if err != nil {
		t.Error("unexpected error:", err)
	}
}
