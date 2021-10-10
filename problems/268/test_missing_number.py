from missing_number import Solution


def test_set_1():
    input = [3,0,1]
    want = 2
    got = Solution().missingNumber(input)

    assert got == want

def test_set_2():
    input = [0, 1]
    want = 2
    got = Solution().missingNumber(input)

    assert got == want

def test_set_3():
    input = [9,6,4,2,3,5,7,0,1]
    want = 8
    got = Solution().missingNumber(input)

    assert got == want

def test_set_4():
    input = [0]
    want = 1
    got = Solution().missingNumber(input)

    assert got == want

def test_set_5():
    input = [1]
    want = 0
    got = Solution().missingNumber(input)

    assert got == want

