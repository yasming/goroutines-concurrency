package main

import (
	"fmt"
	"sync"
)

type SyncCounter struct {
	sync.Mutex
	Value int
}

var counter SyncCounter

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		// opening a new goroutine to increment the counter variable that is shared
		go incremental(&counter, &wg)
	}

	wg.Wait()
	fmt.Println("Final counter:", counter.Value)
}

func incremental(syncCounter *SyncCounter, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		syncCounter.Lock()
		// lock before add a value to the counter
		syncCounter.Value++
		// unlock after the operation being done
		syncCounter.Unlock()
	}
}
