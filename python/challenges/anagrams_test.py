import unittest
from anagrams import checkAnagrams


class TestAnagrams(unittest.TestCase):
    def test_empty_string(self):
      # Test if pass empty string it should return true
        firstString = ''
        secondString = ''

        isValid = checkAnagrams(firstString, secondString)

        self.assertTrue(isValid)

    def test_mismatch_length(self):
      # Test if pass strings with different length it should return false
        firstString = 'amazon'
        secondString = 'amazo'

        isValid = checkAnagrams(firstString, secondString)

        self.assertFalse(isValid)

    def test_mismatch_strings(self):
      # Test if mismatch string should return false
        firstString = 'car'
        secondString = 'rat'

        isValid = checkAnagrams(firstString, secondString)

        self.assertFalse(isValid)

    def test_wrong_frequency(self):
      # Test if pass strings with same words but different frequency should return false
        firstString = 'aaz'
        secondString = 'zza'

        isValid = checkAnagrams(firstString, secondString)

        self.assertFalse(isValid)

    def test_correct_anagram(self):
      # Test if anagram and nagaram is an Anagram it should return true
        firstString = 'anagram'
        secondString = 'nagaram'

        isValid = checkAnagrams(firstString, secondString)

        self.assertTrue(isValid)
