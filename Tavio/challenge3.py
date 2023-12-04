from typing import Dict, List
from file_reader import data_reader

used_indexes: Dict[int, List[int]] = dict()

def check_top_or_bottom(sym_idx: int, line: str, list_idx: int):
    if list_idx not in used_indexes:
        used_indexes[list_idx] = []
    left_num = ""
    middle_num = ""
    right_num = ""
    if sym_idx != 0 and line[sym_idx - 1].isdigit():  # Left side
        num = ""
        for char_idx in range(sym_idx - 1, -1, -1):
            current_num = line[char_idx]
            if current_num.isdigit() and char_idx not in used_indexes[list_idx]:
                num += line[char_idx]
                used_indexes[list_idx].append(char_idx)
            else:
                break
        left_num = num[::-1]
    if line[sym_idx].isdigit():  # Middle
        num = ""
        for char_idx in range(sym_idx, len(line)):
            current_num = line[char_idx]
            if current_num.isdigit() and char_idx not in used_indexes[list_idx]:
                num += line[char_idx]
                used_indexes[list_idx].append(char_idx)
            else:
                break
        middle_num = num
    if sym_idx != len(line) - 1 and line[sym_idx + 1].isdigit() and len(middle_num) < 1:  # Right side
        num = ""
        for char_idx in range(sym_idx + 1, len(line)):
            current_num = line[char_idx]
            if current_num.isdigit() and char_idx not in used_indexes[list_idx]:
                num += line[char_idx]
                used_indexes[list_idx].append(char_idx)
            else:
                break
        right_num = num
    if left_num and middle_num:
        return int(left_num + middle_num)
    if left_num or right_num:
        l_value = int(left_num) if left_num != "" else 0
        r_value = int(right_num) if right_num != "" else 0
        return l_value + r_value
    if middle_num:
        return int(middle_num)
    return 0


    

def check_sides(sym_idx: int, line: str, list_idx: int):
    if list_idx not in used_indexes:
        used_indexes[list_idx] = []
    nums_found = []
    if sym_idx != 0 and line[sym_idx - 1].isdigit():  # Left side
        num = ""
        for char_idx in range(sym_idx - 1, -1, -1):
            current_num = line[char_idx]
            if current_num.isdigit() and char_idx not in used_indexes[list_idx]:
                num += line[char_idx]
                used_indexes[list_idx].append(char_idx)
            else:
                break
        value = int(num[::-1]) if num != "" else 0
        nums_found.append(value)
    if sym_idx != len(line) - 1 and line[sym_idx + 1].isdigit():  # Right side
        num = ""
        for char_idx in range(sym_idx + 1, len(line)):
            current_num = line[char_idx]
            if current_num.isdigit() and char_idx not in used_indexes[list_idx]:
                num += line[char_idx]
                used_indexes[list_idx].append(char_idx)
            else:
                break
        value = int(num) if num != "" else 0
        nums_found.append(value)
    return sum(nums_found)


def search_symbols(engine_line: str, list_idx: int, engine_list: List[str]):
    nums_found = 0
    for char_idx, char in enumerate(engine_line):
        if not char.isdigit() and char != ".":
            print(char, char_idx)
            print(engine_list[list_idx], list_idx)
            if list_idx != len(engine_list) - 1:  # Check bottom
                nums_found += check_top_or_bottom(char_idx, engine_list[list_idx + 1], list_idx + 1)
            if list_idx != 0:  # Check top
                nums_found += check_top_or_bottom(char_idx, engine_list[list_idx - 1], list_idx - 1)
            nums_found += check_sides(char_idx, engine_line, list_idx)
    return nums_found


def part_one(engine_data: List[str]):
    total = 0
    for idx, line in enumerate(engine_data):
        total += search_symbols(line, idx, engine_data)
    print(total)


def check_top_or_bottom_part_two(sym_idx: int, line: str, list_idx: int):
    left_num = ""
    middle_num = ""
    right_num = ""
    if sym_idx != 0 and line[sym_idx - 1].isdigit():  # Left side
        num = ""
        for char_idx in range(sym_idx - 1, -1, -1):
            current_num = line[char_idx]
            if current_num.isdigit():
                num += line[char_idx]
            else:
                break
        left_num = num[::-1]
    if line[sym_idx].isdigit():  # Middle
        num = ""
        for char_idx in range(sym_idx, len(line)):
            current_num = line[char_idx]
            if current_num.isdigit():
                num += line[char_idx]
            else:
                break
        middle_num = num
    if sym_idx != len(line) - 1 and line[sym_idx + 1].isdigit() and len(middle_num) < 1:  # Right side
        num = ""
        for char_idx in range(sym_idx + 1, len(line)):
            current_num = line[char_idx]
            if current_num.isdigit():
                num += line[char_idx]
            else:
                break
        right_num = num
    if left_num and middle_num:
        return [int(left_num + middle_num)]
    if left_num or right_num:
        vals = []
        vals.append(int(left_num)) if left_num != "" else 0
        vals.append(int(right_num)) if right_num != "" else 0
        return vals
    if middle_num:
        return [int(middle_num)]
    return []


    

def check_sides_part_two(sym_idx: int, line: str, list_idx: int):
    nums_found = []
    if sym_idx != 0 and line[sym_idx - 1].isdigit():  # Left side
        num = ""
        for char_idx in range(sym_idx - 1, -1, -1):
            current_num = line[char_idx]
            if current_num.isdigit():
                num += line[char_idx]
            else:
                break
        value = int(num[::-1]) if num != "" else 0
        nums_found.append(value)
    if sym_idx != len(line) - 1 and line[sym_idx + 1].isdigit():  # Right side
        num = ""
        for char_idx in range(sym_idx + 1, len(line)):
            current_num = line[char_idx]
            if current_num.isdigit():
                num += line[char_idx]
            else:
                break
        value = int(num) if num != "" else 0
        nums_found.append(value)
    return nums_found

def search_symbols_part_two(engine_line: str, list_idx: int, engine_list: List[str]):
    total = 0
    for char_idx, char in enumerate(engine_line):
        nums_found = []
        if char == "*":
            if list_idx != len(engine_list) - 1:  # Check bottom
                nums_found.extend(check_top_or_bottom_part_two(char_idx, engine_list[list_idx + 1], list_idx + 1))
            if list_idx != 0:  # Check top
                nums_found.extend(check_top_or_bottom_part_two(char_idx, engine_list[list_idx - 1], list_idx - 1))
            nums_found.extend(check_sides_part_two(char_idx, engine_line, list_idx))
        if len(nums_found) == 2:
            total += nums_found[0] * nums_found[1]
    return total


def part_two(engine_data: List[str]):
    total = 0
    for idx, line in enumerate(engine_data):
        total += search_symbols_part_two(line, idx, engine_data)
    print(total)

data_list: List[str] = data_reader("Tavio/challenge3.txt")
# part_one(data_list)
part_two(data_list)
test_data: List[str] = \
["467..114..",
"...*......",
"..35..633.",
"......#...",
"617*......",
".....+.58.",
"..592.....",
"......755.",
"...$.*....",
".664.598.."]
# ["467..114..",
# "...*......",
# "..35..633.",
# "......#...",
# "617*......",
# ".....+.58.",
# "..592.....",
# "......755.",
# "...$.*....",
# ".664.598.."]
# part_one(test_data)
# part_two(test_data)