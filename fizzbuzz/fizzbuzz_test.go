package main

import (
	"testing"
)

var tests = []struct {
	input  int
	output string
	isErr  bool
}{
	{1, "1", false},
	{3, "fizz", false},
	{5, "buzz", false},
	{7, "7", false},
	{9, "fizz", false},
	{15, "fizzbuzz", false},
	{-3, "invalid", true},
	{-5, "invalid", true},
	{-15, "invalid", true},
	{30, "fizzbuzz", false},
	{42, "fizz", false},
	{50, "buzz", false},
}

func TestFizzBuzz(t *testing.T) {
	for _, tt := range tests {
		got, err := fizzBuzz(tt.input)
		if tt.isErr {
			if err == nil {
				t.Error("Expected an error but didnt got one")
			}
		} else {
			if err != nil {
				t.Error("Didnt expect an error but got one")
			} else if got != tt.output {
				t.Errorf("Expected %v but got %v", tt.output, got)
			}
		}
	}
}
