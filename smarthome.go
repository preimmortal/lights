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

func GetLights(w http.ResponseWriter, r *http.Request) {

	lights := Lights{
		Light{Name: "Post Name", Desc: "Post Description", Status: "On"},
	}
	fmt.Println("Endpoint Hit: Get Lights endpoint")
	json.NewEncoder(w).Encode(lights)
}

func PostLights(w http.ResponseWriter, r *http.Request) {
	lights := Lights{
		Light{Name: "Post Name", Desc: "Post Description", Status: "On"},
	}

	fmt.Println("Endpoint Hit: Post Lights endpoint")
	json.NewEncoder(w).Encode(lights)
}
