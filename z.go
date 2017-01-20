package main

import "fmt"

var WhatIsThe7 = AnswerToLife(2.7) //2.7
var WhatIsThe8 = AnswerToLife(2.8) //2.8
var WhatIsThe9 = AnswerToLife(2.9) //2.9

func init() { //3.5


       fmt.Printf("init WhatIsThe in z.go`s init 3.5: %d\n", 4)

}

func init() { //3.6


       fmt.Printf("init WhatIsThe in z.go`s init 3.6: %d\n", 5)

}
