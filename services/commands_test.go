package commands

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_getSessions(t *testing.T) {

	partner, _ := http.NewRequest("GET", "http://rehabaam.net:8080/TriDubai/v1/sessions/getSessions", nil)
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
			name: "getSessions_failure",
			args: args{
				w: rw,
				r: partner,
			},
			wantPanic: true,
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

	sessions, _ := http.NewRequest("GET", "http://rehabaam.net:8080/TriDubai/v1/sessions/getSessions", nil)
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
			name: "getSessions_failure",
			args: args{
				w: rw,
				r: sessions,
			},
			wantPanic: true,
		},
		{
			name: "getSessions_success",
			args: args{
				w: rw,
				r: sessions,
			},
			wantPanic: true,
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
			getDeals(tt.args.w, tt.args.r)
		})
	}
}

func Test_getPartners(t *testing.T) {

	partner, _ := http.NewRequest("GET", "http://rehabaam.net:8080/TriDubai/v1/sessions/getSessions", nil)
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
			name: "getSessions_failure",
			args: args{
				w: rw,
				r: partner,
			},
			wantPanic: true,
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
			getPartners(tt.args.w, tt.args.r)
		})
	}
}

func TestRunServer(t *testing.T) {
	tests := []struct {
		name      string
		wantPanic bool
	}{
		// {
		// 	name:      "getSessions_failure",
		// 	wantPanic: true,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("getSessions(), wantPanic = %v", tt.wantPanic)
				}
			}()
			RunServer()
		})
	}
}
