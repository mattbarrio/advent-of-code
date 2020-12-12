#!/usr/bin/python
input_file = open('day2input.txt')
array = [];
matching = {};
n = 0;
def populate_array():
    try:
        for i, line in enumerate(input_file):
            array.append(line);
    finally:
        input_file.close();
        return array;

def diff(a, b):
    val = [i for i in range(len(a)) if a[i] != b[i]];
    if len(val) == 1:
        for i, k in zip(a,b):
            matching[i] = k;
        s = set.union(set(matching));
        print "single char diff found: {}".format(s);

array = populate_array();
# print array;
for a in array:
    for b in array:
        diff(a,b);
