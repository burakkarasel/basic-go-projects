package main

import "fmt"

func main() {

	addition := calculator("+", 5, 13)

	subtraction := calculator("-", 23.6, 5.9)

	multiplication := calculator("*", 6, 7.2)

	division := calculator("/", -1, -777)

	fmt.Println(addition, subtraction, multiplication, division)

}

func calculator(operation string, num1 float32, num2 float32) float32 {
	switch operation {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	}
	return 0
}
