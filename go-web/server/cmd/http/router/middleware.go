package router

import (
	"fmt"
	"net/http"
	"os"
	"server/pkg/response"
	"time"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	StatusCode int
	Bytes      int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, 0, 0}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.StatusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func (lrw *loggingResponseWriter) Write(b []byte) (int, error) {
	if lrw.StatusCode == 0 {
		lrw.StatusCode = http.StatusOK
	}
	n, err := lrw.ResponseWriter.Write(b)
	lrw.Bytes += n
	return n, err
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != os.Getenv("AUTH") {
			response.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()

		lrw := NewLoggingResponseWriter(w)
		defer func() {
			scheme := "http"
			if r.TLS != nil {
				scheme = "https"
			}
			fmt.Printf("%s %s://%s%s - %d %dB in %s \n", r.Method, scheme, r.Host, r.RequestURI, lrw.StatusCode, lrw.Bytes, time.Since(t1))
		}()

		next.ServeHTTP(lrw, r)
	})
}
