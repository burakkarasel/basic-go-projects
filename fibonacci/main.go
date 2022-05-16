package main

import (
	"errors"
	"fmt"
)

func main() {
	firstTen, err := fibo(10)

	fmt.Println(firstTen, err)
}

func fibo(num int) ([]int, error) {

	if num < 2 {
		return []int{1}, errors.New("Invalid input")
	}

	num1 := 0
	num2 := 1

	var result []int

	result = append(result, num1)
	result = append(result, num2)

	for i := 2; i < num; i++ {
		total := num1 + num2
		num1 = num2
		num2 = total

		result = append(result, num2)
	}
	return result, nil
}
