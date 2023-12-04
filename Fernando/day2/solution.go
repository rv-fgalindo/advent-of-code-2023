package main

import (
	"aoc/utilities"
	"fmt"
	"strconv"
	"strings"
)

type Blocks struct {
	Red int
	Green int
	Blue int
}

func main() {
	// pass in true to enable debug (additional print statements)
	Part1(false)
	Part2(false)
}

func Part1(debug bool) {
	fmt.Println("Running Part 1")
	path := "input.txt"
	games, err := utilities.LoadFile(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	// used to track possible games
	sum := 0
	GameLoop:
	for gameNumber, game := range games {
		gameRounds := strings.Split(game, ":")
		if debug {
			fmt.Printf("Game %v\n%v\n", gameNumber, gameRounds)
		}
		rounds := strings.Split(gameRounds[1], ";")
		if debug {
			fmt.Printf("Rounds %v\n", rounds)
		}
		for _, round := range rounds {
			pull := strings.Split(round, ",")
			if debug {
				fmt.Printf("Pull %v\n", pull)
			}
			for _, set := range pull {
				// remove leading and trailing spaces
				set := strings.Trim(set, " ")
				if debug {
					fmt.Printf("Set %v\n", set)
				}
				values := strings.Split(set, " ")
				// changing string to int
				number, err := strconv.Atoi(values[0]); if err != nil {
					fmt.Printf("Error converting %v to int: %v in game %v in round %v ", values[0], err, gameNumber+1, round)
				}
				color := values[1]
				max := GetColorLimit(color)
				if max<12 {
					fmt.Println("Error: color out of range")
				}
				if number > max {
					if debug {fmt.Printf("Game %v is impossible: skipping to next game\n", gameNumber+1)}
					continue GameLoop
				}
			}
		}
		// because first index is 0 need to add 1 to gameNumber
		sum += gameNumber +1
	}
	fmt.Printf("Sum of possible games: %v\n", sum)
}

func GetColorLimit(color string) int {
	switch color {
	case "red":
		return 12
	case "green":
		return 13
	case "blue":
		return 14
	default:
		return 0
	}
}

func Part2(debug bool) {
	fmt.Println("Running Part 2")
	path := "input.txt"
	games, err := utilities.LoadFile(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	// used to total sum of powers
	sum := 0
	for gameNumber, game := range games {
		power :=1
		blockNums := map[string]int{
			"Red":   1,
			"Green": 1,
			"Blue":  1,
		}
		gameRounds := strings.Split(game, ":")
		if debug {
			fmt.Printf("Game %v\n%v\n", gameNumber, gameRounds)
		}
		rounds := strings.Split(gameRounds[1], ";")
		if debug {
			fmt.Printf("Rounds %v\n", rounds)
		}
		// looping over each round
		for _, round := range rounds {
			pull := strings.Split(round, ",")
			if debug {
				fmt.Printf("Pull %v\n", pull)
			}
			for _, set := range pull {
				// remove leading and trailing spaces
				set := strings.Trim(set, " ")
				if debug {
					fmt.Printf("Set %v\n", set)
				}
				values := strings.Split(set, " ")
				// changing string to int
				number, err := strconv.Atoi(values[0]); if err != nil {
					fmt.Printf("Error converting %v to int: %v in game %v in round %v ", values[0], err, gameNumber+1, round)
				}
				color := values[1]
				if blockNums[color] < number {
					blockNums[color] = number
				}
			}
		}
		// because first index is 0 need to add 1 to gameNumber
		for _, value := range blockNums {
			power *= value
		}
		sum += power
	}
	fmt.Printf("Sum of possible games: %v\n", sum)
}