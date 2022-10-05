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

	t.Run("TestSecondImplementation", func(t *testing.T) {
		slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		num := 1
		got, _ := two(num, slice)
		expected := num

		if got != expected {
			t.Errorf("GOT: %q \nEXPECTED: %q ", got, expected)
		}
	})

	t.Run("TestThirdImplementation", func(t *testing.T) {
		slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		num := 6
		got, _ := three(num, slice)
		expected := num

		if got != expected {
			t.Errorf("GOT: %q \nEXPECTED: %q ", got, expected)
		}
	})
}
