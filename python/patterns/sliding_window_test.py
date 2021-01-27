import unittest

from sliding_window import max_subarray_sum


class TestMaxSubarraySum(unittest.TestCase):
    def test_empty_list(self):
        # test if the list is with diff length of the nums should return None
        test = []
        nums = 4
        got = max_subarray_sum(test, nums)
        self.assertIsNone(got)

    def test_list_with_single_nums(self):
        # test if the list has max int 6 with nums of 1
        test = [4, 2, 1, 6]
        nums = 1
        got = max_subarray_sum(test, nums)
        want = 6
        self.assertEqual(got, want)

    def test_list_with_multiple_nums(self):
        # test if the list has max of 13 with nums of 4
        test = [4, 2, 1, 6, 2]
        nums = 4
        got = max_subarray_sum(test, nums)
        want = 13
        self.assertEqual(got, want)

    def test_long_list_with_multiple_nums(self):
        # test if the list has max of 10 with nums of 2
        test = [1, 2, 5, 2, 8, 1, 5]
        nums = 2
        got = max_subarray_sum(test, nums)
        want = 10
        self.assertEqual(got, want)

    def test_long_list_with_long_multiple_nums(self):
        # test if the list has max of 10 with nums of 2
        test = [1, 2, 5, 2, 8, 1, 5]
        nums = 4
        got = max_subarray_sum(test, nums)
        want = 17
        self.assertEqual(got, want)