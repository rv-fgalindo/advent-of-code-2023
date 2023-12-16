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
	// part1(games, false)
	part2(games, false)
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
			// spaces werent the issue; it was the empty strings
			if winner == " " || winner == "" {
				continue
			}
			winnersMap[winner] = false
		}
		for _, num := range numsHave {
			if num == " " || num == "" {
				continue
			}
			if _, ok := winnersMap[num]; ok {
				if debug{fmt.Printf("Match Found Num:%v\n", num)}
				winnersMap[num] = true
				matches++
			}
		}
		if matches > 0 {
			points += math.Pow(2, float64(matches-1))
		}
	}
	fmt.Printf("Total Points:%v\n", points)
}

func part2(games []string, debug bool) {
	fmt.Println("Running Part 2")
	points := 0
	gameTracker := make(map[int]int)
	for i := range games {
		// add i to gameTracker map with value of 1
		gameTracker[i] = 1
	}
	for i, game := range games {
		matches := 0
		winners := strings.Split(strings.Trim(strings.Split(strings.Trim(strings.Split(game,": ")[1], " "), "|")[0]," "), " ")
		numsHave := strings.Split(strings.Trim(strings.Split(strings.Trim(strings.Split(game,": ")[1], " "), "|")[1]," "), " ")
		if debug{
			fmt.Printf("Winners:%v\n", winners)
			fmt.Printf("Nums Have:%v\nlength %v\n", numsHave, len(numsHave))
		}
		
		winnersMap := make(map[string]bool)
		for _, winner := range winners {
			// want to disregard spaces
			// spaces werent the issue; it was the empty strings
			if winner == " " || winner == "" {
				continue
			}
			winnersMap[winner] = false
		}
		for _, num := range numsHave {
			if num == " " || num == "" {
				continue
			}
			if _, ok := winnersMap[num]; ok {
				if debug{fmt.Printf("Match Found Num:%v\n", num)}
				winnersMap[num] = true
				matches++
			}
		}
		if matches > 0 {
			multiplyCards(i, matches, gameTracker)
		}
	}
	for _, v := range gameTracker {
		points += v
	}
	fmt.Printf("Total Points:%v\n", points)
}

func multiplyCards(currentWinner int, cardsWon int, gameTracker map[int]int) {
	
	multiplier := gameTracker[currentWinner]
	for i := 0; i < cardsWon; i++ {
		cardWon := currentWinner + i+1
		gameTracker[cardWon]= gameTracker[cardWon] + (1 * multiplier)
	}
	
}