package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("indexinfo.csv")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	count := 1

	for scanner.Scan() {
		line := scanner.Text()
		item := strings.Split(line, ",")

		// fmt.Printf("Line %d : %s\n", count, line)
		fmt.Println(item[3])
		count++
	}
}
