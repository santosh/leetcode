from defang_ip_address import Solution


def test_set_1():
    addr = "1.1.1.1"
    want = "1[.]1[.]1[.]1"
    got = Solution().defangIPaddr(addr)

    assert got == want

def test_set_2():
    addr = "255.100.50.0"
    want = "255[.]100[.]50[.]0"
    got = Solution().defangIPaddr(addr)

    assert got == want

