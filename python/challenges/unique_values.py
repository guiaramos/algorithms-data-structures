# countUniqueValues takes a sorted array of numbers and return the count of unique numbers
def countUniqueValues(numbers):
    # if len of numbers is 0 return 0
    if(len(numbers) == 0):
        return 0

    # create two pointers for the start of the list and next of it
    prev = 0
    next = 1

    # interate with the numbers
    while next < len(numbers):

        # if the values of two pointer are different
        if(numbers[prev] != numbers[next]):
            # increment prev pointer
            prev += 1
            # assign the value on position of next pointer to the position of prev pointer
            numbers[prev] = numbers[next]

        # increment next pointer
        next += 1

    # return the prev pointer plus 1
    return prev + 1
