package gmp

import (
	"errors"
	"testing"
)

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
	_, err = GetPath(1)
	if !errors.Is(err, FailedToFind) {
		t.Fatal(err)
	}
}

func TestGetFolderPath(t *testing.T) {
	res, err := GetFolderPath("go-mod-path")
	if err != nil {
		t.Fatal(err)
	}
	println(res)
}
