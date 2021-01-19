// countUniqueValues receives a sorted array of numbers and return the count of unique values O(n)
export function countUniqueValues(values: number[]): number {
    // if array is empty return 0
    if (values.length === 0) return 0
    // create two pointer for the start and the next number of the array
    let left = 0

    // loop thru the array until it right pointer arrives the end of array
    for (let right = 0; right < values.length; right++) {
        // if the left number is equal to the right, increase the right pointer
        if (values[left] !== values[right]) {
            // move the left pointer
            left++
            // assign the value of right pointer to the left pointer
            values[left] = values[right]
        }
    }

    // return the left pointer + 1
    return left + 1
}
