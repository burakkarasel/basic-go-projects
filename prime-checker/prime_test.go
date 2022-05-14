package main

import "testing"

var tests = []struct {
	Input  int
	Output bool
	Err    bool
}{
	{-3, false, true},
	{2, true, false},
	{4, false, false},
	{5, true, false},
	{9, false, false},
	{47, true, false},
	{49, false, false},
	{149, true, false},
	{331, true, false},
	{816, false, false},
	{831, false, false},
	{827, true, false},
}

// test folder is for some of my cases

// and here with TestPrime func i test my cases by iterating tests

func TestPrime(t *testing.T) {
	for _, tt := range tests {
		result, err := primeCheck(tt.Input)
		if tt.Err {
			if err == nil {
				t.Error("Expected an error but didnt get one")
			}
		} else {
			if err != nil {
				t.Error("Did not expect error but got :", err)
			} else if result != tt.Output {
				t.Errorf("Expected %v but got %v", tt.Output, result)
			}
		}
	}
}
