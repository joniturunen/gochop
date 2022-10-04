package main

import (
	"go.uber.org/zap"
)

func main() {
	sugar := initializeLogger().Sugar()
	sliceOfIntegers := []int{82, 66, 14, 1, 3, 42, 8, 321, 54, 12, 56, 6, 9}
	sortedSliceOfInts := bubbleSort(sliceOfIntegers)

	// Number to search for
	num := 66

	// First implementation
	one, iterations := one(num, sortedSliceOfInts)
	sugar.Info("One: ", one, ". Iterations: ", iterations)
	// Second implementation
	two, iterations := two(num, sortedSliceOfInts)
	sugar.Info("Two: ", two, ". Iterations: ", iterations)
}

func initializeLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	logger.Debug("Logger initialized")

	return logger
}
