package gmp

import "testing"

func TestGetNearestPath(t *testing.T) {
	res, err := GetNearestPath()
	if err != nil {
		t.Fatal(err)
	}
	println(res)
}

func TestGetPath(t *testing.T) {
	res, err := GetPath(0)
	if err != nil {
		t.Fatal(err)
	}
	println(res)
}
