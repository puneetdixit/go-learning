package main

import (
	"fmt"
	// "time"
	"sync"
)


var (
	mu sync.Mutex
	count int
)

func increment(wg *sync.WaitGroup) {
	mu.Lock()
	count ++
	mu.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	for i:=1; i <= 10; i++ {
		wg.Add(1)
		go increment(&wg)

	}
	wg.Wait()
	fmt.Println("Final count : ", count)
}