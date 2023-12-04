from copy import deepcopy
from typing import Dict, List
from file_reader import data_reader

def part_one(data: List[str]):

    original_dict: Dict[int, List[int]] = \
    {0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}
    data = data[0].split(",")
    new_data = [int(item) for item in data]
    for item in new_data:
        if item in original_dict:
            original_dict[item] += 1
    for day in range(256):
        copied_dict = deepcopy(original_dict)
        increased = 0
        for key, value in original_dict.items():
            if key == 0:
                increased = value
            else:
                original_dict[key - 1] = copied_dict[key]
        original_dict[6] += increased
        original_dict[8] = increased
    all_vals = [value for value in original_dict.values()]
    print(sum(all_vals))



data_list = data_reader("Tavio/input.txt")
part_one(data_list)
# test_data = ["3,4,3,1,2"]
# part_one(test_data)