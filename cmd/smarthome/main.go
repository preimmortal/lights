package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/preimmortal/smarthome"
)

var SmartHomeDB *smarthome.Database

type Device struct {
	Key   string `json:"key"`
	Name  string `json:"name"`
	Ip    string `json:"ip"`
	Alias string `json:"alias"`
	State string `json:"state"`
}

type Devices []Device

func GetDevices(w http.ResponseWriter, r *http.Request) {
	it, err := SmartHomeDB.ReadAll()
	if err != nil {
		log.Printf("Could not read from db: %v", err)
	}

	data := Devices{}
	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(*smarthome.DBDevice)
		log.Printf("  %s - %s - %s - %s\n", p.Name, p.Ip, p.Alias, p.State)
		device := Device{Name: p.Name, Ip: p.Ip, Alias: p.Alias, State: p.State}
		data = append(data, device)
	}

	log.Print("Endpoint Hit: Get Devices endpoint")
	json.NewEncoder(w).Encode(data)
}

func GetDeviceInfo(w http.ResponseWriter, r *http.Request) {
	log.Print("Getting device info")
	vars := mux.Vars(r)
	tp := smarthome.Tplink{}
	log.Print("\tDevice: ", vars["deviceip"])
	//info, err := tp.Send(vars["deviceip"], smarthome.TPLINK_API_INFO)
	infoBytes, err := tp.Send(vars["deviceip"], smarthome.TPLINK_API_INFO)
	if err != nil {
		log.Print("Could not get info: ", err)
	}
	var info *smarthome.TplinkInfo
	err = json.Unmarshal(infoBytes, &info)
	if err != nil {
		log.Print("Could not encode the info: ", err)
	}
	log.Print(info)
	if err != nil {
		log.Printf("Could not send info request: %v", err)
	}
	log.Print("Endpoint Hit: Post Device Info endpoint -> ", info)
	json.NewEncoder(w).Encode(info)
}

type API_CONTRACT_PostDeviceAction struct {
	State string `json:"state"`
}

func PostDeviceAction(w http.ResponseWriter, r *http.Request) {
	log.Print("Posting device action")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	api := &API_CONTRACT_PostDeviceAction{}
	action := ""

	//Extract Body information
	if err := json.NewDecoder(r.Body).Decode(api); err != nil {
		log.Print(err)
	}

	log.Print("Got Action: ", api.State)

	switch api.State {
	case "on":
		action = smarthome.TPLINK_API_RELAY_ON
	case "off":
		action = smarthome.TPLINK_API_RELAY_OFF
	default:
		log.Print(errors.New("state must be defined as \"on\" or \"off\""))
	}

	//Initialize TPLink Obj
	tp := smarthome.Tplink{}
	log.Print("\tDevice: ", vars["deviceip"])
	log.Print("\tAction: ", action)

	// Send Action
	info, err := tp.Send(vars["deviceip"], action)
	if err != nil {
		log.Printf("Could not send info request: %v", err)
	}
	log.Print("Endpoint Hit: Post Device Action endpoint")
	json.NewEncoder(w).Encode(info)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	log.Print(w, "Homepage Endpoint hit")
}

func handleRequests() {
	headersOk := handlers.AllowedHeaders([]string{"Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST"})

	router := mux.NewRouter()
	router.HandleFunc("/", homePage)
	router.HandleFunc("/devices", GetDevices).Methods("GET")
	router.HandleFunc("/devices/{deviceip}", GetDeviceInfo).Methods("GET")
	router.HandleFunc("/devices/{deviceip}", PostDeviceAction).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", handlers.CORS(headersOk, originsOk, methodsOk)(router)))
}

func main() {
	log.Print("Starting Scanner and Server")

	// Initialize Database
	SmartHomeDB = &smarthome.Database{}
	if err := SmartHomeDB.Init(); err != nil {
		log.Fatal("Could not initialize SmartHome Database")
	}

	//Start the Scanner
	scanner := &smarthome.Scan{Db: SmartHomeDB}
	go scanner.Start(false)

	//Start Web Server
	handleRequests()
}
