package tplink

import (
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	encodedString := Encrypt("hello")
	expectedBytes := string([]byte{0, 0, 0, 5, 195, 166, 202, 166, 201})

	if string(encodedString) != string(expectedBytes) {
		t.Fatalf("Encoded string is incorrect ")
	}
}

func TestDecrypt(t *testing.T) {
	encodedBytes := []byte{195, 166, 202, 166, 201}
	expectedString := "hello"

	decryptedString := Decrypt(encodedBytes)

	if decryptedString != expectedString {
		t.Fatalf("Decrypted string is incorrect")
	}
}

func TestSend(t *testing.T) {
	lightIPAddress := "192.168.1.105"

	//result, err := Send(lightIPAddress, "{\"system\":{\"get_sysinfo\":{}}}")
	result, err := Send("192.168.1.105", "{\"system\":{\"set_relay_state\":{\"state\":1}}}")
	if err != nil {
		t.Fatalf("Could not execute info command")
	}

	fmt.Println(result)

	// Test  Bad IP
	bad_ip_result, err := Send("999.999.999.999", "{\"system\":{\"get_sysinfo\":{}}}")
	if bad_ip_result != "" || err == nil {
		t.Fatalf("Expected command not to work")
	}

	// Test  Bad Call
	bad_data_result, err := Send(lightIPAddress, "")
	if bad_data_result != "" || err == nil {
		t.Fatalf("Expected command not to work")
	}

}
