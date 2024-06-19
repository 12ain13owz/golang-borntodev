package main

import (
	"fmt"
	"strings"
)

func main() {
	test := "http://localhost:4200/recipes/3"
	arr := strings.Split(test, "recipes/")
	fmt.Println(len(arr[1:]), arr[len(arr)-1])

	// fmt.Println(arr[len(arr)-1])
	// fmt.Println(arr[0:])
	// fmt.Println(arr[1:])
	// fmt.Println(arr[2:])
	// fmt.Println(arr[len(arr)-1])
	//  fmt.Println(len(arr[1:]))

	// value := arr[len(arr)-1]

}
