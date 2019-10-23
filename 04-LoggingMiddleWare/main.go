package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Printf("Started %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		fmt.Printf("Completed  %s %v\n", r.URL.Path, time.Since(start))
	})
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing Index MiddleWare")
	fmt.Fprintf(w, "Welcome")
	fmt.Println()

}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing about MiddleWare")
	fmt.Fprintf(w, "Welcome Rajat")
	fmt.Println()
}

func main() {
	indexHandler := http.HandlerFunc(index)
	aboutHandler := http.HandlerFunc(about)
	http.Handle("/", LoggingMiddleware(indexHandler))
	http.Handle("/about", LoggingMiddleware(aboutHandler))
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
