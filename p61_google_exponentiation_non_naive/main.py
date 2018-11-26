# This problem was asked by Google.
# Implement integer exponentiation. That is, implement the pow(x, y) function, where x and y are integers and returns x^y.
# Do this faster than the naive method of repeated multiplication.
# For example, pow(2, 10) should return 1024.

def pow(x, y):
    if y == 0:
        return 1
    if y == 1:
        return x

    square_result = x
    spare = 1
    while y > 1:
        if y % 2 == 0:
            square_result = square_result * square_result
            y = y // 2
        else:
            spare = square_result * spare
            y = y - 1
    return square_result * spare


print(pow(2, 10))