from square_of_sorted_array import square_of_sorted_array as square_sorted


def test_set_1():
    array = [-4,-1,0,3,10]
    want = [0,1,9,16,100]
    got = square_sorted(array)

    assert got == want

def test_set_2():
    array = [-7,-3,2,3,11]
    want = [4,9,9,49,121]
    got = square_sorted(array)

    assert got == want
