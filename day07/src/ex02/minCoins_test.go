package main

import (
	"reflect"
	"testing"
	"time"
	"fmt"
)

type TestCase struct {
	nbr int
	storage []int
	result []int
}

func testUtils2(t *testing.T, data TestCase) {
	timeout := time.After(3 * time.Second)
    done := make(chan []int)
	go func() {
		res := minCoins2(data.nbr, data.storage)
		done <- res
	}()
	select {
    case <-timeout:
        t.Fatal("Test didn't finish in time")
    case current := <-done:
		if (!reflect.DeepEqual(data.result, current)) {
			t.Errorf("expected result: %d | current result: %d\n ", data.result, current)
		}
    }
}

func TestMinCoins(t *testing.T) {

	// t.Parallel()

	t.Run("block of testing minCoins2", func(t *testing.T) {
		t.Run("test case from subject #1", func(t *testing.T) {testUtils2(t, TestCase{
			13, 
			[]int{1,5,10}, 
			[]int{10,1,1,1}})})
		t.Run("test case from subject #2", func(t *testing.T) {testUtils2(t, TestCase{
			3642, 
			[]int{1,5,10,50,100,500,1000}, 
			[]int{1000,1000,1000,500,100,10,10,10,10,1,1}})})
		t.Run("case with sum = 0", func(t *testing.T) {testUtils2(t, TestCase{
			0, 
			[]int{1,5,10,50,100,500,1000}, 
			[]int{}})})
		t.Run("case with empty slice", func(t *testing.T) {testUtils2(t, TestCase{
			10, 
			[]int{}, 
			[]int{}})})
		t.Run("case with negative sum = -5", func(t *testing.T) {testUtils2(t, TestCase{
			-5, 
			[]int{1,5,10,50,100,500,1000}, 
			[]int{}})})
		t.Run("case with some negative denominations", func(t *testing.T) {testUtils2(t, TestCase{
			111, 
			[]int{-1,5,10,50,100,500,1000}, 
			[]int{}})})
		t.Run("sum < denominations case", func(t *testing.T) {testUtils2(t, TestCase{
			2, 
			[]int{10,50,100,500,1000}, 
			[]int{}})})
		t.Run("unsorted slice", func(t *testing.T) {testUtils2(t, TestCase{
			110, 
			[]int{100,50,10,1000,500}, 
			[]int{100, 10}})})
		t.Run("duplicates in denominations", func(t *testing.T) {testUtils2(t, TestCase{
			80, 
			[]int{10,10,50,50,500}, 
			[]int{50, 10, 10, 10}})})
		t.Run("unsorted duplicates in denominations", func(t *testing.T) {testUtils2(t, TestCase{
			80, 
			[]int{10,50,50,10,500}, 
			[]int{50, 10, 10, 10}})})
	})

}

func BenchmarkTest0(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fmt.Sprintf("hello")
    }
}

func BenchmarkTest1(b *testing.B) {
	for i := 0; i < b.N; i++ {
       minCoins2(13, []int{1,5,10})
    }
}

func BenchmarkTest2(b *testing.B) {
	for i := 0; i < b.N; i++ {
       minCoins2(3642, []int{1,5,10,50,100,500,1000})
    }
}

func BenchmarkTest3(b *testing.B) {
	for i := 0; i < b.N; i++ {
       minCoins2(0, []int{1,5,10,50,100,500,1000})
    }
}

func BenchmarkTest4(b *testing.B) {
	for i := 0; i < b.N; i++ {
       minCoins2(10, []int{})
    }
}

func BenchmarkTest5(b *testing.B) {
	for i := 0; i < b.N; i++ {
       minCoins2(-5, []int{1,5,10,50,100,500,1000})
    }
}

func BenchmarkTest6(b *testing.B) {
	for i := 0; i < b.N; i++ {
       minCoins2(111, []int{1,-5,10,50,100,500,1000})
    }
}
func BenchmarkTest7(b *testing.B) {
	for i := 0; i < b.N; i++ {
       minCoins2(2, []int{5,10,50,100,500,1000})
    }
}

func BenchmarkTest8(b *testing.B) {
	for i := 0; i < b.N; i++ {
       minCoins2(110, []int{100,50,10,1000,500})
    }
}

func BenchmarkTest9(b *testing.B) {
	for i := 0; i < b.N; i++ {
       minCoins2(80, []int{10,50,50,10,500})
    }
}

// go test -bench=. | grep Benchmark | sort -nk2 > top10.txt 