package main

import (
	"fmt"
	"time"
)

func main() {
	// create an event as a channel
	timer1 := time.NewTimer(2 * time.Second)

	// block code below until it receives the timeout signal
	<-timer1.C
	fmt.Println("timer 1 fired!")

	// if you just wanna wait like above, you can use time.Sleep
	// you can cancel timer before firing
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer 2 fired!")
	}()
	// stop timer2
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer 2 stopped")
	}
	time.Sleep(2 * time.Second)
}
