# average_pair gets a sorted list of integers and a target average,
# deterines if there is a pair of values in the list where the average of the pair
# equals the target average O(n)
def average_pair(int_list, target_avg):
    # if the len of int_list is zero returns false
    if len(int_list) == 0:
        return False

    # creates a pointer "right" to the end of the list
    right = len(int_list) - 1

    # loop thru the list
    for left in int_list:
        # calculates the avg (left value + right value / 2) and store in a var
        avg = (left + int_list[right]) / 2

        # if the avg matches with the target returns True
        if avg == target_avg:
            return True
        # else if the avg is bigger than the target, decrease the right pointer
        elif avg > target_avg:
            right -= 1

    # if the avg is not found returns false
    return False
