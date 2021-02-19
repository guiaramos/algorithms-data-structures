package challenges

import "testing"

func TestAveragePair(t *testing.T) {
	t.Run("should return false when the array is empty", func(t *testing.T) {
		sliceInt := []int{}
		targetAvg := 4.0
		got := AveragePair(sliceInt, targetAvg)
		want := false

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("should return true when the array is short", func(t *testing.T) {
		sliceInt := []int{1, 2, 3}
		targetAvg := 2.5
		got := AveragePair(sliceInt, targetAvg)
		want := true

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("should return true when the array is long", func(t *testing.T) {
		sliceInt := []int{1, 3, 3, 5, 6, 7, 10, 12, 19}
		targetAvg := 8.0
		got := AveragePair(sliceInt, targetAvg)
		want := true

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("should return false when the target is wrong", func(t *testing.T) {
		sliceInt := []int{-1, 0, 3, 4, 5, 6}
		targetAvg := 4.1
		got := AveragePair(sliceInt, targetAvg)
		want := false

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
