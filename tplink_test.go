package smarthome

import (
	"log"
	"testing"
)

func TestEncrypt(t *testing.T) {
	tp := Tplink{}
	encodedString := tp.encrypt("hello")
	expectedBytes := string([]byte{0, 0, 0, 5, 195, 166, 202, 166, 201})

	if string(encodedString) != string(expectedBytes) {
		t.Fatalf("Encoded string is incorrect ")
	}
}

func TestDecrypt(t *testing.T) {
	tp := Tplink{}
	encodedBytes := []byte{195, 166, 202, 166, 201}
	expectedString := "hello"

	decryptedBytes := tp.decrypt(encodedBytes)
	decryptedString := string(decryptedBytes)

	if decryptedString != expectedString {
		t.Fatalf("Decrypted string is incorrect")
	}
}

func Test_Positive_Send_Info(t *testing.T) {
	tp := Tplink{}
	sn := Scan{}
	ip, err := sn.FindFirstIP()
	if err != nil {
		t.Fatalf("Could not find any IP")
	}
	result, err := tp.Send(ip, "{\"system\":{\"get_sysinfo\":{}}}")
	if err != nil {
		t.Fatal("Could not execute info command: ", err)
	}

	log.Print(string(result))
}

//func Test_Positive_Send_PlugOn(t *testing.T) {
//	tp := Tplink{}
//	sn := Scan{}
//	ip, err := sn.FindFirstIP()
//	if err != nil {
//		t.Fatalf("Could not find any IP")
//	}
//	result, err := tp.Send(ip, "{\"system\":{\"set_relay_state\":{\"state\":1}}}")
//	if err != nil {
//		t.Fatalf("Could not execute info command")
//	}
//
//	log.Print(result)
//}

//func Test_Positive_Send_PlugOff(t *testing.T) {
//	tp := Tplink{}
//	sn := Scan{}
//	ip, err := sn.FindFirstIP()
//	if err != nil {
//		t.Fatalf("Could not find any IP")
//	}
//	result, err := tp.Send(ip, "{\"system\":{\"set_relay_state\":{\"state\":0}}}")
//	if err != nil {
//		t.Fatalf("Could not execute info command")
//	}
//
//	log.Print(result)
//}

func Test_Negative_Send_BadIP(t *testing.T) {
	tp := Tplink{}
	t.Log("Negative Test - Bad IP")
	// Test  Bad IP
	ip_result, err := tp.Send("999.999.999.999", "{\"system\":{\"get_sysinfo\":{}}}")
	bad_ip_result := string(ip_result)
	t.Logf("Got IP %s, Err %v", bad_ip_result, err)
	if bad_ip_result != "" || err == nil {
		t.Fatalf("Expected command not to work")
	}
}

func Test_Negative_Send_BadCall(t *testing.T) {
	tp := Tplink{}
	sn := Scan{}
	ip, err := sn.FindFirstIP()
	if err != nil {
		t.Fatalf("Could not find any IP")
	}
	t.Log("Negative Test - Bad Call")
	// Test  Bad Call
	data_result, err := tp.Send(ip, "")
	bad_data_result := string(data_result)
	if bad_data_result != "" || err == nil {
		t.Fatalf("Expected command not to work")
	}
}
