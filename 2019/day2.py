#!/usr/bin/python3

import string
# from tools.arr import *

# input_file = open('day2.txt')
# array = populate_array(input_file);


def set_array(noun, verb):
    return [1, noun, verb, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 1, 10, 19, 1, 6, 19, 23, 1, 10, 23, 27, 2, 27, 13, 31, 1, 31, 6, 35, 2, 6, 35, 39, 1, 39, 5, 43, 1, 6, 43, 47, 2, 6, 47, 51, 1, 51, 5, 55, 2, 55, 9, 59, 1, 6, 59, 63, 1, 9, 63, 67, 1, 67, 10, 71, 2, 9, 71, 75, 1, 6, 75, 79, 1, 5, 79, 83, 2, 83, 10, 87, 1, 87, 5, 91, 1, 91, 9, 95, 1, 6, 95, 99, 2, 99, 10, 103, 1, 103, 5, 107, 2, 107, 6, 111, 1, 111, 5, 115, 1, 9, 115, 119, 2, 119, 10, 123, 1, 6, 123, 127, 2, 13, 127, 131, 1, 131, 6, 135, 1, 135, 10, 139, 1, 13, 139, 143, 1, 143, 13, 147, 1, 5, 147, 151, 1, 151, 2, 155, 1, 155, 5, 0, 99, 2, 0, 14, 0]


def set_opcode():
    return 0


array = set_array(0, 0)
opcode = set_opcode()

for noun in range(99):  # reversed(list(range(99))):
    for verb in range(99):
        for i, val in enumerate(array):
            if array[opcode] == 99:
                # print(array[0], array[1], array[2]);
                if array[0] == 19690720:
                    print(100 * array[1] + array[2])
                    break
                break
            lookup1 = array[opcode + 1]
            lookup2 = array[opcode + 2]
            if array[opcode] == 1:
                array[array[opcode + 3]] = array[lookup1] + array[lookup2]
            if array[opcode] == 2:
                array[array[opcode + 3]] = array[lookup1] * array[lookup2]
            opcode = opcode + 4
        verb += 1
        # array[2] = verb;
        array = set_array(noun, verb)
        opcode = set_opcode()
    noun += 1
    # array[1] += noun;
    array = set_array(noun, verb)
    opcode = set_opcode()
