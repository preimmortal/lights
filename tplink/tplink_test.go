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
	//result := Send("192.168.1.105", "{\"system\":{\"get_sysinfo\":{}}}")
	result := Send("192.168.1.105", "{\"system\":{\"set_relay_state\":{\"state\":1}}}")
	fmt.Println(result)
}
