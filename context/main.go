package main

import (
	"context"
	"fmt"
	"time"
)

// A context is a value you pass to functions that carries 3 things:

// A stop signal, which means “cancel this work”.
// A deadline, which means “stop by this time”.
// Small request data, like a request ID or user ID.

func main() {
	// Let's create first context which is Timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("worker: stop!")
				return
			default:
				fmt.Println("worker: working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	<-ctx.Done()                       // main waits until 2 seconds pass
	time.Sleep(100 * time.Millisecond) // give the worker time to print "stop"
	fmt.Println("main: done")
}
