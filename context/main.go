package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// A context is a value you pass to functions that carries 3 things:

// A stop signal, which means “cancel this work”.
// A deadline, which means “stop by this time”.
// Small request data, like a request ID or user ID.
//

func main() {
	// timeOut()
	withCancel()
}

func timeOut() {
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

func withCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	found := make(chan int, 3)

	for i := 1; i <= 3; i++ {
		go func(id int) {
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("worker %d: stopped\n", id)
					return
				default:
					n := rand.Intn(100)
					fmt.Printf("worker %d: checked %d\n", id, n)

					if n == 42 {
						found <- id
						cancel() // tell everyone to stop
						return
					}
					time.Sleep(50 * time.Millisecond)
				}
			}
		}(i)
	}

	winner := <-found
	fmt.Printf("\nworker %d found 42!\n\n", winner)

	time.Sleep(100 * time.Millisecond) // let the others print "stopped"
}
