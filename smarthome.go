package smarthome

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Light struct {
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Status string `json:"status"`
}

type Lights []Light

func GetDevices(w http.ResponseWriter, r *http.Request) {

	lights := Lights{
		Light{Name: "Post Name", Desc: "Post Description", Status: "On"},
	}
	fmt.Println("Endpoint Hit: Get Devices endpoint")
	json.NewEncoder(w).Encode(lights)
}

func PostDevices(w http.ResponseWriter, r *http.Request) {
	lights := Lights{
		Light{Name: "Post Name", Desc: "Post Description", Status: "On"},
	}

	fmt.Println("Endpoint Hit: Post Devices endpoint")
	json.NewEncoder(w).Encode(lights)
}
