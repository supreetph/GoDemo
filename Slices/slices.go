package main

import (
	"fmt"
	"sort"
)

func main() {

	var names = []string{}
	var nameList = append(names, "tom", "harry", "sally")
	sort.Strings(nameList) //sorting
	fmt.Println(nameList)
}
