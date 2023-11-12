package middleware

import (
	"context"
	"fmt"
	"net/http"
	"sync/atomic"
)

// Key to use when setting the request ID.
type ctxKeyRequestID int

// RequestIDKey is the key that holds th unique request ID in a request context.
const RequestIDKey ctxKeyRequestID = 0

var (
	// prefix is const prefix for request ID
	prefix string

	// reqID is counter for request ID
	reqID uint64
)

// AddRequestID is a middleware that injects a request ID into the context of each
// request. A request ID is a string of the form "host.example.com/random-0001",
// where "random" is a base62 random string that uniquely identifies this go
// process, and where the last number is an atomically incremented request
// counter.
func AddRequestID(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		myid := atomic.AddUint64(&reqID, 1)
		ctx := r.Context()
		ctx = context.WithValue(ctx, RequestIDKey, fmt.Sprintf("%s-%06d", prefix, myid))
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetReqID returns a request ID from the given context if one is present.
// Returns the empty string if a request ID cannot be found.
func GetReqID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if reqID, ok := ctx.Value(RequestIDKey).(string); ok {
		return reqID
	}
	return ""
}
