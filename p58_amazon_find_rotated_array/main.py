# An sorted array of integers was rotated an unknown number of times.
# Given such an array, find the index of the element in the array in faster than linear time. If the element doesn't exist in the array, return null.
# For example, given the array [13, 18, 25, 2, 8, 10] and the element 8, return 4 (the index of 8 in the array).


import math


def find_in_rotated_array(array, x):
    return fra_bis(array, 0, len(array)-1, x)


def fra_bis(array, from_index, to_index, x):
    if to_index == from_index:
        return from_index if array[from_index] == x else -1

    if to_index == from_index + 1:
        if array[from_index] == x:
            return from_index
        if array[to_index] == x:
            return to_index
        return -1

    mid_index = int(math.floor((to_index + from_index) / 2))
    is_first_half_in_order = array[from_index] < array[mid_index]
    if is_first_half_in_order:
        can_find_in_first_half = array[mid_index] >= x >= array[from_index]
        if can_find_in_first_half:
            return fra_bis(array, from_index, mid_index, x)
        return fra_bis(array, mid_index, to_index, x)
    else:
        can_find_in_second_half = array[to_index] >= x >= array[mid_index]
        if can_find_in_second_half:
            return fra_bis(array, mid_index, to_index, x)
        return fra_bis(array, from_index, mid_index, x)


print(find_in_rotated_array([13, 18, 25, 2, 8, 10], 8))
