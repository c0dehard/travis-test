package main

import "testing"

func TestAdd(t *testing.T) {
	four := add(2, 2)
	if four != 4 {
		t.Log("False")
		t.Fail()
	} else {
		t.Logf("'four' had the value %d as expected", four)

	}
}
