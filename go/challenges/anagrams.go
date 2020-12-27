package challenges

// AnagramsErr is a type for anagrams errors
type AnagramsErr string

func (e AnagramsErr) Error() string {
	return string(e)
}

const (
	// ErrDifferentLengths defines the error message for strings with different lengths
	ErrDifferentLengths        = AnagramsErr("the strings have different lengths")
	// ErrStringHasDiffLetter defines the error message for strings with different letters
	ErrStringHasDiffLetter     = AnagramsErr("the strings have different letters")
	// ErrStringHasDiffLetterFreq defines the error message for strings with different letters
	ErrStringHasDiffLetterFreq = AnagramsErr("the strings have different letter frequency")
)

// CheckAnagrams validates if two strings a`re anagrams O(n)
func CheckAnagrams(s1, s2 string) error {

	// Check if the strings have the same length
	if len(s1) != len(s2) {
		return ErrDifferentLengths
	}

	// create a map for the letter frequency in the first string
	freqS1 := make(map[rune]int)

	// interate thru the struct and record the frequency
	for _, l := range s1 {
		if c, ok := freqS1[l]; ok {
			freqS1[l] = c + 1
		} else {
			freqS1[l] = 1
		}
	}

	// loop thru the letters of the second string
	for _, l := range s2 {
		// check if the letter exits in the frequency struct
		if c, ok := freqS1[l]; ok {
			// if frequency is zero return false
			if c == 0 {
				return ErrStringHasDiffLetterFreq
			}
			// otherwise, decrease the frequency
			freqS1[l] = c - 1
		} else {
			// otherwise return error
			return ErrStringHasDiffLetter
		}
	}

	return nil
}

