package main

import (
	"fmt"
	"github.com/robfig/cron"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

var table = map[string]func(string){
	"test": test,
	"for":  another,
}

func main() {
	var c = cron.New()
	var once sync.Once
	once.Do(func() {
		c.Start()
		fmt.Println("Init the cron job daemon")
	})

	//var k string
	//var v func(string)
	//var array []func()
	//for k, v = range table {
	//	array = append(array, func(){v(k)})
	//}
	//k = "hahaha"
	//
	//for i, v := range array {
	//	fmt.Printf("index is %d, address is %d\n", i, v)
	//	c.AddFunc("*/10 * * * * *", v)
	//}

	for k, v := range table {
		//fmt.Printf("closure address is %p\n", func(){v(k)})
		var temp = func(k string, v func(string)) func() {
			var kValue = k
			var vValue = v
			return func() {
				vValue(kValue)
			}
		}(k, v)
		c.AddFunc("*/3 * * * * *", temp)
	}
	//c.AddFunc("*/3 * * * * *", func(){test()})
	//c.AddFunc("*/3 * * * * *", func(){another()})

	fmt.Printf("test01 is %p, test02 is %p", func() {
		fmt.Print("test")
	}, func() {
		fmt.Errorf("error is %s", "test")
	})

	http.ListenAndServe("localhost:16060", nil)
	select {}
}

func test(test string) {

	fmt.Printf("one test is %s\n", test)
}

func another(test string) {
	fmt.Printf("another test is %s\n", test)
}
