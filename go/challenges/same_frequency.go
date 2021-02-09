package challenges

import "strconv"

// SameFrequency takes two int and compares if they have the same frequency of digits
// and return true, otherwise false O(n)
func SameFrequency(num1, num2 int) bool {
	// converts the numbers in strings
	strNum1 := strconv.Itoa(num1)
	strNum2 := strconv.Itoa(num2)

	// check if the numbers has diff len and return false if so
	if len(strNum1) != len(strNum2) {
		return false
	}

	// create a map variable to store the frequency
	freqNum1 := make(map[rune]int)

	// loop thru the num1
	for _, v := range strNum1 {
		// save the frequency in the variable
		if _, ok := freqNum1[v]; ok {
			freqNum1[v]++
		} else {
			freqNum1[v] = 1
		}
	}

	// loop thru the num2
	for _, v := range strNum2 {
		// check if the current value is not a key in the map then return false
		if _, ok := freqNum1[v]; !ok {
			return false
			// else if the value in the map is zero or less then return false
		} else if freqNum1[v] <= 0 {
			return false
			// otherwise decrease the value in the map
		} else {
			freqNum1[v]--
		}
	}

	// returns true if all check passes
	return true
}
