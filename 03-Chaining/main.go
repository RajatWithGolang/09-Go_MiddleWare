package main

import (
	"fmt"
	"net/http"
	//"github.com/justinas/alice"
)

func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middlewareOne")
		next.ServeHTTP(w, r)
		fmt.Println("Executing middlewareOne again")
	})
}
func middlewareTwo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middlewareTwo")
		if r.URL.Path != "/" {
			return
		}
		next.ServeHTTP(w, r)
		fmt.Println("Executing middlewareTwo again")
	})
}
func final(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing finalHandler")
	w.Write([]byte("OK"))
}

func main() {
	finalHandler := http.HandlerFunc(final)
	http.Handle("/", middlewareOne(middlewareTwo(finalHandler)))
	//using alice Package
	// httpm.Handle("/",alice.New(middlewareOne,middlewareOne).ThenFunc(http.HandlerFunc(index )))
	http.ListenAndServe(":8080", nil)
}
