package challenges

import (
	"testing"
)

func TestCheckAnagrams(t *testing.T) {
	t.Run("expect to get an error if strings have different length", func(t *testing.T) {
		firstString := "awesome"
		secondString := "awesom"

		err := CheckAnagrams(firstString, secondString)

		assertError(t, err, ErrDifferentLengths)
	})

	t.Run("expect to get an error if the strings have different letter", func(t *testing.T) {
		firstString := "car"
		secondString := "rat"

		err := CheckAnagrams(firstString, secondString)

		assertError(t, err, ErrStringHasDiffLetter)
	})

	t.Run("expect to get an error if the strings have different letter frequency", func(t *testing.T) {
		firstString := "aaz"
		secondString := "zza"

		err := CheckAnagrams(firstString, secondString)

		assertError(t, err, ErrStringHasDiffLetterFreq)
	})

	t.Run("should not return when input is correct anagram", func(t *testing.T) {
		firstString := "anagram"
		secondString := "nagaram"

		err := CheckAnagrams(firstString, secondString)

		if err != nil {
			t.Errorf("expected to get no error, but got %q", err)
		}
	})
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got == nil {
		if want == nil {
			return
		}
		t.Fatal("expected to get an error")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
