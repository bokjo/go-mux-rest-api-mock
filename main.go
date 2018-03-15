package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, " Welcome greeting message...")
}

func requestsHandler() {
	router := mux.NewRouter()

	router.HandleFunc("/", welcome)

	log.Fatal(http.ListenAndServe(":9999", router))

}

func main() {

	fmt.Println("GO MUX REST API MOCK")

	requestsHandler()

	fmt.Println("END OF GO MUX REST API MOCK... Exiting...")
}
