package patterns

import (
	"fmt"
	"math"
)

// SameArraySquared returns true if all int in slice is squared on other O(n)
func SameArraySquared(inArr []int, sqArr []int) bool {
	// check if the arrays have the same len
	if len(inArr) != len(sqArr) {
		return false
	}

	// create a map for the frequency of input array
	freqInArr := make(map[int]int)

	// create a map for the frequency of squared array
	freqSqArr := make(map[int]int)

	// interate thru input array and add the key
	for _, key := range inArr {
		if freq, found := freqInArr[key]; found {
			freqInArr[key] = freq + 1
		} else {
			freqInArr[key] = 1
		}
	}

	// interate tru squared array and add the key
	for _, key := range sqArr {
		if freq, found := freqSqArr[key]; found {
			freqSqArr[key] = freq + 1
		} else {
			freqSqArr[key] = 1
		}
	}

	// interate thru the frequency of input array
	for k, f := range freqInArr {
		sq := int(math.Pow(float64(k), 2))
		freq, found := freqSqArr[sq]

		// check if the squared key on frequency of input array is in the frequency of the squared array
		if !found {
			return false
		}

		// check if the frequency of input array matchs with the frequency of the squared array
		if f != freq {
			return false
		}
	}

	// return true if the checks pass
	return true
}

func main() {
	inArr := []int{3, 2, 1}
	sqArr := []int{9, 4, 1}

	got := SameArraySquared(inArr, sqArr)
	fmt.Printf("got %t", got)
}
