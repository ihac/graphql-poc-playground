package common

import (
	"net/http"
	"time"
)

func LatencyMiddleware(latencyInMills int, next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Duration(latencyInMills) * time.Millisecond)
		next.ServeHTTP(rw, r)
	})
}
