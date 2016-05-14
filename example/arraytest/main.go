package main

import (
	"GoRepositories/Microsoft"
	"fmt"
)

// Employees test
type Employees []Employee

// Employee model for DigitalFleet Sql database
type Employee struct {
	EmployeeID int
	ClientID   int
	LastName   string `sql:"Last_Name"`
	FirstName  string `sql:"First_Name"`
	Login      string
	PIN        string
}

func main() {
	fmt.Printf("\nemployee array length start\n")
	connectionString := "server=192.168.0.157;user id=chad;password=ky1ttk#1;database=DigitalFleet;port=1433"
	userPin := "440440"
	query := "select EmployeeID,ClientID,Last_Name,First_Name, LOGIN, PIN from employees where PIN = " + userPin
	rows, err := Microsoft.GetByQuery(connectionString, query)

	if err != nil {
		fmt.Println(fmt.Errorf("sql error: %v \n", err))
		// os.Exit(1)
	} else {
		defer rows.Close()
		employees := Employees{}
		employeeCount := len(employees)

		fmt.Printf("employee array length : %d \n", employeeCount)
		if len(employees) == 0 {
			fmt.Printf("2-Semployee array length : %d \n", employeeCount)
		}
	}
	fmt.Printf("\nemployee array length ended\n")
}
