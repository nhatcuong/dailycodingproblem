# Implement an LRU (Least Recently Used) cache. It should be able to be initialized with a cache size n, and contain the following methods:
#
# set(key, value): sets key to value. If there are already n items in the cache and we are adding a new item, then it should also remove the least recently used item.
# get(key): gets the value at key. If no such key exists, return null.
# Each operation should run in O(1) time.

class LruCacheNode:
    def __init__(self, key,  value, sooner_node, later_node):
        self.key = key
        self.value = value
        self.sooner_node = sooner_node
        self.later_node = later_node


class LruCache:
    def __init__(self, capacity):
        self.capacity = capacity
        self.dict = dict()
        self.most_recently_used_node = None
        self.least_recently_used_node = None
        self.size = 0

    def get(self, key):
        if not self.dict.get(key, None):
            return None
        node = self.dict[key]
        if node == self.least_recently_used_node and node.later_node:
            self.least_recently_used_node = node.later_node
        if node != self.most_recently_used_node:
            sooner_node = node.sooner_node
            later_node = node.later_node
            later_node.sooner_node = sooner_node
            if sooner_node:
                sooner_node.later_node = later_node
            self.most_recently_used_node = node

        return node.value

    def set(self, key, value):
        # key already exists
        if self.dict.get(key, None):
            node = self.dict[key]
            node.value = value
            if node == self.least_recently_used_node and node.later_node:
                self.least_recently_used_node = node.later_node
            if node != self.most_recently_used_node:
                sooner_node = node.sooner_node
                later_node = node.later_node
                later_node.sooner_node = node.sooner_node
                if sooner_node:
                    sooner_node.later_node = later_node
                self.most_recently_used_node = node
            return

        # new key
        node = LruCacheNode(key, value, self.most_recently_used_node, None)
        if self.most_recently_used_node:
            self.most_recently_used_node.later_node = node
        self.most_recently_used_node = node
        self.least_recently_used_node = self.least_recently_used_node or node
        self.dict[key] = node
        if self.size < self.capacity:
            self.size += 1
        else:
            self.dict.pop(self.least_recently_used_node.key)
            self.least_recently_used_node = self.least_recently_used_node.later_node


cache = LruCache(capacity=5)
cache.set(1, True)
cache.set(2, True)
cache.set(3, True)
cache.set(4, True)
cache.set(5, True)
print('get 5 {}'.format(cache.get(5)))
print('get 1 {}'.format(cache.get(1)))
print('get 9 {}'.format(cache.get(9)))
cache.set(6, True)
print('get 2 {}'.format(cache.get(2)))
print('get 3 {}'.format(cache.get(3)))
cache.set(4, False)
cache.set(7, True)
print('get 4 {}'.format(cache.get(4)))
print('get 3 {}'.format(cache.get(3)))
print('get 1 {}'.format(cache.get(1)))
print('get 5 {}'.format(cache.get(5)))

# set 1, 2, 3, 4, 5 => lru = 1
# get 5
# get 1
# set 6 => remove 2
# get 3
# get 5
# set 7 => remove 4
# get 4 -> null