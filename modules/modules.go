package modules

import (
	"fmt"
)

var x float32 = 1.2

func init() {
    fmt.Printf("p1 package, x:%f\n", x)
}

func Donothing() {
   	fmt.Println("do nothing.\n")
}
