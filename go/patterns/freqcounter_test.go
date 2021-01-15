package patterns

import "testing"

func TestSameArraySquared(t *testing.T) {
	t.Run("should return false if length is not same", func(t *testing.T) {
		inArr := []int{1, 2, 3}
		sqArr := []int{1, 4, 9, 6}

		got := SameArraySquared(inArr, sqArr)
		want := false

		assertBool(t, got, want)
	})

	t.Run("should return false if squared is not found", func(t *testing.T) {
		inArr := []int{1, 2, 3}
		sqArr := []int{1, 4, 3}

		got := SameArraySquared(inArr, sqArr)
		want := false

		assertBool(t, got, want)
	})

	t.Run("should return false if freq is not same", func(t *testing.T) {
		inArr := []int{1, 2, 1}
		sqArr := []int{1, 4, 3}

		got := SameArraySquared(inArr, sqArr)
		want := false

		assertBool(t, got, want)
	})

	t.Run("should return true for correct freq ordered", func(t *testing.T) {
		inArr := []int{1, 2, 3}
		sqArr := []int{1, 4, 9}

		got := SameArraySquared(inArr, sqArr)
		want := true

		assertBool(t, got, want)
	})

	t.Run("should return true for correct freq unordered", func(t *testing.T) {
		inArr := []int{3, 2, 1}
		sqArr := []int{9, 4, 1}

		got := SameArraySquared(inArr, sqArr)
		want := true

		assertBool(t, got, want)
	})
}

func assertBool(t *testing.T, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}
