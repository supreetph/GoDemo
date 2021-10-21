package main

import "fmt"

func main() {
	//types of variable declaration
	var firstname string = "Supreet"
	var age int = 35
	var isActive bool = true
	height := 5.8
	var certifcation [2]string
	certifcation[0] = "csm"
	certifcation[1] = "pmp"
	//certifcation[2]="cspo"
	fmt.Println("Hello", firstname)
	fmt.Println("age", age)
	fmt.Println("isActive", isActive)
	fmt.Println("height", height)
	fmt.Println("certifcations", certifcation)

}
