package main

import "fmt"

func main() {
	result := add(1, 2)
	fmt.Println(result)
}

func add(num1 int, num2 int) int {
	return num1 + num2
}
