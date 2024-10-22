package garray

import (
	"fmt"
	"math"
)

func ExampleMax() {
	// Finding the maximum element of a slice of integers
	max := Max([]int{10, -20, 5, 15, 30, -25})
	fmt.Println(max)
	// Output: 30
}

func ExampleMax_panicsOnEmptySlice() {
	// Trying to find the maximum element of an empty slice should panic
	defer func() {
		if r := recover(); r == nil {
			fmt.Println(`Max didn't panic`)
		} else {
			fmt.Println(r)
		}
	}()
	Max([]int{})
	// Output: Max: empty slice
}

func ExampleMin() {
	// Finding the minimum element of a slice of floats
	min := Min([]float64{3.14, 2.71, 0.0, -1.0, 1.5, -2.0})
	fmt.Println(min)
	// Output: -2
}

func ExampleMaxBy() {
	// Finding the maximum length string from a slice of strings
	maxLen := MaxBy([]string{`apple`, `banana`, `cherry`, `date`}, func(s string) int {
		return len(s)
	})
	fmt.Println(maxLen)
	// Output: 6
}

func ExampleMinBy() {
	// Finding the minimum even number from a slice of integers
	minEven := MinBy([]int{3, 6, -1, 0, 8, 7, 9}, func(n int) int {
		if n%2 == 0 {
			return n
		} else {
			return math.MaxInt32
		}
	})
	fmt.Println(minEven)
	// Output: 0
}
func ExampleMaxIndex() {
	// Finding the index of the maximum element in a slice of integers
	maxIndex := MaxIndex([]int{10, -20, 5, 15, 30, -25}, func(i int) int {
		return []int{10, -20, 5, 15, 30, -25}[i]
	})
	fmt.Println(maxIndex)
	// Output: 4
}

func ExampleMaxIndex_emtpy() {
	// Trying to find the index of the maximum element in an empty slice should return -1
	maxIndex := MaxIndex([]int{}, func(i int) int {
		return []int{}[i]
	})
	fmt.Println(maxIndex)
	// Output: -1
}

func ExampleMinIndex() {
	// Finding the index of the maximum element in a slice of integers
	maxIndex := MinIndex([]int{10, -20, 5, 15, 30, -25}, func(i int) int {
		return []int{10, -20, 5, 15, 30, -25}[i]
	})
	fmt.Println(maxIndex)
	// Output: 5
}
