from file_reader import data_reader

def calculate_long_sum(val: int):
    total = 0
    for x in range(1, val + 1):
        total += x
    return total


def part_one(non_form_data):
    x = non_form_data[0].split(",")
    data = [int(item) for item in x]
    num_list = set()
    for item in data:
        num_list.add(item)
    lowest = float('inf')
    for num in num_list:
        total = 0
        for item in data:
            total += abs(item - num)
        if total < lowest:
            lowest = total
    print(lowest)

def part_two(non_form_data):
    x = non_form_data[0].split(",")
    data = [int(item) for item in x]
    highest = max(data)
    lowest_num = 0
    lowest = float('inf')
    for num in range(highest + 1):
        total = 0
        for item in data:
            total += calculate_long_sum(abs(item - num))
            if total > lowest:
                break
        if total < lowest:
            lowest = total
            lowest_num = num
    print(lowest, lowest_num)

data_list = data_reader("Tavio/input2.txt")
test_data = ["16,1,2,0,4,2,7,1,2,14"]
# part_one(test_data)

part_two(data_list)