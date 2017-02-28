package modules

import (
	"fmt"
)

var x float32 = 1.2

func init() {
	fmt.Printf("p1 package, x:%f\n", x)
	x += 1.0
}

func Donothing() {
	fmt.Println("do nothing.")
	fmt.Printf("the x in modules is %f\n", x)
}
