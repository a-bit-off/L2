package main

import (
	"log"
	"reflect"
	"testing"
)

func TestCase1_k_1(t *testing.T) {
	args := &Args{lcK: 3, filePath: "files/case1_k.txt"}
	data, err := getData(args.filePath)
	if err != nil {
		log.Fatalln(err)
	}
	result := mySort(args, data)
	expected := []string{
		"2 7",
		"6 3",
		"0 9 1",
		"4 5 2",
		"8 1 3",
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestCase1_k_2(t *testing.T) {
	args := &Args{lcK: 2, filePath: "files/case1_k.txt"}
	data, err := getData(args.filePath)
	if err != nil {
		log.Fatalln(err)
	}
	result := mySort(args, data)
	expected := []string{
		"8 1 3",
		"6 3",
		"4 5 2",
		"2 7",
		"0 9 1",
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestCase2_k(t *testing.T) {
	args := &Args{lcK: 1, filePath: "files/case2_k.txt"}
	data, err := getData(args.filePath)
	if err != nil {
		log.Fatalln(err)
	}
	result := mySort(args, data)
	expected := []string{
		"012С",
		"345D",
		"678A",
		"902B",
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestCase3_n_r_1(t *testing.T) {
	args := &Args{lcN: true, filePath: "files/case3_n_r.txt"}
	data, err := getData(args.filePath)
	if err != nil {
		log.Fatalln(err)
	}
	result := mySort(args, data)
	expected := []string{
		"1",
		"2",
		"3",
		"4",
		"4",
		"5",
		"6",
		"9",
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestCase3_n_r_2(t *testing.T) {
	args := &Args{lcR: true, filePath: "files/case3_n_r.txt"}
	data, err := getData(args.filePath)
	if err != nil {
		log.Fatalln(err)
	}
	result := mySort(args, data)
	expected := []string{
		"9",
		"6",
		"5",
		"4",
		"4",
		"3",
		"2",
		"1",
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestCase4_u(t *testing.T) {
	args := &Args{lcU: true, filePath: "files/case4_u.txt"}
	data, err := getData(args.filePath)
	if err != nil {
		log.Fatalln(err)
	}
	result := mySort(args, data)
	expected := []string{
		"123",
		"223",
		"2344",
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestCase5_b(t *testing.T) {
	args := &Args{lcB: true, filePath: "files/case5_b.txt"}
	data, err := getData(args.filePath)
	if err != nil {
		log.Fatalln(err)
	}
	result := mySort(args, data)
	expected := []string{
		"line 1",
		" line 2",
		"  line 3",
		"line 4",
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestCase6_h(t *testing.T) {
	args := &Args{lcH: true, filePath: "files/case6_h.txt"}
	data, err := getData(args.filePath)
	if err != nil {
		log.Fatalln(err)
	}
	result := mySort(args, data)
	expected := []string{
		"500B",
		"1K",
		"100M",
		"10G",
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestCase7_M(t *testing.T) {
	args := &Args{ucM: true, filePath: "files/case7_M.txt"}
	data, err := getData(args.filePath)
	if err != nil {
		log.Fatalln(err)
	}
	result := mySort(args, data)
	expected := []string{
		"Январь",
		"Март",
		"Апрель",
		"Декабрь",
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
