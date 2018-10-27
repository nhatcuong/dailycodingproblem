# Given an array of numbers, find the maximum sum of any contiguous subarray of the array.
# For example, given the array [34, -50, 42, 14, -5, 86], the maximum sum would be 137, since we would take elements 42, 14, -5, and 86.
# Given the array [-5, -1, -8, -9], the maximum sum would be 0, since we would not take any elements.
# Do this in O(N) time.

def biggest_consegutive_sum(input):
    biggest = 0
    current_biggest_extensible = 0
    for i in input:
        if current_biggest_extensible + i > 0:
            current_biggest_extensible = current_biggest_extensible + i
        else:
            current_biggest_extensible = 0
        if current_biggest_extensible > biggest:
            biggest = current_biggest_extensible
    return biggest


print('{}'.format(biggest_consegutive_sum([34, -50, 42, 14, -5, 86])))
print('{}'.format(biggest_consegutive_sum([-5, -1, -8, -9])))
print('{}'.format(biggest_consegutive_sum([19, -17])))