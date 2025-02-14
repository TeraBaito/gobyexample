package main

import (
	"fmt"
	"time"
)

func main() {
	// tickers do something repeatedly every x amount of time
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)
	
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("tick at", t)
			}
		}
	}()

	time.Sleep (1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("ticker stopped")

	// you can restart a ticker by calling ticker.Start()
}
