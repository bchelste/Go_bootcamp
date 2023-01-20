package main

import (
	"fmt"
)

type Present struct {
    Value int
    Size int
}

// func presentMax(first, second Present) Present {
	
// 	if (first.Value == second.Value) {
// 		if (first.Size < second.Size) {
// 			return first
// 		}
// 	}
// 	if (first.Value > second.Value) {
// 		return first
// 	}
// 	return second
// }

func printTable(table *[][]int) {
	for i := 0; i < len(*table); i++ {
		for j := 0; j < len((*table)[i]); j++ {
			fmt.Printf("%d ", (*table)[i][j])
		}
		fmt.Println()
	}
}

func initTable(table *[][]int, capacity int) {
	for i := range *table {
		(*table)[i] = make([]int, capacity + 1)
	}
}

func getSolution(presents *[]Present, keepTable *[][]int, capacity int) []Present {
	var result []Present

	for i := len(*presents); i > 0; i-- {
		if ((*keepTable)[i][capacity] == 1) {
			result = append(result, (*presents)[i - 1])
			capacity -= (*presents)[i - 1].Size
		}
	}

	return result
}

func grabPresents(presents []Present, capacity int) []Present {
	if (capacity <= 0) {
		return nil
	}

	valueTable := make([][]int, (len(presents) + 1))
	keepTable := make([][]int, (len(presents) + 1))
	initTable(&valueTable, capacity)
	initTable(&keepTable, capacity)

	// printTable(&valueTable)
	// fmt.Println("----------")
	// printTable(&keepTable)

	for i := 1; i < (len(presents) + 1); i++ {
		for j := 1; j < (capacity + 1); j++ {
			if (presents[i - 1].Size <= j) {
				prev := valueTable[i - 1][j]
				maxCapacity := presents[i - 1].Value + valueTable[i - 1][j - presents[i - 1].Size]
				if (maxCapacity > prev) {
					valueTable[i][j] = maxCapacity
					keepTable[i][j] = 1
				} else {
					valueTable[i][j] = prev
				}
			}
		}
	}

	printTable(&valueTable)
	fmt.Println("----------")
	printTable(&keepTable)

	return getSolution(&presents, &keepTable, capacity)

}

func main() {

	exampleSlice := []Present{{5, 1}, {4, 5}, {3, 1}, {5, 2}}
	fmt.Println(exampleSlice)
	capacity := 3
	result := grabPresents(exampleSlice, capacity)
	fmt.Println("For capacity =", capacity)
	fmt.Println("The result is:", result)
	fmt.Println("|--------------------------------------|")

	testSlice := []Present{{5, 1}, {4, 5}, {3, 1}, {5, 2}, {4, 1}, {8, 2}, {2, 2}}
	fmt.Println(testSlice)
	capacity = 4
	result = grabPresents(testSlice, capacity)
	fmt.Println("For capacity =", capacity)
	fmt.Println("The result is:", result)
	capacity = -10
	result = grabPresents(testSlice, capacity)
	fmt.Println("For capacity =", capacity)
	fmt.Println("The result is:", result)
	fmt.Println("|--------------------------------------|")

	emptySlice := []Present{}
	capacity = 2
	fmt.Println(emptySlice)
	result = grabPresents(emptySlice, capacity)
	fmt.Println("For capacity =", capacity)
	fmt.Println("The result is:", result)
}