package main

import "fmt"

func main() {
	var courseName []string
	courseName = []string{"Java", "Phyton"}
	fmt.Println(courseName)

	courseName = append(courseName, "C", "C#", "HTML", "CSS", "Javascript") // เพิ่ม Data
	fmt.Println(courseName)

	courseWeb := courseName[4:7] // ดึง Data 3 ตัวหลังคือ ต่ำแหน่ง 5, 6, 7
	fmt.Println(courseWeb)

	courseWeb = courseName[:4] // ดึง Data 4 ตัวแรกคือ ต่ำแหน่ง 1, 2, 3, 4
	fmt.Println(courseWeb)
}
