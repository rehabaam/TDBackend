package commands

import (
	labels "TDBackend/localization"
	applog "TDBackend/logger"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Test_getData(t *testing.T) {

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
				r:        nil,
			},
			wantErr: false,
		},
	}

	_ = applog.Init("debug", time.Now().Format(labels.RFC3339Milli), "TDBackend")
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
