package mylib

import "sync"

type AtomicInt struct {
	counter int16
	lock    sync.Mutex
}

func (ai *AtomicInt) getCount() int16 {
	ai.lock.Lock()
	defer ai.lock.Unlock()
	return ai.counter
}

func (ai *AtomicInt) increment() {
	ai.lock.Lock()
	defer ai.lock.Unlock()
	ai.counter++
}

func (ai *AtomicInt) decrement() {
	ai.lock.Lock()
	defer ai.lock.Unlock()
	ai.counter--
}