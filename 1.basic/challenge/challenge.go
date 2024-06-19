package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getCount() int {
	fmt.Printf("%v", "ต้องการคำนวณกี่ตัวเลข: ")
	input, _ := reader.ReadString('\n')
	count, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil {
		message, _ := fmt.Scanf("%v must number only", count)
		panic(message)
	}
	return count
}

func getValue(i int) float64 {
	fmt.Printf("%v", "Number "+strconv.Itoa(i)+": ")
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		message, _ := fmt.Scanf("%v must number only", value)
		panic(message)
	}

	return value
}

func getOperator() string {
	fmt.Printf("Operator : ")
	op, _ := reader.ReadString('\n')

	return strings.TrimSpace(op)
}

func add(value1, value2 float64) float64 {
	return value1 + value2
}

func subtract(value1, value2 float64) float64 {
	return value1 - value2
}

func multiply(value1, value2 float64) float64 {
	return value1 * value2
}

func divide(value1, value2 float64) float64 {
	return value1 / value2
}

func calculate(operator string, value1, value2 float64) float64 {
	result := 0.00

	switch operator {
	case "+":
		result = add(value1, value2)

	case "-":
		result = subtract(value1, value2)

	case "*":
		result = multiply(value1, value2)

	case "/":
		result = divide(value1, value2)
	default:
		panic("wrong operator")
	}

	return float64(result)
}

func main() {
	i := 1
	operator := ""
	value := 0.00
	count := getCount()
	result := getValue(1)

	for {
		i++
		operator = getOperator()
		value = getValue(i)
		result = calculate(operator, result, value)

		if i >= count {
			break
		}
	}
}
