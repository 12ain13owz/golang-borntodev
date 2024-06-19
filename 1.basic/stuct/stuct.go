package main

import "fmt"

type Employee struct {
	employeeID   string
	employeeName string
	phone        string
}

func main() {
	// Ex.1 struct
	// employee := Employee{
	// 	employeeID:   "101",
	// 	employeeName: "Pradoo",
	// 	phone:        "0909874517",
	// }

	// fmt.Println(employee)

	// -----------------------------------------------------------------------------------

	// Ex.2 struct + array
	// employeeList := [3]Employee{}
	// employeeList[0] = Employee{
	// 	employeeID:   "101",
	// 	employeeName: "Pradoo",
	// 	phone:        "090-854-8746",
	// }
	// employeeList[1] = Employee{
	// 	employeeID:   "102",
	// 	employeeName: "Prayad",
	// 	phone:        "090-854-8746",
	// }
	// employeeList[2] = Employee{
	// 	employeeID:   "103",
	// 	employeeName: "Pranee",
	// 	phone:        "090-854-8746",
	// }

	// fmt.Println("Employee List =", employeeList)

	// -----------------------------------------------------------------------------------

	// Ex.3 strict + slice
	employeeList := []Employee{}
	employee1 := Employee{
		employeeID:   "101",
		employeeName: "Pradoo",
		phone:        "090-845-8874",
	}
	employee2 := Employee{
		employeeID:   "102",
		employeeName: "Prayad",
		phone:        "090-845-8874",
	}
	employee3 := Employee{
		employeeID:   "103",
		employeeName: "Pranee",
		phone:        "090-845-8874",
	}

	employeeList = append(employeeList, employee1)
	employeeList = append(employeeList, employee2)
	employeeList = append(employeeList, employee3)

	fmt.Println("Employee List =", employeeList)
}
