package main

import "github.com/YoungLiu/GolangPlan/modules" //1
import "fmt"

var WhatIsThe4 = AnswerToLife(2.4) //2.4
var WhatIsThe5 = AnswerToLife(2.5) //2.5
var WhatIsThe6 = AnswerToLife(2.6) //2.6

func AnswerToLife(index float32) float32 {
       fmt.Printf("init package level variable, WhatIsThe: %f\n", index)
       return index
}

func init() { //3.3

       fmt.Printf("init WhatIsThe in testinit.go`s init3.3: %d\n", 0)

}

func init() { //3.4

       fmt.Printf("init WhatIsThe in testinit.go`s init3.4: %d\n", 1)

}

func main() { //4

       modules.Donothing() //5
}
