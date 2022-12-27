package arraysandslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("for one slice", func(t *testing.T) {
		sliceOfNumbers := [][]int{
			{1, 2, 3},
		}

		got := SumAll(sliceOfNumbers...)
		want := []int{6}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("for two slice", func(t *testing.T) {
		sliceOfNumbers := [][]int{
			{1, 2, 3},
			{4, 5, 6, 7, 8},
		}

		got := SumAll(sliceOfNumbers...)
		want := []int{6, 30}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("for one slice", func(t *testing.T) {
		sliceOfNumbers := [][]int{
			{1, 2, 3},
		}

		got := SumAllTails(sliceOfNumbers...)
		want := []int{5}

		checkSums(t, got, want)
	})

	t.Run("for two slice", func(t *testing.T) {
		sliceOfNumbers := [][]int{
			{1, 2, 3},
			{4, 5, 6, 7, 8},
		}

		got := SumAllTails(sliceOfNumbers...)
		want := []int{5, 26}

		checkSums(t, got, want)
	})

	t.Run("empty slice", func(t *testing.T) {
		sliceOfNumbers := [][]int{
			{},
			{4, 5, 6, 7, 8},
		}

		got := SumAllTails(sliceOfNumbers...)
		want := []int{0, 26}

		checkSums(t, got, want)
	})
}
