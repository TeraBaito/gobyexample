// with worker pools we can put various workers to work on the
// same channel of jobs concurrently
// they take the jobs like piñata: each one takes one and works on it

package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		// do work 
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		// place result
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// start 3 workers, initially they are blocked since
	// there are still no jobs
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// add jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// collect results
	for a := 1; a <= numJobs; a++ {
		fmt.Println(<-results)
	}
}
