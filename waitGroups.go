// package main

// import (
// 	"fmt"
// 	"sync"
// 	// "time"
// )

// func worker(id int, wg *sync.WaitGroup) {
//     fmt.Printf("Worker %d starting\n", id)
//     // Simulate work
//     // time.Sleep(time.Second)
//     fmt.Printf("Worker %d done\n", id)
//     wg.Done()
// }

// func main() {
// 	var wg sync.WaitGroup

// 	for i:= 1; i<= 10; i++ {
// 		wg.Add(1)
// 		go worker(i, &wg)
// 	}
// 	wg.Wait()
// 	fmt.Println("All worker done")
// }