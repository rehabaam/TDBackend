package commands

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_getSessions(t *testing.T) {

	data, _ := http.NewRequest("GET", "https://apps.tridubai.org/api/v1/sessions", nil)
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
				r: data,
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

	data, _ := http.NewRequest("GET", "https://apps.tridubai.org/api/v1/deals", nil)
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
			name: "getDeals_failure",
			args: args{
				w: rw,
				r: data,
			},
			wantPanic: true,
		},
		{
			name: "getDeals_success",
			args: args{
				w: rw,
				r: data,
			},
			wantPanic: true,
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

	data, _ := http.NewRequest("GET", "https://apps.tridubai.org/api/v1/partners", nil)
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
			name: "getPartner_failure",
			args: args{
				w: rw,
				r: data,
			},
			wantPanic: true,
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

	data, _ := http.NewRequest("GET", "https://apps.tridubai.org/api/v1/kit", nil)
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
			name: "getKit_failure",
			args: args{
				w: rw,
				r: data,
			},
			wantPanic: true,
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

	data, _ := http.NewRequest("GET", "https://apps.tridubai.org/api/v1/faqs", nil)
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
			name: "getFAQs_failure",
			args: args{
				w: rw,
				r: data,
			},
			wantPanic: true,
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
