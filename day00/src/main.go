package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"math"
	"sort"
)

func findMode(table map[int64]int64) (float64) {

	if (len(table) == 0) {
		return (0)
	}
	
	var value int64 = 0
	var key int64 = math.MaxInt64

	for i, j := range table {
		if (j > value) {
			value = j
			key = i;
		} else if (j == value) {
			if (i < key) {
				key = i
			}
		}
	}
	return float64(key)
}

func findMedian(storage []int64) (float64) {

	if (len(storage) == 0) {
		return (0)
	} else if (len(storage) == 1) {
		return (float64(storage[0]))
	}


	var pos int
	// fmt.Println(pos)
	// fmt.Println(pos & 1)
	// if (pos & 1) != 0 {
	// 	return (float64(storage[pos]))
	// }
	// return (float64(storage[pos] + storage[pos - 1]) / 2)

	pos = len(storage) / 2
	if (len(storage) & 1) == 0 {
		return (float64(storage[pos] + storage[pos - 1]) / 2)
	}
	return (float64(storage[pos]))
}

func findSD(storage []int64, mean float64) (float64) {

	if (len(storage) == 0) {
		return (0)
	}
		
	var sum float64 = 0;

	for i := range storage {
		sum += math.Pow((float64(storage[i]) - mean), 2)
	}

	return (math.Sqrt(sum / float64(len(storage))))
}

func outputdata(key string, value float64) {
	fmt.Print(key + ": ")
	s := fmt.Sprintf("%.2f", value)
	fmt.Println(s)
}

func conclusion(args []string, storage []int64, table map[int64]int64, sum int64) {

	var mean float64
	if (len(storage) != 0) {
		mean = float64(sum) / float64(len(storage))
	}
	median := findMedian(storage)
	mode := findMode(table)
	sd := findSD(storage, mean)

	if (len(args) == 1) {
		outputdata("Mean", mean)
		outputdata("Median", median)
		outputdata("Mode", mode)
		outputdata("SD", sd)
	} else {
		for item := range args {
			switch (args[item]) {
			case "-mean":
				outputdata("Mean", mean)
			case "-med":
				outputdata("Median", median)
			case "-mod":
				outputdata("Mode", mode)
			case "-sd":
				outputdata("SD", sd)
			}
		}
	}
}
	
func main() {
	fmt.Println("Please put the sequence of integers\ninside of [-100000, 100000] bounds ->")
	
	scanner := bufio.NewScanner(os.Stdin)
	storage := []int64{}
	table := make(map[int64]int64)
	var sum int64 = 0

	for (scanner.Scan()){
		if len(scanner.Text()) == 0 {
			break
		}
		input, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if (err != nil) {
			fmt.Println("Input error! Try again")
			return
		} else {
			if ((input > 100000) || (input < -100000)) {
				fmt.Println("The input number out of bounds!")
				return
			}
			storage = append(storage, input)
			table[input]++
			sum += input
		}
	}
	sort.SliceStable(storage, func (i, j int) bool { return(storage[i] < storage[j])})
	
	conclusion(os.Args, storage, table, sum)
}