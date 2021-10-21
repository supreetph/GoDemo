package main

import "fmt"

func main() {

	names := make(map[string]string)
	names["fistname"] = "tom"
	names["lastname"] = "hanks"
	fmt.Println(names)
	fmt.Println(names["fistname"])
}
