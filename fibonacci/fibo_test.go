package main

import "testing"

var tests = []struct {
	input  int
	output []int
	isErr  bool
}{
	{10, []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}, false},
	{15, []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377}, false},
	{16, []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610}, false},
	{-3, []int{1}, true},
}

func TestFibonacci(t *testing.T) {
	for _, tt := range tests {
		got, err := fibo(tt.input)
		if tt.isErr {
			if err == nil {
				t.Error("Expected error but didnt get one")
			}
		} else {
			if err != nil {
				t.Error("Didnt expect an error but got one", err)
			} else if got == nil {
				t.Errorf("Expected %v but got %v", got, nil)
			}
		}
	}
}
