package middleware

import (
	labels "github.com/rehabaam/TDBackend/localization"

	"net/http"
	"strings"
	"time"

	"go.uber.org/zap"
)

var (
	logger *zap.Logger
)

// AddLogger logs request/response pair
func AddLogger(logger *zap.Logger, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// We do not want to be spammed by Kubernetes health check.
		// Do not log Kubernetes health check.
		// You can change this behavior as you wish.
		if r.Header.Get("X-Liveness-Probe") == "Healthz" {
			h.ServeHTTP(w, r)
			return
		}

		id := GetReqID(ctx)

		// Prepare fields to log
		var scheme string
		if r.TLS != nil {
			scheme = "https"
		} else {
			scheme = "http"
		}
		proto := r.Proto
		method := r.Method
		remoteAddr := r.RemoteAddr
		userAgent := r.UserAgent()
		uri := strings.Join([]string{scheme, "://", r.Host, r.RequestURI}, "")

		// Log HTTP request
		logger.Debug("request started",
			zap.String(labels.RequestID, id),
			zap.String(labels.HTTPScheme, scheme),
			zap.String(labels.HTTPProto, proto),
			zap.String(labels.HTTPMethod, method),
			zap.String(labels.RemoteAddr, remoteAddr),
			zap.String(labels.UserAgent, userAgent),
			zap.String(labels.URI, uri),
		)

		t1 := time.Now()

		h.ServeHTTP(w, r)

		// Log HTTP response
		logger.Debug("request completed",
			zap.String(labels.RequestID, id),
			zap.String(labels.HTTPScheme, scheme),
			zap.String(labels.HTTPProto, proto),
			zap.String(labels.HTTPMethod, method),
			zap.String(labels.RemoteAddr, remoteAddr),
			zap.String(labels.UserAgent, userAgent),
			zap.String(labels.URI, uri),
			zap.Float64("elapsed-ms", float64(time.Since(t1).Nanoseconds())/1000000.0),
		)
	})
}
