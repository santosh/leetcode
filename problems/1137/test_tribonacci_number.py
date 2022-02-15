from tribonacci_number import Solution


def test_set_1():
    input = 4
    want = 4
    got = Solution().tribonacci(input)

    assert got == want

def test_set_2():
    input = 25
    want = 1389537
    got = Solution().tribonacci(input)

    assert got == want

# def test_set_3():
#     input = 4
#     want = 3
#     got = Solution().tribonacci(input)

#     assert got == want

# def test_set_4():
#     input = 5
#     want = 5
#     got = Solution().tribonacci(input)

#     assert got == want

