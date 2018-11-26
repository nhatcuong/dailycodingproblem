def count_path(width, height):
    # instantiate matrix
    M = [[0 for _ in range(height)] for _ in range(width)]
    # fill the last column
    for i in reversed(range(height)):
        M[i][height-1] = 1
    # fill the last row
    for j in reversed(range(width)):
        M[width-1][j] = 1
    # fill all column from right to left
    for j in reversed(range(width-1)):
        for i in reversed(range(height-1)):
            M[i][j] = M[i+1][j] + M[i][j+1]
    # in each column from bottom to top
    return M[0][0]


print(count_path(5, 5))
