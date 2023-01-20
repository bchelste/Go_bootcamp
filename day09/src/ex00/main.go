package main

import (
	"sync"
	"time"
	"fmt"
)

func addNumber(x int, wg *sync.WaitGroup, output *chan int) {

	var tmp int = 1
	defer wg.Done()
	if (x < 0) {
		tmp = -1
	}
	time.Sleep(time.Duration(x * tmp) * time.Millisecond)
	*output <- x
}

func sleepSort(numbers []int) (<-chan int){
	result := make(chan int, len(numbers))
	resultNeg := make(chan int, len(numbers))
	var wg sync.WaitGroup
	var wgNeg sync.WaitGroup
	
	for _, x := range numbers {
		if (x < 0) {
			wgNeg.Add(1)
			go addNumber(x, &wgNeg, &resultNeg)
		} else {
			wg.Add(1)
			go addNumber(x, &wg, &result)
		}
	}

	wg.Wait()
	wgNeg.Wait()
	close(result)
	close(resultNeg)

	final := make(chan int, len(numbers))
	var tmp []int

	for item := range resultNeg {
		tmp = append(tmp, item)
	}
	for i := len(tmp) - 1; i >= 0; i-- {
		final <- tmp[i]
	}
	for item := range result {
		final <- item
	}

	close(final)
	return final
}

func main() {
	testArray := []int{3, -99, 7, 9, 16, -9, -16, -1}
	fmt.Println(testArray)
	res := sleepSort(testArray)
	fmt.Print("The result is: ")
	for item := range res {
		fmt.Printf("%d ", item)
	}
	fmt.Println()
	fmt.Println("--")

	testArray2 := []int{3, 3, 3, 4, 5, 3, 3}
	fmt.Println(testArray2)
	res2 := sleepSort(testArray2)
	fmt.Print("The result is: ")
	for item := range res2 {
		fmt.Printf("%d ", item)
	}
	fmt.Println()
	fmt.Println("--")

	testArray3 := []int{}
	fmt.Println(testArray3)
	res3 := sleepSort(testArray3)
	fmt.Print("The result is: ")
	for item := range res3 {
		fmt.Printf("%d ", item)
	}
	fmt.Println()
	fmt.Println("--")
}