package challenges

/*
CountUniqueValues takes a sorted slice of numbers
 and returns the count of unique numbers
*/
func CountUniqueValues(nums []int) int {
	// check if the len of the slice is 0 and return 0
	if len(nums) == 0 {
		return 0
	}

	// make a pointer for the start of the slice
	start := 0

	// interate with the slice with a pointer next to the start pointer while it's small than the length of the array
	for next := 1; next < len(nums); next++ {
		// check if the value on the pointer are diff
		if nums[start] != nums[next] {
			// move the start pointer one position
			start++
			// assign the value from the next pointer to start pointer
			nums[start] = nums[next]
		}
	}

	// return the start pointer + 1
	return start + 1
}
