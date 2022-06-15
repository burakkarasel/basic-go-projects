package main

import "fmt"

func main() {
	fmt.Println(Between(4, 9))
}

func Between(a, b int) []int {
	var mySlice []int
	for i := a; i <= b; i++ {
		mySlice = append(mySlice, i)
	}
	return mySlice
}
