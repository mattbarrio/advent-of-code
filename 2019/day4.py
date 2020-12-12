#!/usr/bin/python3

import string
from collections import Counter


def check1(s):
    return list(s) == sorted(s) and len(set(s)) < len(s)


def check2(s):
    return list(s) == sorted(s) and 2 in Counter(s).values()


print(sum(check1(str(x)) for x in range(359282, 820401)))
print(sum(check2(str(x)) for x in range(359282, 820401)))

# print('Found {} matches'.format(matches_found));
