# sumZero receives an array of numbers and return the first pair where sums is zero
def sumZero(numbers):
    # if the len is zero should return None
    if (len(numbers) == 0):
        return None

    # create two pointer for the start and end of the array
    left = 0
    right = len(numbers) - 1

    # interate until left is less than right pointer
    while (left < right):
        sum = numbers[left] + numbers[right]
        # if the sum of the pointers value is equal zero return the pointers
        if (sum == 0):
            return [numbers[left], numbers[right]]
        # if the sum is greater than zero decrease right pointer
        elif (sum > 0):
            right -= 1
        # otherwise increase the left pointer
        else:
            left += 1

    # if the pair is not found return None
    return None
