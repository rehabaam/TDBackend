package commands

import (
	labels "TDBackend/localization"
	applog "TDBackend/logger"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Test_readFileData(t *testing.T) {

	partner, _ := http.NewRequest("GET", "http://rehabaam.net:8080/TriDubai/v1/sessions/getPartners", nil)
	session, _ := http.NewRequest("GET", "http://rehabaam.net:8080/TriDubai/v1/sessions/getSessions", nil)
	rw := httptest.NewRecorder()

	type args struct {
		endPoint string
		w        http.ResponseWriter
		r        *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "readFileData_success",
			args: args{
				endPoint: "Partners",
				w:        rw,
				r:        partner,
			},
		},
		{
			name: "readFileData_failure",
			args: args{
				endPoint: "partners!",
				w:        rw,
				r:        session,
			},
		},
		{
			name: "readFileData_nil",
			args: args{
				endPoint: "partners",
				w:        rw,
				r:        nil,
			},
		},
	}

	applog.Init("debug", time.Now().Format(labels.RFC3339Milli), "TDBackend")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			readFileData(tt.args.endPoint, tt.args.w, tt.args.r)
		})
	}
}
