package main

import (
	"fmt"
	"time"
)

func main() {
	for i := range 10 {
		go func() {
			fmt.Println(i)
		}()
	}

	time.Sleep(10 * time.Millisecond)

	ch := make(chan string)
	go func() {
		ch <- "text"
	}()
	msg := <-ch
	fmt.Println(msg)

	go func() {
		for i := range 3 {
			msg := fmt.Sprintf("message: #%d", i+1)
			fmt.Println("sending")
			ch <- msg
		}
		close(ch)
	}()

	for msg := range ch {
		fmt.Println("receiving")
		fmt.Println("got: ", msg)
	}

	msg = <-ch
	fmt.Printf("value from closed channel: %#v \n", msg)
}
