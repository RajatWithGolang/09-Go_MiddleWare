package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/handlers"
	"os"
	"log"
)


func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing Index MiddleWare")
	fmt.Fprintf(w, "Welcome")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing about MiddleWare")
	fmt.Fprintf(w, "Welcome Rajat")
}
func iconHandler(w http.ResponseWriter, r *http.Request){
	 
}
func main() {  
	http.HandleFunc("/favicon.ico",iconHandler)
	indexHandler := http.HandlerFunc(index)
	aboutHandler := http.HandlerFunc(about)
	http.Handle("/", handlers.LoggingHandler(os.Stdout,handlers.CompressHandler(indexHandler)))
	http.Handle("/about", handlers.LoggingHandler(os.Stdout,handlers.CompressHandler(aboutHandler)))
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
