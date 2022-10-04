package main

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
// Runtime complexity: O(log n)
func two(num int, slice []int) (int, int) {
	var numberOfIterations int = 0
	var n = (len(slice) + 1) / 2
	for i := 0; i < len(slice); i++ {
		numberOfIterations++
		if slice[n] == num {
			return n, numberOfIterations
		} else if slice[n] > num {
			n = n / 2
		} else {
			n = n + n/2
		}
	}
	return -1, numberOfIterations

}
