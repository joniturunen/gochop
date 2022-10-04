package main

import (
	"go.uber.org/zap"
)

func main() {
	sugar := initializeLogger().Sugar()
	sliceOfIntegers := generateSliceOfIntegers(1000000)
	numberOfInts := len(sliceOfIntegers)

	// Number to search for
	num := 415143
	sugar.Info("Number of indexes ", numberOfInts, ". Number to search for ", num, ".")

	// First implementation
	one, iterations := one(num, sliceOfIntegers)
	sugar.Info("One: ", one, ". Iterations: ", iterations)
	// Second implementation
	two, iterations := two(num, sliceOfIntegers)
	sugar.Info("Two: ", two, ". Iterations: ", iterations)
	// Third implementation
	three, iterations := three(num, sliceOfIntegers)
	sugar.Info("Three: ", three, ". Iterations: ", iterations)
}

func initializeLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	logger.Debug("Logger initialized")

	return logger
}

func generateSliceOfIntegers(size int) []int {
	var sliceOfIntegers []int
	for i := 0; i < size; i++ {
		sliceOfIntegers = append(sliceOfIntegers, i)
	}
	return sliceOfIntegers
}
