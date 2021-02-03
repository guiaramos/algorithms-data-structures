package patterns

// MaxSubarraySum takes an slice and an "n" int to find the max subarray sum with the length equal to the "n" O(n)
func MaxSubarraySum(slice []int, n int) int {
	// checks if the len of slice is less from the n
	if len(slice) < n {
		// then resturns zero
		return 0
	}

	// create a variable for storing the max with zero
	maxSum := 0

	// loop thru the slice until the n
	for i, v := range slice {
		// if the i is bigger than the n -1 break
		if i > n-1 {
			break
		}
		maxSum += v
	}
	// create another variable for storing the max sum with the value of max sum
	tempSum := maxSum

	// loop thru the array from the n value
	for i, v := range slice {
		// if the i is less than the n skip the turn
		if i < n {
			continue
		}
		// reassign the tempSum with the value of tempSum
		// minos the tempSum and the value slice poisition where current index minos the n value
		// plus the current value
		// in this way, it removes the previous value of the subarray and adds the next one
		tempSum = tempSum - slice[i-n] + v

		// if the temp sum is bigger than the max sum
		if tempSum > maxSum {
			// assign the max sum with the temp sum
			maxSum = tempSum
		}
	}

	// return the max
	return maxSum
}
