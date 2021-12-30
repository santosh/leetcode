# Given an array of integers and an integer target, return indices of the two numbers such that they add up to target.

def two_sum_hashmap(nums, target):
    index = {}
    for i, num in enumerate(nums):
        if target - num in index:
            return [index[target - num], i]
        index[num] = i

def two_sum_two_pointer(nums, target):
    pointer_one = 0
    pointer_two = len(nums) - 1

    while pointer_one < pointer_two:
        sum = nums[pointer_one] + nums[pointer_two]

        if sum == target:
            return [pointer_one, pointer_two]
        elif sum < target:
            pointer_one += 1
        else:
            pointer_two -= 1

    return [-1, -1]
