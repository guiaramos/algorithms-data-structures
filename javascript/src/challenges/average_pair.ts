// averagePair gets a sorted array of int and determine if there is a pair of values
// where the average of the pair equals the target average passed as second arg. O(n)
function averagePair(intArr: number[], targetAvg: number): boolean {
    // if the array has len zero returns false
    if (intArr.length === 0) return false

    // create a pointer to the end of the array
    let right = intArr.length - 1

    // loop thru the array
    for (let left = 0; left < intArr.length; left++) {
        // get the avg of the 2 pointers (left + right) / 2
        const avg = (intArr[left] + intArr[right]) / 2
        // if the avg is equal to the target return true
        if (avg === targetAvg) return true
        // if the right cross with left return false
        else if (left >= right) return false
        // if the avg is bigger than the decrease the right pointer
        else if (avg > targetAvg) right--
    }

    // if no avg found returns false
    return false
}

export default averagePair
