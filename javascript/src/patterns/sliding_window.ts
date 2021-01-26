// maxSubarraySum gets an array of number and another number that represents the consecutive sum that the function returns O(n)
function maxSubarraySum(arr: number[], num: number): number | null {
    // return null if the array len is less than num
    if (arr.length < num) {
        return null
    }

    // create a variable for max number and initiate with the minimum number
    let maxSum = 0

    // create a variable for temp max number
    let tempSum = 0

    // loop until the num
    for (let i = 0; i < num; i++) {
        // add the values maxSum var
        maxSum += arr[i]
    }

    // assign the value of maxSum to the tempSum
    tempSum = maxSum

    // loop thru the array starting from the tempMax
    for (let i = num; i < arr.length; i++) {
        // reassign the tempSum with tempSum minus the value on array at minus num plus the value at i
        tempSum = tempSum - arr[i - num] + arr[i]

        // if the tempMax is bigger than max reassign max with tempMax
        maxSum = Math.max(maxSum, tempSum)
    }

    // return the max
    return maxSum
}

export default maxSubarraySum
