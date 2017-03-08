package main

import "fmt"

func main() {
	var test []string
	var another = []string{"test01", "test02"}
	for i := 0; i < 50; i++ {
		test= append(test, another...)
	}
	fmt.Println(test)
}
