package main

import (
	"time"
	"fmt"
	"math/rand"
	"sync"
)

func multiplex(input ...<-chan interface{}) (<-chan interface{}) {
	result := make(chan interface{})
	// errChannel := make(chan error, 1)

	wg := &sync.WaitGroup{}

	for _, v := range input {
		wg.Add(1)
		go func(i <-chan interface{}) {
			for item := range i {
				result <- item
			}
			defer wg.Done()
		}(v)
	}

	// Put the wait group in a go routine.
	// By putting the wait group in the go routine we ensure either all pass
	// and we close the "finished" channel or we wait forever for the wait group
	// to finish.
	//
	// Waiting forever is okay because of the blocking select below.

	go func ()  {
		wg.Wait()
		close(result)
	}()

	// This select will block until one of the two channels returns a value.
	// This means on the first failure in the go routines above the errChannel will release a
	// value first. Because there is a "return" statement in the err check this function will
	// exit when an error occurs.
	//
	// Due to the blocking on wg.Wait() the finished channel will not get a value unless all
	// the go routines before were successful because not all the wg.Done() calls would have
	// happened.
	// select {
	// case <-result:
	// case err := <-errChannel:
	// 	if err != nil {
	// 		fmt.Println("error ", err)
	// 		return make(chan interface{})
	// 	}
	// }

	return result

}

func track(name string) <-chan any {
    c := make(chan any)
    go func() {
       // Simulate random race time for a horse
       d := time.Duration(rand.Intn(1e2)) * time.Millisecond
       time.Sleep(d)
       // End simulation
       c <- fmt.Sprintf("%s %d%s", name, d/1e6, "ms")
	   close(c)
    }()

    return c
}

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(" - Race starts. - \n ")

	intChan := make(chan any, 2)
	intChan <- 999
	intChan <- "ooo"
	close(intChan)

	nilChan := make(chan any, 1)
	nilChan <- nil
	close(nilChan)

	c := multiplex(track("bchelste"), track("artem"), track("hello"), intChan, nilChan)
	for {

		tmp, ok := <-c
		if !ok {
			break
		}
		fmt.Println(tmp)
	}

	fmt.Println("\n -- Race ends. --")

}