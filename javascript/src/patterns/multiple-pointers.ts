// sumZero accepts a sorted array of integers and find the first pair where the sum is. O(n)
function sumZero(sortedArr: Array<number>): Array<number> | undefined {
    // if the len is 0 return undefined
    if (sortedArr.length === 0) return undefined

    // add first pointer to the start of the array
    let left = 0

    // add the second pointer to the end of the array
    let right = sortedArr.length - 1

    // loop until the start pointer is smaller than the end of the array
    while (left < right) {
        const numLeft = sortedArr[left]
        const numRight = sortedArr[right]
        const sum = numLeft + numRight
        // check if the sum is 0 and return the array locations
        if (sum === 0) return [numLeft, numRight]
        // otherwise check if the sum is bigger than zero and decrease the right pointer
        else if (sum > 0) right--
        // otherwise increase left pointer
        else left++
    }

    // if none matches return undefined
    return undefined
}

export { sumZero }
