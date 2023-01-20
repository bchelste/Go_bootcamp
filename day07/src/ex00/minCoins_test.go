package main

import (
	"reflect"
	"testing"
	"time"
)

type TestCase struct {
	nbr int
	storage []int
	result []int
}

func testUtils(t *testing.T, data TestCase) {
	timeout := time.After(3 * time.Second)
    done := make(chan []int)
	go func() {
		res := minCoins(data.nbr, data.storage)
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

	t.Parallel()

	t.Run("block#1 of testing minCoins", func(t *testing.T) {
		t.Run("test case from subject #1", func(t *testing.T) {testUtils(t, TestCase{
			13, 
			[]int{1,5,10}, 
			[]int{10,1,1,1}})})
		t.Run("test case from subject #2", func(t *testing.T) {testUtils(t, TestCase{
			3642, 
			[]int{1,5,10,50,100,500,1000}, 
			[]int{1000,1000,1000,500,100,10,10,10,10,1,1}})})
		t.Run("case with sum = 0", func(t *testing.T) {testUtils(t, TestCase{
			0, 
			[]int{1,5,10,50,100,500,1000}, 
			[]int{}})})
		t.Run("case with empty slice", func(t *testing.T) {testUtils(t, TestCase{
			10, 
			[]int{}, 
			[]int{}})})
		t.Run("case with negative sum = -5", func(t *testing.T) {testUtils(t, TestCase{
			-5, 
			[]int{1,5,10,50,100,500,1000}, 
			[]int{}})})
		t.Run("case with some negative denominations", func(t *testing.T) {testUtils(t, TestCase{
			111, 
			[]int{-1,5,10,50,100,500,1000}, 
			[]int{}})})
		t.Run("sum < denominations case", func(t *testing.T) {testUtils(t, TestCase{
			2, 
			[]int{10,50,100,500,1000}, 
			[]int{}})})
		t.Run("unsorted slice", func(t *testing.T) {testUtils(t, TestCase{
			110, 
			[]int{100,50,10,1000,500}, 
			[]int{100, 10}})})
		t.Run("duplicates in denominations", func(t *testing.T) {testUtils(t, TestCase{
			80, 
			[]int{10,10,50,50,500}, 
			[]int{50, 10, 10, 10}})})
		t.Run("unsorted duplicates in denominations", func(t *testing.T) {testUtils(t, TestCase{
			80, 
			[]int{10,50,50,10,500}, 
			[]int{50, 10, 10, 10}})})
	})

	t.Run("block#2 of testing minCoins2", func(t *testing.T) {
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