import unittest

import duplicatenum


class TestFindDuplicate(unittest.TestCase):
    def test_with_one_duplicate(self):
        array = [2, 1, 3, 4, 2]
        want = 2
        got = duplicatenum.find_duplicate(array)

        self.assertEqual(got, want)

    def test_with_two_duplicate(self):
        array = [2, 1, 3, 4, 2, 2]
        want = 2
        got = duplicatenum.find_duplicate(array)

        self.assertEqual(got, want)
