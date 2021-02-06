import unittest
from same_frequency import same_frequency


class TestSameFrequency(unittest.TestCase):
    def test_different_length(self):
      # test should return false if the inputs have diff length
        num1 = 22
        num2 = 222
        is_same = same_frequency(num1, num2)
        self.assertFalse(is_same)

    def test_one_freq(self):
      #  test should return false if the inputs has only one match
        num1 = 34
        num2 = 14
        is_same = same_frequency(num1, num2)
        self.assertFalse(is_same)

    def test_match_small_num(self):
      # test should return true when inputs has same frequency
        num1 = 182
        num2 = 281
        is_same = same_frequency(num1, num2)
        self.assertTrue(is_same)

    def test_match_big_num(self):
      # test should return true when inputs has same frequency
        num1 = 3589578
        num2 = 5879385
        is_same = same_frequency(num1, num2)
        self.assertTrue(is_same)
