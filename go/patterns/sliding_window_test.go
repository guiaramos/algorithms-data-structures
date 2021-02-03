package patterns

import "testing"

func TestMaxSubarraySum(t *testing.T) {
	t.Run("test should return 0 when the n is less than the slice len", func(t *testing.T) {
		s := []int{}
		n := 4
		got := MaxSubarraySum(s, n)
		if got != 0 {
			t.Errorf("expected zero but received value %v", got)
		}
	})

	t.Run("test should return the max int when slice has no repetitive int and n is 1", func(t *testing.T) {
		s := []int{4, 2, 1, 6}
		n := 1
		got := MaxSubarraySum(s, n)
		want := 6
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("test should return the max sum int when slice has repetitive int and n is 4", func(t *testing.T) {
		s := []int{4, 2, 1, 6, 2}
		n := 4
		got := MaxSubarraySum(s, n)
		want := 13
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("test should return the max sum int when slice has multiple repetitive int and n is 4", func(t *testing.T) {
		s := []int{1, 2, 5, 2, 8, 1, 5}
		n := 4
		got := MaxSubarraySum(s, n)
		want := 17
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("test should return the max sum int when slice has multiple repetitive int and n is 2", func(t *testing.T) {
		s := []int{1, 2, 5, 2, 8, 1, 5}
		n := 2
		got := MaxSubarraySum(s, n)
		want := 10
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
