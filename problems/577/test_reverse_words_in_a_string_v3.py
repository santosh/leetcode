import reverse_words_in_a_string_v3 as reverse_words

def test_set_1():
    input_str = "Let's take LeetCode contest"
    want = "s'teL ekat edoCteeL tsetnoc"
    got = reverse_words.reverse_words(input_str)

    assert got == want

def test_set_2():
    input_str = "God Ding"
    want = "doG gniD"
    got = reverse_words.reverse_words(input_str)

    assert got == want

def test_set_3():
    input_str = "Reverse Words in a String"
    want = "esreveR sdroW ni a gnirtS"
    got = reverse_words.reverse_words(input_str)

    assert got == want
