package main

import (
	"DIY_2/mylib"
	"fmt"
	"time"
)

func main() {
	pool := mylib.InitPool(4, 2)
	pool2 := mylib.InitPool(4, 4)

	pool.Submit(mylib.Task{
		Id:      1,
		Fn:      printTillMillion,
		Timeout: time.Second * 5,
	})

	pool2.Submit(mylib.Task{
		Id:      2,
		Fn:      printTillMillion,
		Timeout: time.Second * 10,
	})
	pool.Submit(mylib.Task{
		Id:      3,
		Fn:      printTillMillion,
		Timeout: time.Second * 10,
	})
	pool.Submit(mylib.Task{
		Id:      4,
		Fn:      printTillMillion,
		Timeout: time.Second * 10,
	})
	time.Sleep(time.Minute * 1)
}

func printTillMillion() {
	for i := 0; i < 2*1e1; i++ {
		fmt.Println(i)
	}
}
