from binary_search import Solution

def test_set_1():
    nums = [-1,0,3,5,9,12]
    target = 9
    want = 4
    got = Solution().search(nums, target)

    assert got == want

def test_set_2():
    nums = [-1,0,3,5,9,12]
    target = 2
    want = -1
    got = Solution().search(nums, target)

    assert got == want
