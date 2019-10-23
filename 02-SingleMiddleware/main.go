package main

import (
	"fmt"
	"net/http"
)

func middleWare(next http.Handler) http.Handler {
	fmt.Println("MiddleWare is running before MainLogic")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			next.ServeHTTP(w,r)
		fmt.Println("MiddleWare is running after MainLogic")
	})
	 
}

func mainLogic(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Main Logic handler Running")
	w.Write([]byte("OK"))
}

func main() {
	mainLogicHandler := http.HandlerFunc(mainLogic)
	http.Handle("/",middleWare(mainLogicHandler))
	http.ListenAndServe(":8080",nil)
}
