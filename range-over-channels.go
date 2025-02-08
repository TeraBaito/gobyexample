package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	// even when you close the channel, you can fetch its messages
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
