package main

import "log"

func main() {
	log.Println(factorialRecursion(9))
	log.Println(factorialLoop(9))
}

func factorialRecursion(num int) int {
	if num == 1 || num == 0 {
		return 1
	}
	fact := factorialRecursion(num - 1)
	fact *= num
	return fact
}

func factorialLoop(num int) int {
	fact := 1
	for i := 1; i <= num; i++ {
		fact *= i
	}
	return fact
}
