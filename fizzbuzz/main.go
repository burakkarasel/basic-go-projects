package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(fizzBuzz(7))
}

func fizzBuzz(num int) (string, error) {
	if num <= 0 {
		return "invalid", errors.New("invalid input")
	} else {
		switch {
		case num%15 == 0:
			return "fizzbuzz", nil
		case num%3 == 0:
			return "fizz", nil
		case num%5 == 0:
			return "buzz", nil
		default:
			return fmt.Sprintf("%v", num), nil
		}
	}
}
