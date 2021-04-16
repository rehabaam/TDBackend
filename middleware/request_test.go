package middleware

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddRequestID(t *testing.T) {
	type args struct {
		reqID uint64
	}
	type resp struct {
		requestID string
	}

	tests := []struct {
		name string
		args args
		exp  resp
	}{
		{
			name: "success_652",
			args: args{
				reqID: 652,
			},
			exp: resp{
				requestID: prefix + "-000653",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotReq *http.Request
			h := http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
				gotReq = r
			})

			reqID = tt.args.reqID
			r, _ := http.NewRequest(http.MethodGet, "http://example.com", nil)
			AddRequestID(h).ServeHTTP(nil, r)

			gotID := GetReqID(gotReq.Context())

			assert.Equal(t, tt.exp.requestID, gotID, "RequestID from context is as expected")
		})
	}
}

func TestGetReqID(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ctx_nil",
			args: args{},
		},
		{
			name: "ctx_does_not_contain_id",
			args: args{
				ctx: context.Background(),
			},
		},
		{
			name: "ctx_contains_id",
			args: args{
				ctx: context.WithValue(context.Background(), RequestIDKey, "123"),
			},
			want: "123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetReqID(tt.args.ctx)

			assert.Equal(t, tt.want, got, "GetReqID returned value as expected")
		})
	}
}
