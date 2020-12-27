# checkAnagrams check if two strings are anagrams O(n)
def checkAnagrams(firstString, secondString):
    # check if the length is same
    if len(firstString) != len(secondString):
        return False

    # create a lookup dict for record the frequency of letters
    freqFirstString = {}

    # loop thru the first string and increase the count for each letter
    for letter in firstString:
        if letter in freqFirstString:
            freqFirstString[letter] += 1
        else:
            freqFirstString[letter] = 1

    # loop thru the second string
    for letter in secondString:
        if letter in freqFirstString:
            # check if the letter is present on freq dict
            if freqFirstString[letter] == 0:
                # if freq is 0 than return false
                return False
            else:
                # decrease the count of the frequency
                freqFirstString[letter] -= 1
        else:
            return False

    return True


checkAnagrams('anagram', 'nagaram')
