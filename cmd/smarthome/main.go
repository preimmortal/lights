package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/preimmortal/smarthome"
)

var SmartHomeDB *smarthome.Database

type Device struct {
	Name string `json:"name"`
	Ip   string `json:"ip"`
	Port string `json:"port"`
}

type Devices []Device

func GetDevices(w http.ResponseWriter, r *http.Request) {
	it, err := SmartHomeDB.ReadAll()
	if err != nil {
		glog.Errorf("Could not read from db: %v", err)
	}

	data := Devices{}
	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(*smarthome.DBScan)
		fmt.Printf("  %s - %s - %s\n", p.Name, p.Ip, p.Port)
		device := Device{Name: p.Name, Ip: p.Ip, Port: p.Port}
		data = append(data, device)
	}

	glog.Info("Endpoint Hit: Get Devices endpoint")
	json.NewEncoder(w).Encode(data)
}

func PostDevices(w http.ResponseWriter, r *http.Request) {
	devices := Devices{
		Device{Name: "Post Name", Ip: "1.1.1.1", Port: "9999"},
	}

	glog.Info("Endpoint Hit: Post Devices endpoint")
	json.NewEncoder(w).Encode(devices)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Homepage Endpoint hit")
	glog.Info(w, "Homepage Endpoint hit")
}

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage)
	router.HandleFunc("/devices", GetDevices).Methods("GET")
	router.HandleFunc("/devices", PostDevices).Methods("POST")
	glog.Fatal(http.ListenAndServe(":8081", router))
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: example -stderrthreshold=[INFO|WARN|FATAL] -log_dir=[string]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func init() {
	flag.Usage = usage
	// NOTE: This next line is key you have to call flag.Parse() for the command line
	// options or "flags" that are defined in the glog module to be picked up.
	flag.Parse()
}

func main() {
	glog.Info("Starting Scanner and Server")

	// Initialize Database
	SmartHomeDB = &smarthome.Database{}
	if err := SmartHomeDB.Init(); err != nil {
		glog.Fatal("Could not initialize SmartHome Database")
	}

	//Start the Scanner
	scanner := &smarthome.Scan{Db: SmartHomeDB}
	go scanner.Start()

	//Start Web Server
	handleRequests()
}
