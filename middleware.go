package main

import (
	"log"
	"net/http"
	"runtime/debug"
)

func PanicRecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				log.Printf(":( Panic: %v", err)
				log.Println(string(debug.Stack()))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
