package middleware

import (
	applog "TDBackend/logger"
	labels "TDBackend/localization"
	"crypto/tls"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type mockResponseWriter struct {
	data       []byte
	statusCode int
	h          http.Header
}

func (m *mockResponseWriter) Header() http.Header {
	return m.h
}

func (m *mockResponseWriter) Write(d []byte) (int, error) {
	m.data = append(m.data, d...)
	return len(d), nil
}

func (m *mockResponseWriter) WriteHeader(statusCode int) {
	m.statusCode = statusCode
}

func TestAddLogger(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		resp []byte
	}{
		// TODO: Add test cases.
		{
			name: "just_run_through",
			args: args{
				r: func() *http.Request {
					r, _ := http.NewRequest(http.MethodGet, "http://example.com", nil)
					return r
				}(),
			},
			resp: nil,
		},
		{
			name: "just_run_through_https",
			args: args{
				r: func() *http.Request {
					r, _ := http.NewRequest(http.MethodGet, "https://example.com", nil)
					r.TLS = &tls.ConnectionState{}
					return r
				}(),
			},
			resp: nil,
		},
		{
			name: "just_run_through_healtcheck_header",
			args: args{
				r: func() *http.Request {
					r, _ := http.NewRequest(http.MethodGet, "http://example.com", nil)
					r.Header = http.Header{
						"X-Liveness-Probe": []string{
							"Healthz",
						},
					}
					return r
				}(),
			},
			resp: nil,
		},
	}

	h := http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {

	})

	applog.Init("info", time.Now().Format(labels.RFC3339Milli), "TDBackend")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockW := &mockResponseWriter{}

			AddLogger(applog.Log, h).ServeHTTP(mockW, tt.args.r)

			assert.Equal(t, tt.resp, mockW.data, "Written response is as expected")
		})
	}
}
