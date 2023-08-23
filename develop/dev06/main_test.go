package main

import (
	"log"
	"testing"
)

func TestCat_1(t *testing.T) {
	args := &Args{
		flags: Flags{
			fields:    0,
			delimiter: "",
			separated: true,
		},
		files: []string{"files/case1.txt"},
	}

	expectedResult := []string{
		"1 2 3",
		"",
		"4",
		"",
		"7 8 9",
	}
	res, err := processFile(args)
	if err != nil {
		log.Fatalln(err)
	}
	result := res[0]

	if len(result) != len(expectedResult) {
		t.Errorf("Expected result length: %d, but got: %d", len(expectedResult), len(result))
	}

	for i, line := range result {
		if line != expectedResult[i] {
			t.Errorf("Expected line: %s, but got: %s", expectedResult[i], line)
		}
	}
}

func TestCat_2(t *testing.T) {
	args := &Args{
		flags: Flags{
			fields:    2,
			delimiter: "",
			separated: false,
		},
		files: []string{"files/case1.txt"},
	}

	expectedResult := []string{
		"2",
		"",
		"",
		"",
		"",
		"8",
	}
	res, err := processFile(args)
	if err != nil {
		log.Fatalln(err)
	}
	result := res[0]

	if len(result) != len(expectedResult) {
		t.Errorf("Expected result length: %d, but got: %d", len(expectedResult), len(result))
	}

	for i, line := range result {
		if line != expectedResult[i] {
			t.Errorf("Expected line: %s, but got: %s", expectedResult[i], line)
		}
	}
}

func TestCat_3(t *testing.T) {
	args := &Args{
		flags: Flags{
			fields:    0,
			delimiter: "\t",
			separated: false,
		},
		files: []string{"files/case1.txt"},
	}

	expectedResult := []string{
		"1	2	3",
		"",
		"4",
		"",
		"",
		"",
		"7	8	9",
	}
	res, err := processFile(args)
	if err != nil {
		log.Fatalln(err)
	}
	result := res[0]

	if len(result) != len(expectedResult) {
		t.Errorf("Expected result length: %d, but got: %d", len(expectedResult), len(result))
	}

	for i, line := range result {
		if line != expectedResult[i] {
			t.Errorf("Expected line: %s, but got: %s", expectedResult[i], line)
		}
	}
}
