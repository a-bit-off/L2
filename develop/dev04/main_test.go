package main

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	words := []string{"cat", "dog", "tac", "god", "act"}
	expected := map[string][]string{
		"act": {"act", "cat", "tac"},
		"dgo": {"dog", "god"},
	}

	result := findAnagrams(words)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestSortString(t *testing.T) {
	str := "bca"
	expected := "abc"

	result := sortString(str)

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestGetUnique(t *testing.T) {
	words := []string{"cat", "cat", "dog", "dog", "dog"}
	expected := []string{"cat", "dog"}

	getUnique(&words)

	if !reflect.DeepEqual(words, expected) {
		t.Errorf("Expected %v, but got %v", expected, words)
	}
}
