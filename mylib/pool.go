package mylib

import (
	"log"
	"runtime"
	"sync"
	"time"
)

type GoRoutinesPool struct {
	maxNumOfRoutines      int16
	inQ                   chan Task
	currExecutingRoutines AtomicInt
}

var pool *GoRoutinesPool
var once sync.Once

func InitPool(numOfCPUs int, maxNumOfRoutines int16) *GoRoutinesPool {
	once.Do(func() {
		p := &GoRoutinesPool{
			maxNumOfRoutines: maxNumOfRoutines,
			inQ:              make(chan Task),
			currExecutingRoutines: AtomicInt{
				counter: 0,
				lock:    sync.Mutex{},
			},
		}
		pool = p
		log.Println("Initiating pool")
		go pool.executor()
		runtime.GOMAXPROCS(numOfCPUs)
	})
	return pool
}

type Task struct {
	Id      int16
	Fn      func()
	Timeout time.Duration
}

func (pool *GoRoutinesPool) Submit(task Task) {
	pool.inQ <- task
}

func (pool *GoRoutinesPool) executor() {
	for t := range pool.inQ {
		//new task is available
		//check whether we are allowed to take this up or not
		for {
			log.Println("Trying to start task:", t.Id)
			currCount := pool.currExecutingRoutines.getCount()
			if currCount < pool.maxNumOfRoutines {
				log.Println("Picking task:", t.Id)
				pool.currExecutingRoutines.increment()
				go func(task Task) {
					log.Println("Starting task:", t.Id)
					task.Fn()
					pool.currExecutingRoutines.decrement()
					log.Println("Completed task:", t.Id)
				}(t)
				break
			}
			log.Println("Task q is full, waiting to start task:", t.Id)
		}

	}
}
