from contains_duplicate import Solution

def test_set_1():
    nums = [1,2,3,1]
    want = True
    got = Solution().containsDuplicate(nums)

    assert got == want

def test_set_2():
    nums = [-1,0,3,5,9,12]
    want = False
    got = Solution().containsDuplicate(nums)

    assert got == want

def test_set_3():
    nums = [1,1,1,3,3,4,3,2,4,2]
    want = True
    got = Solution().containsDuplicate(nums)

    assert got == want
