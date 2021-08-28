from fibonacci_number import Solution


def test_set_1():
    input = 2
    want = 1
    got = Solution().fib(input)

    assert got == want

def test_set_2():
    input = 3
    want = 2
    got = Solution().fib(input)

    assert got == want

def test_set_3():
    input = 4
    want = 3
    got = Solution().fib(input)

    assert got == want

def test_set_4():
    input = 5
    want = 5
    got = Solution().fib(input)

    assert got == want

