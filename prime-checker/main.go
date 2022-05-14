package main

import "fmt"

func main() {
	fmt.Println(primeCheck(49))
}

func primeCheck(num int) bool {
	if num < 2 {
		return false
	} else if num == 2 {
		return true
	} else {
		for i := 2; i < num; i++ {
			if num%i == 0 {
				return false
			}
		}
		return true
	}

}
