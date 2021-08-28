def maximum_wealth(arrays):
    sum_of_array = []
    for array in arrays:
        sum_of_array.append(sum(array))

    return max(sum_of_array)
