package challenges

import "testing"

func TestSameFrequency(t *testing.T) {
	t.Run("should return false if the len of the num is diff", func(t *testing.T) {
		num1 := 22
		num2 := 222
		got := SameFrequency(num1, num2)
		want := false
		if got != want {
			t.Errorf("expected %v got %v", got, want)
		}
	})

	t.Run("should return false when some has not same frequency of digs", func(t *testing.T) {
		num1 := 34
		num2 := 14
		got := SameFrequency(num1, num2)
		want := false
		if got != want {
			t.Errorf("expected %v got %v", got, want)
		}
	})

	t.Run("should return true when short number has same frequency of digs", func(t *testing.T) {
		num1 := 182
		num2 := 281
		got := SameFrequency(num1, num2)
		want := true
		if got != want {
			t.Errorf("expected %v got %v", got, want)
		}
	})

	t.Run("should return true when large number has same frequency of digs", func(t *testing.T) {
		num1 := 3589578
		num2 := 5879385
		got := SameFrequency(num1, num2)
		want := true
		if got != want {
			t.Errorf("expected %v got %v", got, want)
		}
	})
}
