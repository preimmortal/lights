package lights

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/preimmortal/lights/scan"
)

type Light struct {
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Status string `json:"status"`
}

type Lights []Light

func GetLights(w http.ResponseWriter, r *http.Request) {
	local_ip := "192.168.1.0/24"
	//result, err := tplink.Send("192.168.1.105", "{\"system\":{\"get_sysinfo\":{}}}")
	result, err := scan.Scan(local_ip)
	if err != nil {
		fmt.Println("got an error", err)
	}
	fmt.Println(result)

	fmt.Println("Endpoint Hit: Get Lights endpoint")
	json.NewEncoder(w).Encode(result)
}

func PostLights(w http.ResponseWriter, r *http.Request) {
	lights := Lights{
		Light{Name: "Post Name", Desc: "Post Description", Status: "On"},
	}

	fmt.Println("Endpoint Hit: Post Lights endpoint")
	json.NewEncoder(w).Encode(lights)
}
