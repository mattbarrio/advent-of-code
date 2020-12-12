#!/usr/bin/python3

import string
# from tools.arr import *

input_file = open('day1.txt')

def populate_array(input_file):
    array = [];
    try:
        for i, line in enumerate(input_file):
            array.append(int(line));
    finally:
        input_file.close();
        return array;

array = populate_array(input_file);

a = 0;
sum_array = [];
print(array);
for i in array:
    fuel = i;
    negative_fuel_found = False;
    while not negative_fuel_found:
        calc = fuel // 3 - 2
        if calc > 0:
            sum_array.append(calc);
            fuel = calc
        else:
            negative_fuel_found = True;
            #break;

print('The answer is {}'.format(sum(sum_array)))

# if __name__ == '__main__':
#    main()
