# max_subarray_sum takes an array of int and a nums that represents the length of subarray
# it returns the max sum of subarray with the length defined by nums O(n)
def max_subarray_sum(arr, nums):
    # return None if the list length is less than the nums
    if len(arr) < nums:
        return None

    # create a max variable to store the max sum found
    max_sum = 0
    # create a temp variable to store the max sum while checking the subarray
    temp_sum = 0

    # var helper for the loop
    i = 0
    # loop thru the arr until the nums
    while i < nums:
        # store the first subarray max sum on the max variable
        max_sum += arr[i]
        # increase the helper var
        i += 1

    # assign the value of max var to the temp var
    temp_sum = max_sum

    # var helper for the loop
    i = nums
    # loop thru the array starting from the nums until the length
    while i < len(arr):
        # reassign the temp variable by removing the first value of subarray from the temp value
        # and add the the current value of the array on position num
        # example: [4, 2, 1, 6, 2] nums: 2
        #          |- temp| +
        #            |- temp| +
        #               |- temp|
        temp_sum = temp_sum - arr[i-nums] + arr[i]
        # if the temp is bigger than max, reassign the max with the temp value
        max_sum = max(max_sum, temp_sum)
        # increase the helper var
        i += 1

    # return the max
    return max_sum
