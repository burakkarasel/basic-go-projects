package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(primeCheck(49))
}

func primeCheck(num int) (bool, error) {
	if num < 2 {
		return false, errors.New("numbers that lesser than 2 cannot be a prime number")
	} else if num == 2 {
		return true, nil
	} else {
		for i := 2; i < num; i++ {
			if num%i == 0 {
				return false, nil
			}
		}
		return true, nil
	}

}
