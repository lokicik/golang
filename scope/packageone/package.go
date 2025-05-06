package packageone

import (
	"fmt"
)

var PackageVar = "PackageVar"

func PrintMe(a, b string) {
	fmt.Println(a, b, PackageVar)
}
