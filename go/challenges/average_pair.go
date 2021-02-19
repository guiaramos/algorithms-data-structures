package challenges

// AveragePair takes a sorted array of int and a target avg,
// determine if there is a pair of values in the array where the average of the pair
// equals the target average
func AveragePair(sliceInt []int, targetAvg float64) bool {
	// check if the len of the slice is 0 and return false
	if len(sliceInt) == 0 {
		return false
	}

	// create a "right" pointer to the end of the slice
	r := len(sliceInt) - 1

	// loop thru the slice with the "left" pointer
	for l, v := range sliceInt {
		// calculate avg of the 2 pointers
		avg := float64(v+sliceInt[r]) / 2
		// if the avg is same as the target avg returns true
		if avg == targetAvg {
			return true
			// if the left pointer is bigger than the right pointer, return false
		} else if l > r {
			return false
			// if the avg is bigger than the target avg, moves the right pointer
		} else if avg > targetAvg {
			r--
		}
	}

	// if the checks pass, return false
	return false
}
