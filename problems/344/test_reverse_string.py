import reverse_string


def test__reverse_santosh():
    string = ["s", "a", "n", "t", "o", "s", "h"]
    want = ["h", "s", "o", "t", "n", "a", "s"]
    got = reverse_string.reverse_string(string)

    assert want == got


def test__reverse_cartoon():
    string = ["c", "a", "r", "t", "o", "o", "n"]
    want = ["n", "o", "o", "t", "r", "a", "c"]
    got = reverse_string.reverse_string(string)

    assert want == got


def test__reverse_pointer():
    string = ["p", "o", "i", "n", "t", "e", "r"]
    want = ["r", "e", "t", "n", "i", "o", "p"]
    got = reverse_string.reverse_string(string)

    assert want == got


import reverse_string


def test__reverse_santosh_two_pointers():
    string = ["s", "a", "n", "t", "o", "s", "h"]
    want = ["h", "s", "o", "t", "n", "a", "s"]
    got = reverse_string.reverse_string_two_pointers(string)

    assert want == got


def test__reverse_cartoon_two_pointers():
    string = ["c", "a", "r", "t", "o", "o", "n"]
    want = ["n", "o", "o", "t", "r", "a", "c"]
    got = reverse_string.reverse_string_two_pointers(string)

    assert want == got


def test__reverse_pointer_two_pointers():
    string = ["p", "o", "i", "n", "t", "e", "r"]
    want = ["r", "e", "t", "n", "i", "o", "p"]
    got = reverse_string.reverse_string_two_pointers(string)

    assert want == got
