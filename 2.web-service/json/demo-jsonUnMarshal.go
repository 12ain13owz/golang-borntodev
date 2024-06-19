package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	ID           int
	EmployeeName string
	Tel          string
	Email        string
}

// json.Marshal แปลง Json เป็น Go
func main() {
	e := employee{}
	err := json.Unmarshal([]byte(`{"ID": 101, "EmployeeName": "Sirasit Boonklang", "Tel": "098-784-8888", "Email": "sirasak@gmail.com"}`), &e)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(e.ID)
}
