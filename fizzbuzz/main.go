package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(fizzBuzz(3))
}

func fizzBuzz(num int) error {
	if num <= 0 {
		return errors.New("invalid input")
	} else {
		switch {
		case num%15 == 0:
			fmt.Println("fizzbuzz")
			return nil
		case num%3 == 0:
			fmt.Println("fizz")
			return nil
		case num%5 == 0:
			fmt.Println("buzz")
			return nil
		default:
			fmt.Println(num)
			return nil
		}
	}
}
