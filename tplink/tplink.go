package tplink

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
)

func Encrypt(unenc string) []byte {
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

func Decrypt(enc []byte) string {
	key := 171
	result := []byte("")
	for _, c := range enc {
		a := key ^ int(c)
		key = int(c)
		result = append(result, byte(a))
	}
	return string(result)
}

func Send(ip string, command string) (string, error) {
	port := "9999"
	address := net.JoinHostPort(ip, port)

	conn, err := net.Dial("tcp", address)
	if err != nil {
		return "", errors.New("Could not connect to address")
	}
	defer conn.Close()
	fmt.Fprintf(conn, string(Encrypt(command)))

	result, err := ioutil.ReadAll(conn)
	if err != nil || len(result) == 0 {
		return "", errors.New("Could not read data back: ")
	}

	return Decrypt(result[4:]), nil
}
