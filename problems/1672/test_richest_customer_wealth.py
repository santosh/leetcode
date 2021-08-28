from richest_customer_wealth import maximum_wealth


def test_set_1():
    array = [[1,2,3],[3,2,1]]
    want = 6
    got = maximum_wealth(array)

    assert got == want

def test_set_2():
    array = [[1,5],[7,3],[3,5]]
    want = 10
    got = maximum_wealth(array)

    assert got == want
