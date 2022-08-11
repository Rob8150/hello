package main

import (
	"testing"
)

func TestHelloGo(t *testing.T) {
	expected := "Welcome to Go Lang"
	actual := HelloGo()
	if actual != expected {
		t.Errorf("%s, %s, Hello Test Failed", expected, actual)
	}
}

func TestAdd(t *testing.T) {
	expected := 11
	actual := Add(5, 6)
	if actual != expected {
		t.Errorf("%d, %d, Add Test Failed", expected, actual)
	}

}

func TestAdd2(t *testing.T) {
	expected := 6
	actual := Add2(1, 2, 3)
	if actual != expected {
		t.Errorf("%d, %d, Add2 Test Failed", expected, actual)
	}

}

func TestSwap(t *testing.T) {
	expected1, expected2 := 6, 3
	actual1, actual2 := Swap(3, 6)
	if actual1 != expected1 {
		t.Errorf("%d, %d, Swap Test Failed", expected1, actual1)
	}
	if actual2 != expected2 {
		t.Errorf("%d, %d, Swap Test Failed", expected2, actual2)
	}

}

func TestSwap2(t *testing.T) {
	expected1, expected2 := 6, 3
	actual1, actual2 := Swap(3, 6)
	if actual1 != expected1 {
		t.Errorf("%d, %d, Swap2 Test Failed", expected1, actual1)
	}
	if actual2 != expected2 {
		t.Errorf("%d, %d, Swap2 Test Failed", expected2, actual2)
	}

}
