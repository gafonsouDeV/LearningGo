package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(resWriter http.ResponseWriter, req *http.Request) {
		start := time.Now()

		next.ServeHTTP(resWriter, req)

		log.Printf("%s %s %s", req.Method, req.RequestURI, time.Since((start)))
	})
}
