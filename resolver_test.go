package smarthome

import "testing"

func TestInit(t *testing.T) {
	t.Log("Testing Resolver Initialization")
	r := resolver{}
	if err := r.Init(); err != nil {
		t.Fatal("Could not initialize resolver: ", err)
	}
	t.Log("Done Testing Resolver Initialization")
}
