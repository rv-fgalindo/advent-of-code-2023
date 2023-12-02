from typing import List
from file_reader import data_reader


def part_one(data_list: List[str]):    
    sum = 0
    for data_str in data_list:
        first_digit = ""
        last_digit = ""
        found_digit = False
        for char in data_str:
            if char.isdigit() and not found_digit:
                first_digit = char
                last_digit = char
                found_digit = True
            elif char.isdigit():
                last_digit = char
        sum += int(first_digit + last_digit)
    print(sum)

def part_two(data_list: List[str]):
    digits = {"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}
    sum = 0
    for data_str in data_list:
        first_digit = "0"
        last_digit = "0"
        found_digit = False

        for char in data_str:
            if char.isdigit() and not found_digit:
                first_digit = char
                last_digit = char
                found_digit = True
            elif char.isdigit():
                last_digit = char

        highest_idx = -1
        lowest_idx = float('inf')
        for digit in digits.keys():
            idx = data_str.find(digit)
            while idx != -1:
                if idx > highest_idx:
                    no_nums_found = True
                    for char in data_str[idx:]:
                        if char.isdigit():
                            no_nums_found = False
                    if no_nums_found:
                        last_digit = digits[digit]
                        highest_idx = idx
                if idx < lowest_idx:
                    no_nums_found = True
                    for char in reversed(data_str[:idx+1]):
                        if char.isdigit():
                            no_nums_found = False
                    if no_nums_found:
                        first_digit = digits[digit]
                        lowest_idx = idx
                idx = data_str.find(digit, idx + 1)
        sum += int(first_digit + last_digit)
    print(sum)

data_list: List[str] = data_reader("Tavio\challenge1.txt")
# part_one(data_list)
part_two(data_list)

# part_two(["dttwonezbgmcseven5seven"])

# test_data = List[dict] = \
# ["1abc2",
# "pqr3stu8vwx",
# "a1b2c3d4e5f",
# "treb7uchet"]
# part_one(test_data)

# test_data_part_two = \
# ["two1nine",
# "eightwothree",
# "abcone2threexyz",
# "xtwone3four",
# "4nineeightseven2",
# "zoneight234",
# "7pqrstsixteen"]
# part_two(test_data_part_two)

# word += char
# word_in_digit = False
# for digit in digits.keys():
#     if word in digit:
#         word_in_digit = True
#     if word == digit and not found_digit:
#         first_digit = digits[digit]
#         last_digit = digits[digit]
#         found_digit = True
#         word = char
#     elif word == digit:
#         last_digit = digits[digit]
#         word = char
# if not word_in_digit:
#     word=char