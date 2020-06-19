from typing import List

# Given an array nums containing n + 1 integers where each integer
# is between 1 and n (inclusive), prove that at least one duplicate
# number must exist. Assume that there is only one duplicate
# number, find the duplicate one.


def find_duplicate(nums: List[int]) -> int:
    """find_duplicate uses binary search to find duplicate numbers."""

    start, end = 1, len(nums)

    while start <= end:
        mid = start + (end - start) // 2
        cnt = 0

        # Find how many numbers are on <= of mid.
        # If there are more than mid number of nums means
        # there is a duplicate in left
        for n in nums:
            if n <= mid:
                cnt = cnt + n

        if cnt > mid:
            end = mid - 1
        else:
            start = mid + 1

    return start
