package smarthome

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

type TplinkInfo struct {
	System struct {
		GetSysInfo struct {
			ErrorCode  int     `json:"err_code"`
			SwVersion  string  `json:"sw_ver"`
			HwVersion  string  `json:"hw_ver"`
			Type       string  `json:"type"`
			Model      string  `json:"model"`
			Mac        string  `json:"mac"`
			DeviceId   string  `json:"deviceId"`
			HwID       string  `json:"hwId"`
			FwID       string  `json:"fwId"`
			OemId      string  `json:"oemId"`
			Alias      string  `json:"alias"`
			DevName    string  `json:"dev_name"`
			IconHash   string  `json:"icon_hash"`
			RelayState int     `json:"relay_state"`
			OnTime     int     `json:"on_time"`
			ActiveMode string  `json:"active_mode"`
			Feature    string  `json:"feature"`
			Updating   int     `json:"updating"`
			Rssi       int     `json:"rssi"`
			LedOff     int     `json:"led_off"`
			Latitude   float32 `json:"latitude"`
			Longitude  float32 `json:"longitude"`
		} `json:"get_sysinfo"`
	} `json:"system"`
}
type Tplink struct{}

const (
	TPLINK_API_PORT      = "9999"
	TPLINK_API_PORT_INT  = 9999
	TPLINK_API_INFO      = "{\"system\":{\"get_sysinfo\":{}}}"
	TPLINK_API_RELAY_ON  = "{\"system\":{\"set_relay_state\":{\"state\":1}}}"
	TPLINK_API_RELAY_OFF = "{\"system\":{\"set_relay_state\":{\"state\":0}}}"
)

func (t Tplink) encrypt(unenc string) []byte {
	key := 171
	result := make([]byte, 4)
	binary.BigEndian.PutUint32(result, uint32(len(unenc)))
	for _, c := range unenc {
		a := key ^ int(c)
		key = a
		result = append(result, byte(a))
	}
	return result
}

func (t Tplink) decrypt(enc []byte) []byte {
	key := 171
	result := []byte("")
	for _, c := range enc {
		a := key ^ int(c)
		key = int(c)
		result = append(result, byte(a))
	}
	return result
}

func (t Tplink) Send(ip, command string) ([]byte, error) {
	log.Printf("Sending command \"%s\"", command)
	address := net.JoinHostPort(ip, TPLINK_API_PORT)
	log.Printf("\tAddress: \"%s\" ", address)

	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	fmt.Fprintf(conn, string(t.encrypt(command)))

	result, err := ioutil.ReadAll(conn)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("No results read")
	}

	json.Marshal(t.decrypt(result[4:]))

	return t.decrypt(result[4:]), nil
}
