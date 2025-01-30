package main

import "fmt"

func main() {
	// make a channel that can send up to 2 strings
	// without them being received
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}