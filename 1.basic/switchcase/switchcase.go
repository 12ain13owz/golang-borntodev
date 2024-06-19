package main

import "fmt"

var color, hex string

func main() {
	// input := 10
	// switch input {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// default:
	// 	fmt.Println("default")
	// }

	// fmt.Scanf("%s", &color) รับเฉพาะตัวอักษร
	// fmt.Scand("%d", &color) รับเฉพาะตัวเลข
	fmt.Scan(&color)

	switch color {
	case "blue":
		hex = "#0000FF"
	case "green":
		hex = "#008000"
	case "pink":
		hex = "#FFC0CB"
	case "yellow":
		hex = "#FFFF00"
	default:
		hex = "No color"
	}

	fmt.Println(color, "to", hex)

}
