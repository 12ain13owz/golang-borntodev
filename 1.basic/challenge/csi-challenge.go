package main

import "fmt"

var n_tail_sum int

func main() {
	testinput := []int{1, 2, 3, 4, 5}
	k := 0
	n_tail_sum = isTailSum(testinput, k)

	fmt.Println(n_tail_sum)

	type b = string
	var a b = "Test"
	fmt.Println(a)
}

func isTailSum(testinput []int, k int) int {
	length := len(testinput)
	sum := 0
	n_tail_sum := 0
	n := 0

	for i := length; i > 0; i-- {
		sum += testinput[i-1]
		n++

		if sum == k {
			if n_tail_sum == 0 {
				n_tail_sum = n
			} else if n_tail_sum > n {
				n_tail_sum = n
			}
		}
	}

	return n_tail_sum
}

func isTailSum_1(testinput []int, k int) int {
	length := len(testinput)
	sum := 0
	n_tail_sum := 0

	for i := length; i > 0; i-- {
		sum += testinput[i-1]
		n_tail_sum++

		if sum == k {
			break
		}
	}

	if sum != k {
		return 0
	}

	return n_tail_sum
}
