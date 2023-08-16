package main

import "testing"

func TestType1(t *testing.T) {
	api := NewApi(1)
	s := api.Say("Tom")
	if s != "Hi,Tom" {
		t.Fatal("Type1 test fail")
	}
}

func TestType2(t *testing.T) {
	api := NewApi(2)
	s := api.Say("Tom")
	if s != "Hi,Tom" {
		t.Fatal("Type1 test fail")
	}
}
