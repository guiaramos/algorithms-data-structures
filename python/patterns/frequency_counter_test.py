import unittest
from frequency_counter import sameArraySquared


class TestSameArraySquared(unittest.TestCase):
    def test_wrong_length(self):
        # Test if arrays have the dif length should return false
        firstArray = [1, 2, 3]
        secondArray = [1, 2, 3, 4]

        isSame = sameArraySquared(firstArray, secondArray)
        self.assertFalse(isSame)

    def test_no_squared(self):
        # Test if the array with no correct squared returns false
        firstArray = [1, 3, 3]
        secondArray = [1, 2, 2]

        isSame = sameArraySquared(firstArray, secondArray)
        self.assertFalse(isSame)

    def test_wrong_squared(self):
        # Test if the array has same frequency but wrong squared
        firstArray = [1, 2, 1]
        secondArray = [1, 4, 6]

        isSame = sameArraySquared(firstArray, secondArray)
        self.assertFalse(isSame)

    def test_correct_squared_ordered(self):
        # Test with correct array and squared - ordered
        firstArray = [3, 4, 5]
        secondArray = [9, 16, 25]

        isSame = sameArraySquared(firstArray, secondArray)
        self.assertTrue(isSame)

    def test_correct_squared_unordered(self):
        # Test with correct array and squared - unordered
        firstArray = [4, 5, 3]
        secondArray = [16, 25, 9]

        isSame = sameArraySquared(firstArray, secondArray)
        self.assertTrue(isSame)
