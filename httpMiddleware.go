package main

import (
	"net/http"

	"fmt"
)





// basic http features.
func basic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// write logs etc here
		fmt.Println("basic")
		next.ServeHTTP(w, r)
	})
}



// do auth check
func auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// do auto check here

		fmt.Println("auth")
		next.ServeHTTP(w, r)
	})
}
