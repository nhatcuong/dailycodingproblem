# Given a function that generates perfectly random numbers between 1 and k (inclusive), where k is an input, write a function that shuffles a deck of cards represented as an array using only swaps.
# It should run in O(N) time.

from random import randint

def shuffle(input):
    n = len(input)
    for i in range(0, n-1):
        j = randint(i, n-1)
        input[i], input[j] = input[j], input[i]
    return input
