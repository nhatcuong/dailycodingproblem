# This problem was asked by Facebook.
# Given a multiset of integers, return whether it can be partitioned into two subsets whose sums are the same.
# For example, given the multiset {15, 5, 20, 10, 35, 15, 10}, it would return true, since we can split it up into {15, 5, 10, 15, 10} and {20, 35}, which both add up to 55.
# Given the multiset {15, 5, 20, 10, 35}, it would return false, since we can't split it up into two subsets that add up to the same sum.


def can_array_split_half_sum(input_array):
    # create matrix of size (sum(input_array)/2)+1, len(input_array)
    if sum(input_array) % 2 == 1:
        return False
    half = sum(input_array) // 2
    M = [[False for _ in range(len(input_array))] for _ in range(half+1)]
    # fill all first column (j=0)
    for i in range(half+1):
        M[i][0] = input_array[0] == i
    for j in range(1, len(input_array)):
        for i in range(half+1):
            if M[i][j-1]:
                M[i][j] = True
            if i - input_array[j] >= 0 and M[i-input_array[j]][j-1]:
                M[i][j] = True
        if M[half][j]:
            return True
    return False


print(can_array_split_half_sum([15, 5, 20, 10, 35, 15, 10]))
print(can_array_split_half_sum([15, 5, 20, 10, 35]))
