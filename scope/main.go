package main

import (
	"myapp/packageone"
)

var myVar = "myVar"


func main() {
	var blockVar = "blockVar"
	packageone.PrintMe(myVar, blockVar)
}
