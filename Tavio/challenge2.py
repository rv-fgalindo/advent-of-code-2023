from typing import List
from file_reader import data_reader

def part_one(game_data: List[str]):
    real_games = {"red": 12, "green": 13, "blue": 14}
    total_games = 0
    for game in game_data:
        game = game[5:]
        game_id = ""
        game_passed = True
        for info in game:
            if info == ":":
                break
            game_id += info
        game = game[len(game_id) + 2:]
        rounds = game.split("; ")
        for round in rounds:
            single_round = round.split(", ")
            for data in single_round:
                bruh = tuple(data.split(" "))
                if real_games[bruh[1]] < int(bruh[0]):
                    game_passed = False
        if game_passed:
            total_games += int(game_id)
    print(total_games)


def part_two(game_data: List[str]):
    total_games = 0
    for game in game_data:
        minimum_needed = {"red": 0, "green": 0, "blue": 0}
        game = game[5:]
        game_id = ""
        for info in game:
            if info == ":":
                break
            game_id += info
        game = game[len(game_id) + 2:]
        rounds = game.split("; ")
        for round in rounds:
            single_round = round.split(", ")
            for data in single_round:
                bruh = tuple(data.split(" "))
                color = bruh[1]
                amount = int(bruh[0])
                if minimum_needed[color] < amount:
                    minimum_needed[color] = amount
        total_games += minimum_needed["red"] * minimum_needed["blue"] * minimum_needed["green"]
    print(total_games)

data_list = data_reader("Tavio/challenge2.txt")
part_one(data_list)
# part_two(data_list)

test_data = [
"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"
]
# part_one(test_data)
# part_two(test_data)