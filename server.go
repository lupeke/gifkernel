package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

type respWriter struct {
	http.ResponseWriter
	status int
}

func (rw *respWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func logHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &respWriter{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(wrapped, r)

		log.Printf(
			"%s %s %s %d %s %s",
			r.RemoteAddr,
			r.Method,
			r.URL.Path,
			wrapped.status,
			time.Since(start),
			r.UserAgent(),
		)
	})
}

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC)

	dir := "www"
	fs := http.FileServer(http.Dir(dir))
	http.Handle("/", logHandler(fs))

	addr := ":8888"
	log.Printf("Starting server on %s", addr)
	println("Use Ctrl+C to stop")

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
