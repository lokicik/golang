package main

import (
	"fmt"
	"sort"
)

// basic types (numbers, strings, booleans)
var myInt int
var myUint uint
var myFloat float32
var myFloat64 float64

// aggregate types (array, struct)
var myStrings [3]string

type Car struct {
	NumberOfTires int
	Luxury        bool
	BucketSeats   bool
	Make          string
	Model         string
	Year          int
}

// myCar := Car{
// 	NumberOfTires: 4,
// 	Luxury:        true,
// 	BucketSeats:   true,
// 	Make:          "Volvo",
// 	Model:         "XC90",
// 	Year:          2019,
// }

// reference types (pointers, slices, maps, functions, channels)

// // pointers
// func main() {
// 	x := 10

// 	myFirstPointer := &x

// 	fmt.Println("x is", x)
// 	fmt.Println("myFirstPointer is", myFirstPointer)

// 	*myFirstPointer = 15

// 	fmt.Println("x is now", x)

// 	changeValueOfPointer(&x)

// 	fmt.Println("After function call, x is now", x)
// }

// func changeValueOfPointer(num *int) {
// 	*num = 25
// }

// // slices
func DeleteFromSlice(a []string, i int) []string {
	a[i] = a[len(a)-1]
	a[len(a)-1] = ""
	a = a[:len(a)-1]
	sort.Strings(a)
	return a
}

// functions
func sumMany(nums ...int) int {
	total := 0

	for _, x := range nums {
		total = total + x
	}

	return total
}

type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs string
}

func (a *Animal) Says() {
	fmt.Printf("A %s says %s", a.Name, a.Sound)
}

// var dog Animal
// dog.Name = "Dog"
// dog.Sound = "Bark"
// dog.Says()

// interface type

func main() {

}
