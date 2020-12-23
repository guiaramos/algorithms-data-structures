
# sameArraySquared should check if two array have the correspoding squared number
def sameArraySquared(arrOne, arrTwo):
    # check if the length of the arrays are same
    if len(arrOne) != len(arrTwo):
        return False

    # Make a dict for each arrays to save the frequency
    freqArrayOne = {}
    freqArrayTwo = {}

    # Interate thru the first array and save the frequency
    for key in arrOne:
        if key in freqArrayOne:
            freqArrayOne[key] += 1
        else:
            freqArrayOne[key] = 1

    # Interate thru the second array and save the frequency
    for key in arrTwo:
        if key in freqArrayTwo:
            freqArrayTwo[key] += 1
        else:
            freqArrayTwo[key] = 1

    # Interate thru the first frequency dict
    for key in freqArrayOne:
        # Check if the key squared is on second array frequency
        if key ** 2 not in freqArrayTwo:
            return False

            # Check if the value is same
        if freqArrayTwo[key ** 2] != freqArrayOne[key]:
            return False

    return True
