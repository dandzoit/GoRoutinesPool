package mylib

import (
	"fmt"
	"testing"
	"time"
)

func TestSubmit(t *testing.T) {
	pool := InitPool(2, 2)
	pool.Submit(Task{
		Id:      1,
		Fn:      printTillMillion,
		Timeout: time.Second * 5,
	})

	pool.Submit(Task{
		Id:      2,
		Fn:      printTillMillion,
		Timeout: time.Second * 5,
	})

	//Tasks #1 and #2 will be picked up instantly. Tasks #3 and #4
	//will be queued up

	pool.Submit(Task{
		Id:      3,
		Fn:      printTillMillion,
		Timeout: time.Second * 5,
	})

	pool.Submit(Task{
		Id:      4,
		Fn:      printTillMillion,
		Timeout: time.Second * 5,
	})
}

func printTillMillion() {
	for i := 0; i < 1*1e2; i++ {
		fmt.Println(i)
	}
}
