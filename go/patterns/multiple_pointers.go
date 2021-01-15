package patterns

// SumZero should return the first pair where sum is equal to zero, otherwise nil
func SumZero(numbers []int) []int {
	// if the len is equal to zero return nil
	if len(numbers) == 0 {
		return nil
	}

	// create two pointers for the start and end of the array
	left := 0
	right := len(numbers) - 1

	// interate with the numbers arrays
	for left < right {
		// get the sum of the two pointers values
		sum := numbers[left] + numbers[right]
		// if the sum is equal to zero return the pointers values
		if sum == 0 {
			return []int{numbers[left], numbers[right]}
			// if the sum is greater than zero decrease right pointer
		} else if sum > 0 {
			right--
			// otherwise increase the left pointer
		} else {
			left++
		}
	}

	// otherwise, return nil
	return nil
}
