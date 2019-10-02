package database

import (
	"testing"
)

func TestDatabase(t *testing.T) {
	t.Logf("Testing Database Init")
	if err := Init(); err != nil {
		t.Fatalf("Could not create database: %v", err)
	}

	t.Logf("Testing Database Insert")
	err := Insert("TPLink_Plug", "192.168.1.105", "9999")
	if err != nil {
		t.Fatalf("Could not write database: %v", err)
	}

	t.Logf("Testing Database ReadAll")
	it, err := ReadAll()
	if err != nil {
		t.Fatalf("Could not read database: %v", err)
	}

	for obj := it.Next(); obj != nil; obj = it.Next() {
		p := obj.(*Scan)
		t.Logf("  %s - %s - %s\n", p.Name, p.Ip, p.Port)
	}

	check1, err := HasIp("192.168.1.105")
	if err != nil {
		t.Fatalf("Error checking for HasIP: %v", err)
	}
	t.Log(check1)

	check2, err := HasIp("192.168.1.104")
	if err != nil {
		t.Fatalf("Error checking for HasIP: %v", err)
	}
	t.Log(check2)
}
