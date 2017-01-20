package main

import "fmt"

var WhatIsThe1 = AnswerToLife(2.1) //2.1
var WhatIsThe2 = AnswerToLife(2.2) //2.2
var WhatIsThe3 = AnswerToLife(2.3) //2.3

func init() { //3.1
	fmt.Printf("init WhatIsThe in a.go `s init 3.1: %d\n", 2)
}

func init() { //3.2
	fmt.Printf("init WhatIsThe in a.go`s init 3.2: %d\n", 3)
}
