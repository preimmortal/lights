package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/preimmortal/lights"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint hit")
}

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage)
	router.HandleFunc("/lights", lights.GetLights).Methods("GET")
	router.HandleFunc("/lights", lights.PostLights).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	handleRequests()
	fmt.Println("Hello World")
}
