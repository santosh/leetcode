class Solution:
    def __init__(self):
        self.lookup_table = dict()

        # base case
        self.lookup_table[0] = 0
        self.lookup_table[1] = 1
        self.lookup_table[2] = 1

    def tribonacci(self, n: int) -> int:
        if n in self.lookup_table:
            return self.lookup_table[n]
        else :
            self.lookup_table[n] = self.tribonacci(n-3) + self.tribonacci(n-2) + self.tribonacci(n-1)
            return self.lookup_table[n]

if __name__ == "__main__":
    s = Solution()
    print(s.tribonacci(20))

# 136174

# 145996

# 140648

