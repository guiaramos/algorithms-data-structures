# same_frequency takes two integers and compare the frequency
# should returns false if the frequency is different
#  and returns true when the frequency match
def same_frequency(num1, num2):
    # create list for both numbers
    num1_list = list(str(num1))
    num2_list = list(str(num2))
    # checks if the length
    if len(num1_list) != len(num2_list):
        return False
    # create a dict for storing the first num frequency of digs
    num1_freq = {}

    # loop thru the first num list and add the frequency to the dict
    for dig in num1_list:
        if dig in num1_freq:
            num1_freq[dig] += 1
        else:
            num1_freq[dig] = 1

    # loop thru the second num list
    for dig in num2_list:
        # check if the dig is on the frequency dict and returns false
        if dig not in num1_freq:
            return False
        # else if the frequency value is zero or less returns false
        elif num1_freq[dig] <= 0:
            return False
        # otherwise, decrease the frequency in dict
        else:
            num1_freq[dig] -= 1

    # if all checks pass returns true
    return True
