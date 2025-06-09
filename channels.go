// package main

// import (
// 	"fmt"
// 	"time"
// )

// func worker(done chan bool) {
// 	fmt.Println("Working...")
// 	time.Sleep(time.Second)
// 	fmt.Println("Work done")
// 	done<- true
// }

// func worker2(done2 chan string) {
// 	fmt.Println("Worker 2 running...")
// 	done2 <- "Work Done"
// }

// func main() {
// 	done := make(chan bool)
// 	done2 := make(chan string)
// 	go worker2(done2)
// 	go worker(done)
// 	// <- done
// 	fmt.Println("Main end ", <- done)
// 	fmt.Println("Work 2 output ", <- done2)
// }