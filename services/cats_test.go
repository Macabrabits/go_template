package services

import (
	"testing"
)

func TestTest(t *testing.T) {
	param := "world of tests"
	msg := toTest(param)
	expected := "hello world of tests"
	if msg != expected {
		t.Errorf("toTest(%q) = %q; want %q", param, msg, expected)
	}

}

// func TestGetCats(t *testing.T) {
// 	msg, err := GetCats()

// }

// func TestCreateCat() {

// }
