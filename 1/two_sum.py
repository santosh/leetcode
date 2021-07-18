# Given an array of integers and an integer target, return indices of the two numbers such that they add up to target.

def two_sum(nums, target):
    index = {}
    for i, num in enumerate(nums):
        if target - num in index:
            return [index[target - num], i]
        index[num] = i


