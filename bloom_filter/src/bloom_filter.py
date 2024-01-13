import mmh3
from bitarray import bitarray
from functools import partial
import random
import argparse


bloom_filter = bitarray(2**32)

class BloomFilter:
    def __init__(self, size, num_hashes):
        self.bit_array = bitarray(size)
        self.bit_array.setall(0)
        self.hashes = [partial(mmh3.hash, seed=random.randint(1,1000)) for _ in range(num_hashes)]
    
    
    def insert(self, key):
        # bit_positions = []
        for hash_function in self.hashes:
            value = hash_function(key)
            self.bit_array[value % len(self.bit_array)] = 1
        return self.bit_array

    def check_key(self, key):
        for hash_function in self.hashes:
            value = hash_function(key)
            if self.bit_array[value % len(self.bit_array)] == 0:
                return False
        return True


    
if __name__ == "__main__":
    
    parser = argparse.ArgumentParser()
    import sys
    parser.add_argument("-s", "--size", help = "Bloom filter size")
    parser.add_argument("-n", "--num_hash", help = "Number of hashes to use")
    args = parser.parse_args()
    b = BloomFilter(int(args.size), int(args.num_hash))
    
    while True:
        user_input = input("User commands ADD: and CHECK: \n")
        action, key = user_input.split(":")
        if action.strip().upper() == "ADD":
            b.insert(key.strip())
        elif action.strip().upper() == "CHECK":
            print(b.check_key(key.strip()))
        else:
            print("Invalid input")