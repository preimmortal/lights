package main

import (
	"encoding/json"
	"errors"
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

func GetDeviceInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting device info")
	vars := mux.Vars(r)
	tp := smarthome.Tplink{}
	fmt.Println("\tDevice: ", vars["deviceip"])
	info, err := tp.Send(vars["deviceip"], smarthome.TPLINK_API_INFO)
	if err != nil {
		glog.Errorf("Could not send info request: %v", err)
	}
	glog.Info("Endpoint Hit: Post Device Info endpoint")
	json.NewEncoder(w).Encode(info)
}

type API_CONTRACT_PostDeviceAction struct {
	State string `json:"state"`
}

func PostDeviceAction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Posting device action")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	api := &API_CONTRACT_PostDeviceAction{}
	action := ""

	//Extract Body information
	if err := json.NewDecoder(r.Body).Decode(api); err != nil {
		glog.Error(err)
	}

	fmt.Println("Got Action: ", api.State)

	switch api.State {
	case "on":
		action = smarthome.TPLINK_API_RELAY_ON
	case "off":
		action = smarthome.TPLINK_API_RELAY_OFF
	default:
		glog.Error(errors.New("state must be defined as \"on\" or \"off\""))
	}

	//Initialize TPLink Obj
	tp := smarthome.Tplink{}
	fmt.Println("\tDevice: ", vars["deviceip"])
	fmt.Println("\tAction: ", action)

	// Send Action
	info, err := tp.Send(vars["deviceip"], action)
	if err != nil {
		glog.Errorf("Could not send info request: %v", err)
	}
	glog.Info("Endpoint Hit: Post Device Action endpoint")
	json.NewEncoder(w).Encode(info)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Homepage Endpoint hit")
	glog.Info(w, "Homepage Endpoint hit")
}

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage)
	router.HandleFunc("/devices", GetDevices).Methods("GET")
	router.HandleFunc("/devices/{deviceip}", GetDeviceInfo).Methods("GET")
	router.HandleFunc("/devices/{deviceip}", PostDeviceAction).Methods("POST")
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
