package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Scratch() {
	// Show this when talking about channel management
	// And talk about go routine resource/memory leak
	newStream := func(stop chan bool) <-chan int {
		stream := make(chan int)
		go func() {
			defer fmt.Println("routine closed...")
			defer close(stream)
			for {
				select {
				case <-stop:
					return
				case stream <- rand.Int():
				}
			}
		}()
		return stream
	}
	var stop = make(chan bool)
	randomStream := newStream(stop)
	for i := 0; i < 10; i++ {
		fmt.Printf("random number %d : %d\n", i, <-randomStream)
	}
	close(stop)
	time.Sleep(1 * time.Second)
}
