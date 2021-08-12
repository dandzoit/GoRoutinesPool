package mylib

import (
	"sync"
	"testing"
)

func TestInitPool(t *testing.T) {
	var pools [10]*GoRoutinesPool
	var maxNumOfRoutines int16 = 5
	var wg sync.WaitGroup
	for i := 0; i < len(pools); i++ {
		wg.Add(1)
		go func(idx int) {
			pools[idx] = InitPool(2, maxNumOfRoutines)
			defer wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 1; i < len(pools); i++ {
		if pools[0] != pools[i] {
			t.Error("Go routine pool is not singleton")
		}
	}

}
