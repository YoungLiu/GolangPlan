package main

import (
	"fmt"
)

var tableVar = map[string]int{
	"one": 1,
	"two": 2,
	"three": 3,
}

var tableFunction = map[int]functionType{
	3: isOdd,
	4: isEven,
}

type functionType func(int) bool

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

func main() {
	var arrayVar []int
	for k, v := range tableVar {
		fmt.Printf("before into closure: k->%s, v->%d\n", k, v)
		func() {
			arrayVar = append(arrayVar, v)
		}()
	}
	fmt.Println(arrayVar)

	var arrayFunc []func()
	for k, v := range tableFunction {
		fmt.Printf("before into closure: k->%d, v->%d\n", k, v)
		func() {
			//arrayFunc = append(arrayFunc, func(){v(k)})
			arrayFunc = append(arrayFunc, {v(k)})
		}()
	}
	fmt.Println(arrayFunc)

}
