# Given an undirected graph represented as an adjacency matrix and an integer k, write a function to determine whether each vertex in the graph can be colored such that no two adjacent vertices share the same color using at most k colors.

# 2 possible optimization compared to what I have below:
# - compute all vertices in a list, so that we just need to go through this vertices list
# - add an "adjacent_colors" array that maintain the adjacent colors to any node so that we don't have to go through
# all the matrix again
# Current solution has (kN)^N in complexity
# With optimizations above, we can have k^N running time

# !!! This code has not been tested


class NoUncoloredVertex(Exception):
    pass


def can_color(adjacent_matrix, colors_count):
    colorization_matrix = [[0]*len(adjacent_matrix) for _ in range(len(adjacent_matrix))]
    return _can_color(adjacent_matrix, colorization_matrix, colors_count)


def _can_color(adjacent_matrix, colorization_matrix, colors_count):
    # colorization is temporary, and has len < len(adjacent_matrix)
    # colorization is full when len == len
    try:
        i, j = _get_possible_next_vertex(adjacent_matrix, colorization_matrix)
    except NoUncoloredVertex:
        return True

    possible_next_colors = _get_possible_next_colors(adjacent_matrix, colorization_matrix, colors_count)
    if not possible_next_colors:
        return False

    possible_next_temp_colorizations = _build_next_colorizations(colorization_matrix, i, j, possible_next_colors)
    return any([
        can_color(adjacent_matrix, c, colors_count) for c in possible_next_temp_colorizations
    ])


def _get_possible_next_vertex(adjacent_matrix, colorization_matrix):
    for i in range(len(adjacent_matrix)):
        for j in range(i+1, len(adjacent_matrix)):
            if adjacent_matrix[i][j] == 1 and colorization_matrix[i][j] == 0:
                return i, j
    raise NoUncoloredVertex


def _get_possible_next_colors(adjacent_matrix, colorization, vertex_i, vertex_j, colors_count):
    all_colors = {k+1 for k in range(colors_count)}
    adjacent_vertex_i_colors = {
        colorization[vertex_i][j] for j in range(vertex_i+1, len(adjacent_matrix))
    }
    adjacent_vertex_j_colors = {
        colorization[i][vertex_j] for i in range(0, vertex_j)
    }
    return all_colors - adjacent_vertex_i_colors - adjacent_vertex_j_colors


def _build_next_colorizations(colorization_matrix, vertex_i, vertex_j, next_colors):
    results = []
    for c in next_colors:
        new_colorization_matrix = [row[:] for row in colorization_matrix]
        new_colorization_matrix[vertex_i][vertex_j] = c
        results = results.append(c)
    return results
