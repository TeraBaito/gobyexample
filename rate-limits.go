package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("----- rate limit 1")

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	
	// every 200 millis, accept a request
	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	fmt.Println("----- rate limit 2")
	// maybe we want to allow bursts?
	// make a buffered channel for the limiter
	// up to 3 events every 2 millis
	burstsLimiter := make(chan time.Time, 3)
	
	// simulating initial fill
	for i := 0; i < 3; i++ {
		burstsLimiter <- time.Now()
	}
	// every 200 millis, add another request
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstsLimiter <- t
		}
	}()

	burstsRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstsRequests <- i
	}

	close(burstsRequests)

	// thus, we will receive 3 instantly, then the next every 200 ms
	for req := range burstsRequests {
		<-burstsLimiter
		fmt.Println("request", req, time.Now())
	}
}
