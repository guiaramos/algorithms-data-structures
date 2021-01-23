package challenges

import "testing"

func TestUniqueValues(t *testing.T) {
	t.Run("should return 0 when slice is empty", func(t *testing.T) {
		a := []int{}
		got := CountUniqueValues(a)
		want := 0
		if got != want {
			t.Fatalf("got %d, want %d", got, want)
		}
	})

	t.Run("should return 2 when slice has length 2 and 2 unique numbers", func(t *testing.T) {
		a := []int{1, 2}
		got := CountUniqueValues(a)
		want := 2
		if got != want {
			t.Fatalf("got %d, want %d", got, want)
		}
	})

	t.Run("should return 2 when slice has large length and 2 unique numbers", func(t *testing.T) {
		a := []int{1, 1, 1, 1, 1, 2}
		got := CountUniqueValues(a)
		want := 2
		if got != want {
			t.Fatalf("got %d, want %d", got, want)
		}
	})

	t.Run("should return 7 when slice has large length and 7 unique numbers", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 4, 4, 7, 7, 12, 12, 13}
		got := CountUniqueValues(a)
		want := 7
		if got != want {
			t.Fatalf("got %d, want %d", got, want)
		}
	})

	t.Run("should return 4 when slice has large length and 4 unique numbers with negative numbers", func(t *testing.T) {
		a := []int{-2, -1, -1, 0, 1}
		got := CountUniqueValues(a)
		want := 4
		if got != want {
			t.Fatalf("got %d, want %d", got, want)
		}
	})
}
