from typing import List
from file_reader import data_reader

def part_one(datas: List[str]):
    total = 0
    for data in datas:
        single_card = data[8:]
        split_card = single_card.split("|")
        winning_card = []
        guessing_card = []
        for idx, card in enumerate(split_card):
            card = card.split(" ")
            for value in card:
                if value.isdigit() and idx == 0:
                    guessing_card.append(value)
                elif value.isdigit():
                    winning_card.append(value)
        points = 0
        for win in winning_card:
            for guess in guessing_card:
                if guess == win:
                    if points == 0:
                        points += 1
                    else:
                        points += points
        total += points
    print(total)


def part_two(datas: List[str]):
    total = 0
    card_dict = {idx + 1: 1 for idx, _ in enumerate(datas)}
    for idx, data in enumerate(datas):
        card_idx = idx + 1
        winning_nums = 0
        single_card = data[8:]
        split_card = single_card.split("|")
        winning_card = []
        guessing_card = []
        for idx, card in enumerate(split_card):
            card = card.split(" ")
            for value in card:
                if value.isdigit() and idx == 0:
                    guessing_card.append(value)
                elif value.isdigit():
                    winning_card.append(value)
        for win in winning_card:
            for guess in guessing_card:
                if guess == win:
                    winning_nums += 1
        if winning_nums:
            multiplier = 1 * (card_dict[card_idx])
            for _ in range(winning_nums):
                card_dict[card_idx + 1] += multiplier
                card_idx += 1
    for num in card_dict.values():
        total += num
    print(total)


data_list: List[str] = data_reader("Tavio/challenge4.txt")
test_data: List[str] = \
["Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"]
# part_one(test_data)
part_two(data_list)