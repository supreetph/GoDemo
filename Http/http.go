package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/1"
	res, err := http.Get(url)
	if err != nil {
		panic((err))
	}
	databyte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	responseActual := string(databyte)
	fmt.Println(responseActual)

}
