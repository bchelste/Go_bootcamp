package main

import (
	"container/heap"
	"errors"
	"fmt"
)

type Present struct {
    Value int
    Size int
}

type PresentHeap struct {
	PresentStorage []Present
}

func (presents PresentHeap) Len() int { 
	return len(presents.PresentStorage)
}

func (presents PresentHeap) Less(i, j int) bool {
	if (i > presents.Len()) {
		return true
	} else if (j > presents.Len()) {
		return false
	}
	if (presents.PresentStorage[i].Value == presents.PresentStorage[j].Value) {
		return presents.PresentStorage[i].Size < presents.PresentStorage[j].Size
	}
	return presents.PresentStorage[i].Value > presents.PresentStorage[j].Value
}

func (presents *PresentHeap) Swap(i, j int) {
	if (i >presents.Len()) || (j > presents.Len()) {
		return 
	}
	presents.PresentStorage[i].Value, presents.PresentStorage[j].Value = presents.PresentStorage[j].Value, presents.PresentStorage[i].Value
	presents.PresentStorage[i].Size, presents.PresentStorage[j].Size = presents.PresentStorage[j].Size, presents.PresentStorage[i].Size
}

func (presents *PresentHeap) Push(x any) {
	presents.PresentStorage = append(presents.PresentStorage, x.(Present))
	if (presents.IsSorted() == false) {
		presents.heapSort()
	}
}

func (presents *PresentHeap) Pop() any {
	old := presents.PresentStorage
	n := len(old)
	item := old[n-1]
	presents.PresentStorage = old[0 : n-1]
	return item
}

func (presents *PresentHeap) update(item *Present, value int, size int) {
	item.Value = value
	item.Size = size
	heap.Fix(presents, item.Value)
}

func printOutPresentHeap(presents *PresentHeap) {
	for pos, item := range presents.PresentStorage {
		fmt.Printf("%d| v: %d s: %d\n", pos, item.Value, item.Size)
	}
}

func getNCoolestPresents(storage []Present, n int) ([]Present, error) {
	if (n < 0) {
		return []Present{},  errors.New("n is negative!")
	} else if (n > len(storage)) {
		return []Present{}, errors.New("n is larger than the size of the slice")
	}

	result := new(PresentHeap)

	for _, item := range storage {
		result.Push(item)
	}
	for n < result.Len() {
		result.Pop()
	}

	return result.PresentStorage, nil
}

// heapSort

func (presents *PresentHeap) shiftDown(lo, hi, first int) {
	root := lo
	for {
		child := 2 * root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && presents.Less(first + child, first + child + 1) {
			child++
		}
		if !presents.Less(first + root, first + child) {
			return
		}
		presents.Swap(first + root, first + child)
		root = child
	}
}

func (presents *PresentHeap) heapSort() {
	size := presents.Len()

	for i := (size - 1) / 2; i >= 0; i-- {
		presents.shiftDown(i, size, 0)
	}

	for i := size - 1; i >= 0; i-- {
		presents.Swap(0, 0+i)
		presents.shiftDown(0, i, 0)
	}


}

func (presents *PresentHeap) IsSorted() bool {
	n := presents.Len()
	for i := n - 1; i > 0; i-- {
		if presents.Less(i, i-1) {
			return false
		}
	}
	return true
}


// ---

func main() {

	exampleSlice := []Present{{5, 1}, {4, 5}, {3, 1}, {5, 2}}
	myHeap := PresentHeap{exampleSlice}
	heap.Init(&myHeap)

	printOutPresentHeap(&myHeap)

	tmp, err := getNCoolestPresents(exampleSlice, 2)

	fmt.Println("getNCoolestPresents (n = 2): ", tmp, err)

	fmt.Println("-----------------------------------------")


	testSlice := []Present{{1, 7}, {5, 3}, {2, 8}, {2, 1}, {4, 2}, {5, 2}, {5, 7}}
	fmt.Println(testSlice)
	testHeap := PresentHeap{}
	testHeap.Push(Present{1,7})
	testHeap.Push(Present{5,3})
	testHeap.Push(Present{2,8})
	testHeap.Push(Present{2,1})
	testHeap.Push(Present{4,2})
	testHeap.Push(Present{5,2})
	testHeap.Push(Present{5,7})
	printOutPresentHeap(&testHeap)

	tmp2, err := getNCoolestPresents(testSlice, 3)
	fmt.Println("getNCoolestPresents (n = 3): ", tmp2, err)

	fmt.Println("-----------------------------------------")
	emptySlice := []Present{}
	tmp3, err := getNCoolestPresents(emptySlice, 3)
	fmt.Println("getNCoolestPresents (n = 3): ", tmp3, err)
	tmp3, err = getNCoolestPresents(emptySlice, -3)
	fmt.Println("getNCoolestPresents (n = -3): ", tmp3, err)
}