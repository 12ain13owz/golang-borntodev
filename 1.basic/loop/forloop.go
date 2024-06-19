package main

import "fmt"

const count int = 15

func main() {
	i := 0

	for i < 10 {
		// fmt.Println("i =", i)
		i++
	}

	for j := 0; j < count; j++ {
		// fmt.Println("j = ", j)
	}

	var input string
	for {
		fmt.Scanf("%s", &input)
		fmt.Println("Input =", input)

		if input == "exit" {
			break
		}
	}
}
