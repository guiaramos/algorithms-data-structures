import unittest
from average_pair import average_pair


class TestAveragePair(unittest.TestCase):
    def test_empty_list(self):
        # should return false if the list is empty
        int_list = []
        target_avg = 4
        is_pair = average_pair(int_list, target_avg)
        self.assertFalse(is_pair)

    def test_with_short_list(self):
      # should return true with small list of integers
        int_list = [1, 2, 3]
        target_avg = 2.5
        is_pair = average_pair(int_list, target_avg)
        self.assertTrue(is_pair)

    def test_with_long_list(self):
      # should return true with long list of integers
        int_list = [1, 3, 3, 5, 6, 7, 10, 12, 19]
        target_avg = 8
        is_pair = average_pair(int_list, target_avg)
        self.assertTrue(is_pair)

    def test_with_long_list_no_target(self):
      # should return false with long list of integers but no target
        int_list = [-1, 0, 3, 4, 5, 6]
        target_avg = 4.1
        is_pair = average_pair(int_list, target_avg)
        self.assertFalse(is_pair)
