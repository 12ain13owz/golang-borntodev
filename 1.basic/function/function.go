package main

import "fmt"

func hello() {
	fmt.Println("Hello BorntoDev")
}

func plus(value1 int, value2 int) int {
	// result := value1 + value2
	// fmt.Println("result =", result)
	return value1 + value2
}

func plus3value(value1, value2, value3 int) int {
	return value1 + value2 + value3
}

func main() {
	hello()
	plus(10, 20)

	result := plus(10, 20)
	fmt.Println("result =", result)

	result2 := plus3value(100, 100, 150)
	fmt.Println("result =", result2)
}
