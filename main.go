package main

import (
	"fmt"
	"sync"
)

var counter int
var mutex sync.Mutex

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		// opening a new goroutine to increment the counter variable that is global
		go incremental(&wg)
	}

	wg.Wait()
	fmt.Println("Final counter:", counter)
}

func incremental(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mutex.Lock()
		// lock before add a value to the counter
		counter++
		// unlock after the operation being done
		mutex.Unlock()
	}
}
