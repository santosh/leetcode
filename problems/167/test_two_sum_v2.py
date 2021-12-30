import two_sum_v2


def test_set_1():
    array = [2, 7, 11, 15]
    want = [0, 1]
    got = two_sum_v2.two_sum(array, 9)

    assert got == want

def test_set_2():
    array = [3, 3]
    want = [0, 1]
    got = two_sum_v2.two_sum(array, 6)

    assert got == want

def test_set_3():
    array = [2, 7, 11, 15]
    want = [1, 2]
    got = two_sum_v2.two_sum(array, 9)

    assert got == want

def test_set_4():
    array = [2, 3, 4]
    want = [1, 3]
    got = two_sum_v2.two_sum(array, 6)

    assert got == want

def test_set_5():
    array = [-1, 0]
    want = [1, 2]
    got = two_sum_v2.two_sum(array, -1)

    assert got == want
