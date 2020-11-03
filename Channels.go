package main

import "fmt"

func main() {

	message := make(chan string, 2)

	message <- "buffered"
	message <- "channel"

	//go func() { message <- "Go lang Channel" }()

	//msg := <-message
	//fmt.Println(msg)

	fmt.Println(<-message)
	fmt.Println(<-message)

}
