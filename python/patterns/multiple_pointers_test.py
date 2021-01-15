import unittest
from multiple_pointers import sumZero


class TestSumZero(unittest.TestCase):
    def test_empty_array(self):
        # test if the array of numbers have length 0 returns none
        test = []
        got = sumZero(test)
        self.assertIsNone(got)

    def test_no_pair_match(self):
        # test if the array with no pair where sum is zero returns none
        test = [-3, -2, - 1, 0, 4, 5, 6]
        got = sumZero(test)
        self.assertIsNone(got)

    def test_first_pair_match(self):
        # test if the array with first pair where sum is zero retruns the pair
        test = [-3, -2, -1, 0, 1, 2, 3]
        got = sumZero(test)
        want = [-3, 3]

        self.assertCountEqual(got, want)

    def test_middle_pair_match(self):
        # test if the array with middle pair where sum is zero retruns the pair
        test = [-9, -2, -1, 0, 1, 2, 8]
        got = sumZero(test)
        want = [-2, 2]

        self.assertCountEqual(got, want)
