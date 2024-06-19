package main

import "os"

func main() {

	// เขียนไฟล์ ทับไฟล์เดิม
	data1 := []byte("Hello\nWorld2")
	err := os.WriteFile("data.txt", data1, 0644)

	if err != nil {
		panic(err)
	}

	// สร้างไฟล์
	f, err2 := os.Create("employeeName")

	if err2 != nil {
		panic(err2)
	}

	defer f.Close()

	// สร้างไฟล์พร้อมใส่ Data
	data2 := []byte("Sira\nManee")
	os.WriteFile("employeeName.txt", data2, 0644)

	// เขียนไฟล์ต่อจากของเดิม
	f3, err3 := os.OpenFile("employeeName.txt", os.O_APPEND|os.O_WRONLY, 0644)

	if err3 != nil {
		panic(err3)
	}

	defer f3.Close()
	newData := []byte("\nPitak")
	_, err4 := f3.Write(newData)

	if err4 != nil {
		panic(err4)
	}

}
