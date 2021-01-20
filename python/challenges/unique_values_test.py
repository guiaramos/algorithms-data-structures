import unittest
from unique_values import countUniqueValues


class TestCountUniqueValues(unittest.TestCase):
    def test_empty_list(self):
        # Test should return 0 when empty list
        emptyList = []
        got = countUniqueValues(emptyList)
        want = 0
        self.assertEqual(got, want)

    def test_one_number_list(self):
        # Test should return 1 when list has 1 number
        numList = [1]
        got = countUniqueValues(numList)
        want = 1
        self.assertEqual(got, want)

    def test_two_numbers_list(self):
      # Test should return 2 when list has 2 unique numbers
        numList = [1, 2]
        got = countUniqueValues(numList)
        want = 2
        self.assertEqual(got, want)

    def test_multiple_numbers_list(self):
      # Test should return 2 when list has multiple number, but 2 unique numbers
        numList = [1, 1, 1, 1, 1, 2]
        got = countUniqueValues(numList)
        want = 2
        self.assertEqual(got, want)

    def test_large_number_list(self):
      # Test should return 7 when large number list, but 7 unique numbers
        numList = [1, 2, 3, 5, 5, 5, 7, 7, 12, 12, 13]
        got = countUniqueValues(numList)
        want = 7
        self.assertEqual(got, want)

    def test_negative_number_list(self):
      # Test should return 4 when list has negative and positive number, but 4 unique numbers
        numList = [-2, -1, -1, 0, 1]
        got = countUniqueValues(numList)
        want = 4
        self.assertEqual(got, want)
