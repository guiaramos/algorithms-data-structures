import unittest
from frequency_counter import sameArraySquared


class TestSameArraySquared(unittest.TestCase):
    def test_array(self):
        # Test if arrays have the dif length should return false
        firstArray = [1, 2, 3]
        secondArray = [1, 2, 3, 4]

        isSame = sameArraySquared(firstArray, secondArray)
        self.assertFalse(isSame)

    def test_array_with_no_squared(self):
        # Test if the array with no correct squared returns false
        firstArray = [1, 2, 3]
        secondArray = [1, 2, 3]

        isSame = sameArraySquared(firstArray, secondArray)
        self.assertFalse(isSame)
