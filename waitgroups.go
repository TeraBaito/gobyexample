// with waitgroups we can wait for various workers to finish

package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("worker %d starting\n", id)
	// do work
	time.Sleep(time.Second)
	fmt.Printf("worker %d finished\n", id)
}

func main() {
	// note: if you wanna pass this on functions, do so as a pointer
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		// increment wait group counter
		wg.Add(1)
		
		// wrap in closure to make it async without
		// actually making the function async
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}
	// wait until wait group counter is 0
	wg.Wait()
	fmt.Println("all done!")
}
