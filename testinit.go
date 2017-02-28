package main

import "github.com/YoungLiu/GolangPlan/modules" //1
import (
	"fmt"
	"github.com/kardianos/osext"
	"strings"
)

var WhatIsThe4 = AnswerToLife(2.4) //2.4
var WhatIsThe5 = AnswerToLife(2.5) //2.5
var WhatIsThe6 = AnswerToLife(2.6) //2.6

func AnswerToLife(index float32) float32 {
	fmt.Printf("init package level variable, WhatIsThe: %f\n", index)
	return index
}

func init() {
	//3.3

	fmt.Printf("init WhatIsThe in testinit.go`s init3.3: %d\n", 0)

}

func init() {
	//3.4

	fmt.Printf("init WhatIsThe in testinit.go`s init3.4: %d\n", 1)

}

func main() {
	//4

	modules.Donothing() //5
	executablePath, _ := osext.ExecutableFolder()
	fmt.Println("the executable path is :" + executablePath)

	// test strings.split after N
	test := "liuyang:haah:no:yes"
	res := strings.SplitAfterN(test, ":", 4)
	fmt.Println(res)

	testMap := map[int]string{
		1 : "liuyang",
		2 : "hjaha",
		3 : "jaijsij",
	}
	for k, v := range testMap {
		fmt.Printf("k>>>%d   v>>>%s\n", k, v)
	}

	testSlice := []string{"haha", "liuyang", "hdhhd", "hjdjsijf", "fhuehu"}
	for index, item := range testSlice {
		fmt.Printf("index is %d, value is %s\n", index, item)
	}
}
