/**
 * sameArraySquared gets two array and compares if the second array represents the square number of the first array
 * Time complexity: O(n)
 * */
function sameArraySquared(firstArray: number[], secondArray: number[]): boolean {
    // check if the array have same length
    if (firstArray.length !== secondArray.length) {
        return false
    }

    // create object for save frequency of each array
    const freqCounterOne: Record<string, number> = {}
    const freqCounterTwo: Record<string, number> = {}

    // iterate on first array to get the frequency
    for (const key of firstArray) {
        freqCounterOne[key] = (freqCounterOne[key] || 0) + 1
    }

    // iterate on second array to get the frequency
    for (const key of secondArray) {
        freqCounterTwo[key] = (freqCounterTwo[key] || 0) + 1
    }

    // iterate on first array and compares the frequency, if mismatch return false
    for (const key in freqCounterOne) {
        if (!(Number(key) ** 2 in freqCounterTwo)) {
            return false
        }
        if (freqCounterTwo[Number(key) ** 2] !== freqCounterOne[key]) {
            return false
        }
    }

    return true
}

export { sameArraySquared }
