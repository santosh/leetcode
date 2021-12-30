def two_sum(num_array, target):
    pointer_one = 0
    pointer_two = len(num_array) - 1

    while pointer_one < pointer_two:
        sum = num_array[pointer_one] + num_array[pointer_two]

        if sum == target:
            return [pointer_one + 1, pointer_two + 1]  # +1 is done as per description
        elif sum < target:
            pointer_one += 1
        else:
            pointer_two -= 1

