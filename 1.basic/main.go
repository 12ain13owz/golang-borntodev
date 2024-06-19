package main

import (
	"fmt"
)

var numberInt, numberInt2, numberInt3 int = 1000, 2000, 3000
var msg string = "hello"

func main() {
	numberFloat := 25.4
	msgString := "hell world"

	fmt.Println(numberInt, numberInt2, numberFloat, msg, msgString)
	// fmt.Println(float64(numberInt) + float64(numberInt2) + numberfloat)
	// fmt.Println(msg + msg + " world"  + strconv.Itoa(numberInt) )
}
