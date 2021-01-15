package patterns

import (
	"reflect"
	"testing"
)

func TestSumZero(t *testing.T) {
	t.Run("it should return nil if the array is empty", func(t *testing.T) {
		test := []int{}

		got := SumZero(test)

		assertNil(t, got)
	})

	t.Run("it should return the first pair of the array", func(t *testing.T) {
		test := []int{-8, -5, -3, 0, 1, 4, 8}

		got := SumZero(test)
		want := []int{-8, 8}

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v, want %v", got, want)
		}
	})

	t.Run("it should return nil when there is no pair where sum is zero", func(t *testing.T) {
		test := []int{1, 2, 3, 4, 5}

		got := SumZero(test)

		assertNil(t, got)
	})

	t.Run("it should return pair in the middle of the array", func(t *testing.T) {})
	test := []int{-8, -5, -3, 0, 3, 4, 9}

	got := SumZero(test)
	want := []int{-3, 3}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func assertNil(t *testing.T, got []int) {
	t.Helper()

	if got != nil {
		t.Fatalf("expected nil but received %v", got)
	}
}
