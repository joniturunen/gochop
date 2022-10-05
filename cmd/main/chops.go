package main

import "fmt"

// Simple first implementation
// Runtime complexity: O(n)
func one(num int, slice []int) (int, int) {
	var numberOfIterations int = 0
	for i, v := range slice {
		numberOfIterations++
		if v == num {
			return i, numberOfIterations
		}
	}
	return -1, numberOfIterations
}

// Start binary search from the middle of the slice
// Runtime complexity: ???
func two(num int, slice []int) (int, int) {
	numberOfIterations := 0
	n := (len(slice) + 1) / 2
	for i := 0; i < len(slice); i++ {
		numberOfIterations++
		if slice[n] == num {
			return n, numberOfIterations
		} else if slice[n] > num {
			n--
		} else {
			n++
		}
	}
	return -1, numberOfIterations
}

// Start binary search from the middle of the slice
// Runtime complexity: ???
func three(num int, slice []int) (int, int) {
	numberOfIterations := 0
	mid := (len(slice)) / 2
	max := len(slice)
	min := 0
	for i := 0; i < max; i++ {
		showIterInfo(mid, i)
		numberOfIterations++
		if slice[mid] == num {
			return mid, numberOfIterations
		} else if slice[mid] > num {
			max = mid
			mid = (max + min) / 2
		} else {
			min = mid
			mid = (max + min) / 2
		}
	}
	return -1, numberOfIterations

}

func showIterInfo(num int, index int) {
	fmt.Println("#", index+1, " =", num)
}
