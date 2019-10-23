package main
import (
"github.com/gorilla/handlers"
"github.com/gorilla/mux"
"log"
"os"
"net/http"
)

func mainLogic(w http.ResponseWriter, r *http.Request) {
		log.Println("Processing request!")
		w.Write([]byte("OK"))
		log.Println("Finished processing request")
}

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/", mainLogic)
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	log.Fatalln(http.ListenAndServe(":8080",loggedRouter))
}