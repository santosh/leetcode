from typing import List

class Solution:
    def search(self, nums: List[int], target: int) -> int:
        left_index = 0  # set at start of the sorted array
        right_index = len(nums) - 1  # set at the end of the sorted array

        while left_index <= right_index:
            mid_index = left_index + (right_index - left_index) // 2

            if target == nums[mid_index]:
                return mid_index

            if target < nums[mid_index]:
                right_index = mid_index - 1
            else:
                left_index = mid_index + 1

        return -1
