package main

import (
	"aoc/utilities"
	"fmt"
	"math"
	"strings"
)

func main() {
	games, err := utilities.LoadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	part1(games, false)
	
	
	// fmt.Printf("Game 1:%v\n", strings.Trim(strings.Split(games[0],": ")[1], " "))

	// Part 1
	// fmt.Println("Part 1: ", part1(input))

	// Part 2
	// fmt.Println("Part 2: ", part2(input))
}

func part1(games []string, debug bool) {
	fmt.Println("Running Part 1")
	points := 0.0
	for _, game := range games {
		matches := 0
		winners := strings.Split(strings.Trim(strings.Split(strings.Trim(strings.Split(game,": ")[1], " "), "|")[0]," "), " ")
		numsHave := strings.Split(strings.Trim(strings.Split(strings.Trim(strings.Split(game,": ")[1], " "), "|")[1]," "), " ")
		if debug{
			fmt.Printf("Winners:%v\n", winners)
			fmt.Printf("Nums Have:%v\nlength %v\n", numsHave, len(numsHave))
		}
		// fmt.Printf("Nums Have:%v\nlength %v\nfinding7:%v\n", numsHave, len(numsHave), numsHave[15])
		// will need to account for some elements just being spaces
		// if encounter will just continue
		winnersMap := make(map[string]bool)
		for _, winner := range winners {
			// want to disregard spaces
			if winner == " " {
				continue
			}
			winnersMap[winner] = true
		}
		for _, num := range numsHave {
			if num == " " {
				continue
			}
			if _, ok := winnersMap[num]; ok {
				if debug{fmt.Printf("Match Found Num:%v\n", num)}
				matches++
			}
		}
		if matches > 0 {
			points += math.Pow(2, float64(matches-1))
		}
	}
	fmt.Printf("Total Points:%v\n", points)
}