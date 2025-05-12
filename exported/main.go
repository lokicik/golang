package main

import (
	"log"
	"myapp/staff"
)

var employees = []staff.Employee{
	{FirstName: "John", LastName: "A", Salary: 30000, FullTime: true},
	{FirstName: "Sally", LastName: "B", Salary: 50000, FullTime: true},
	{FirstName: "Mark", LastName: "C", Salary: 60000, FullTime: true},
	{FirstName: "Allan", LastName: "D", Salary: 15000, FullTime: false},
	{FirstName: "Margaret", LastName: "E", Salary: 100000, FullTime: false},
}

func main() {
	myStaff := staff.Office{
		AllStaff: employees,
	}

	myStaff.All()

	log.Println(myStaff.All())
	staff.OverPaidLimit = 30000
	log.Println("Overpaid staff", myStaff.Overpaid())
	log.Println("Underpaid staff", myStaff.Underpaid())

}
