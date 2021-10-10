from typing import List


class Solution:
    def missingNumber(self, nums: List[int]) -> int:
        len_of_nums = len(nums)
        total_sum = len_of_nums * (len_of_nums+1) // 2

        return total_sum - sum(nums)

if __name__ == "__main__":
    s = Solution()

    n = s.missingNumber([9,6,4,2,3,5,7,0,1])
    print(n)
