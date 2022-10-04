package main

//Bubble sort
func bubbleSort(sliceOfInts []int) []int {
	for i := 0; i < len(sliceOfInts); i++ {
		for j := 0; j < len(sliceOfInts)-1; j++ {
			if sliceOfInts[j] > sliceOfInts[j+1] {
				sliceOfInts[j], sliceOfInts[j+1] = sliceOfInts[j+1], sliceOfInts[j]
			}
		}
	}
	return sliceOfInts
}
