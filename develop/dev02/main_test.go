package main

import (
	"testing"
)

func TestUnpack1_Case1(t *testing.T) {
	input := "a4bc2d5e"
	expected := "aaaabccddddde"
	result, err := unpackString(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestUnpack1_Case2(t *testing.T) {
	input := "abcd"
	expected := "abcd"
	result, err := unpackString(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestUnpack1_Case3(t *testing.T) {
	input := "45"
	expected := ""
	result, err := unpackString(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestUnpack1_Case4(t *testing.T) {
	input := ""
	expected := ""
	result, err := unpackString(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestUnpack1_Case5(t *testing.T) {
	input := "qwe\\4\\5"
	expected := "qwe45"
	result, err := unpackString(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestUnpack1_Case6(t *testing.T) {
	input := "qwe\\45"
	expected := "qwe44444"
	result, err := unpackString(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestUnpack1_Case7(t *testing.T) {
	input := "qwe\\\\5"
	expected := "qwe\\\\\\\\\\"
	result, err := unpackString(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
