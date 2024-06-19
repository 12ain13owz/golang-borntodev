package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	ID           int
	EmployeeName string
	Tel          string
	Email        string
}

// json.Marshal แปลง Go เป็น Json
func main() {
	data, _ := json.Marshal(&employee{101, "Sirasit Boonklang", "098-784-8888", "sirasak@gmail.com"})

	fmt.Println(string(data))
}
