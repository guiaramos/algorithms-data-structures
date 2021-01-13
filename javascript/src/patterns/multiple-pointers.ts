// sumZero accepts a sorted array of integers and find the first pair where the sum is. O(n)
function sumZero (sortedArr: Array<number>):Array<number> | undefined {
  if (sortedArr.length === 0) return undefined
  let left = 0
  let right = sortedArr.length - 1

  while (left < right) {
    const numLeft = sortedArr[left]
    const numRight = sortedArr[right]
    const sum = numLeft + numRight
    if (sum === 0) return [numLeft, numRight]
    else if (sum > 0) right--
    else left++
  }

  return undefined
}

export { sumZero }
