package main

import (
	"reflect"
	"testing"
)

func isEqual(t *testing.T, got, expected int) {
	if got != expected {
		t.Errorf("Expected %q but got %q", expected, got)
	}
}

func isArrayEqual(t *testing.T, got, expected []int) {
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %q but got %q", expected, got)
	}
}

func TestSum(t *testing.T) {
	t.Run("Sum: 1...5 = 15", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		expected := 15
		isEqual(t, Sum(numbers), expected)
	})

	t.Run("Sum: 1...10 = 55", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		expected := 55
		isEqual(t, Sum(numbers), expected)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Sum: Two arrays", func(t *testing.T) {
		numbers := SumAll([]int{1, 2}, []int{0, 9})
		expected := []int{3, 9}
		isArrayEqual(t, numbers, expected)
	})
	t.Run("Sum: One arrays", func(t *testing.T) {
		numbers := SumAll([]int{1, 1, 1})
		expected := []int{3}
		isArrayEqual(t, numbers, expected)
	})
	t.Run("Sum: Zero arrays", func(t *testing.T) {
		numbers := SumAll([]int{})
		expected := []int{0}
		isArrayEqual(t, numbers, expected)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("SumTail: Two arrays", func(t *testing.T) {
		numbers := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}
		isArrayEqual(t, numbers, expected)
	})

	t.Run("SumTail: Three Different arrays", func(t *testing.T) {
		numbers := SumAllTails([]int{1}, []int{0, 9}, []int{0, 9, 8})
		expected := []int{1, 9, 8}
		isArrayEqual(t, numbers, expected)
	})

	t.Run("SumTail: safely sum empty slices", func(t *testing.T) {
		numbers := SumAllTails([]int{}, []int{0, 9})
		expected := []int{0, 9}
		isArrayEqual(t, numbers, expected)
	})
}
