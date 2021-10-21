package main

import "fmt"

func main() {

	user1 := Person{1, "Supreet", "Bangalore", 1674631764, true}
	fmt.Println(user1)
}

type Person struct {
	id            int
	Name          string
	City          string
	ContactNumber int
	IsActive      bool
}
