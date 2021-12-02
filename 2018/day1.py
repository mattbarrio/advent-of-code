#!/usr/bin/python
input_file = open('day1input.txt')
a = 0
frequency_array = []
frequencies = []
dup_frequency_found = False
l = 0


def check_frequency_list(frequencies, current):
    if current in frequencies:
        return True
    else:
        frequencies.append(current)
        return False


def populate_frequency_array():
    try:
        for i, line in enumerate(input_file):
            frequency_array.append(int(line))
    finally:
        input_file.close()
        return frequency_array


frequency_array = populate_frequency_array()

while not dup_frequency_found:
    l += 1
    print("loop {}".format(l))
    for i in frequency_array:
        a += i
        if check_frequency_list(frequencies, a):
            print("line: {} current: {} \n{}".format(i, a, frequencies))
            dup_frequency_found = True
            break
        # print "line: {} current: {} \n{}".format(i, a, frequencies);

print("answer: {}".format(a))
