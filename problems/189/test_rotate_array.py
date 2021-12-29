from rotate_array import Solution


def test__rotate_by_3():
    array = [1,2,3,4,5,6,7]
    want = [5,6,7,1,2,3,4]
    got = Solution().rotate(array, 3)

    assert got == want

def test__rotate_by_2():
    array = [-1,-100,3,99]
    want = [3,99,-1,-100]
    got = Solution().rotate(array, 2)

    assert got == want

def test__rotate_by_2_v2():
    array = [1,2,3,4,5,6,7]
    want = [6,7,1,2,3,4,5]
    got = Solution().rotate(array, 2)

    assert got == want

def test__rotate_by_3_v2():
    array = [1,2]
    want = [2,1]
    got = Solution().rotate(array, 3)

    assert got == want
