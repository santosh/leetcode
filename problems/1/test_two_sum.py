import pytest
import two_sum


def test_set_1_with_hashmap():
    array = [2, 7, 11, 15]
    want = [0, 1]
    got = two_sum.two_sum_hashmap(array, 9)

    assert got == want

def test_set_2_with_hashmap():
    array = [3, 2, 4]
    want = [1, 2]
    got = two_sum.two_sum_hashmap(array, 6)

    assert got == want

def test_set_3_with_hashmap():
    array = [3, 3]
    want = [0, 1]
    got = two_sum.two_sum_hashmap(array, 6)

    assert got == want

def test_set_1_with_two_pointer():
    array = [2, 7, 11, 15]
    want = [0, 1]
    got = two_sum.two_sum_two_pointer(array, 9)

    assert got == want

@pytest.mark.xfail(reason='two pointer requires sorted array')
def test_set_2_with_two_pointer():
    array = [3, 2, 4]
    want = [1, 2]
    got = two_sum.two_sum_two_pointer(array, 6)

    assert got == want

def test_set_3_with_two_pointer():
    array = [3, 3]
    want = [0, 1]
    got = two_sum.two_sum_two_pointer(array, 6)

    assert got == want
