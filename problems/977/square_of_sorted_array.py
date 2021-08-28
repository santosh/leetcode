def square_of_sorted_array(array):
    for i, num in enumerate(array):
        array[i] = num * num

    return sorted(array)
