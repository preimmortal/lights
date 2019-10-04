package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/preimmortal/smarthome"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint hit")
}

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage)
	router.HandleFunc("/devices", smarthome.GetDevices).Methods("GET")
	router.HandleFunc("/devices", smarthome.PostDevices).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	fmt.Println("Starting Server")
	//Start a background goroutine to scan the network
	db := &smarthome.Database{}
	scanner := &smarthome.Scan{Db: db}
	go scanner.Start()
	handleRequests()
	fmt.Println("Hello World")
}
