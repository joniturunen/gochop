package main

import "testing"

func Test_example(t *testing.T) {
	t.Run("TestFirstImplementation", func(t *testing.T) {
		slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		num := 7
		got, gotIterations := one(num, slice)
		expected, expectedIterations := num, 11

		if got != expected {
			t.Errorf("GOT: %q and %q\nEXPECTED: %q and %q", got, gotIterations, expected, expectedIterations)
		}
	})
}
