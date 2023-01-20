package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
)

type mutexInput struct {
    mu sync.Mutex
    data <-chan string
}

func crawlWeb(input <-chan string, ctx context.Context) <-chan *string  {

	var tokens = make(chan struct{}, 8)

	mutex := mutexInput{data: input}
	var wg sync.WaitGroup
	result := make(chan *string)
	
	for item := range input {
		wg.Add(1)
		go func (src string) {
			defer wg.Done()
			select {
			case <- ctx.Done():
				return
			default:
				mutex.mu.Lock()
				tmp, err := http.Get(src)
				mutex.mu.Unlock()

				tokens <- struct{}{}

				if err != nil {
					log.Fatalf("some error with http.Get: %v\n", err)
				}
				body, err := ioutil.ReadAll(tmp.Body)
				if err != nil {
					log.Fatalf("some error with ioutil.ReadAll: %v\n", err)
				}

				res := string(body)
				result <- &res

				<-tokens
			}

		}(item)
	}
	
	go func ()  {
		wg.Wait()
		close(result)
	}()

	// wg.Wait()
	// close(result)

	return result
}

func chanCreat(nbr int) <-chan string {
	result := make(chan string, nbr)
	for i := 0; i < nbr; i++ {
		result <- "http://abracadabra_"
	}
	close(result)

	return result
}

func main() {

	fmt.Println("crawl Web was started")
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	go func(){
    	<-sigChan
		fmt.Println("\n- - - CTRL + C was pressed - - -")
		stop()
		os.Exit(1)
	}()

	tmp := chanCreat(10)

	result := crawlWeb(tmp, ctx)
	for {
		item, ok := <-result
		if !ok {
			break
		}
		fmt.Printf("%s\n", *item)
	}
}