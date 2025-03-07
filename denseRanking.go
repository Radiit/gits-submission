package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func denseRanking(player int, playerScore []int, gitsAttempt int, gitsScore []int) []int {
	uniqueScores := make([]int, len(playerScore))
	copy(uniqueScores, playerScore)

	sort.Slice(uniqueScores, func(i, j int) bool {
		return uniqueScores[i] > uniqueScores[j]
	})

	j := 0
	for i := 1; i < len(uniqueScores); i++ {
		if uniqueScores[i] != uniqueScores[j] {
			j++
			uniqueScores[j] = uniqueScores[i]
		}
	}
	uniqueScores = uniqueScores[:j+1]

	gitsRanks := make([]int, gitsAttempt)
	for i, score := range gitsScore {
		rank := 1
		for _, uniqueScore := range uniqueScores {
			if score >= uniqueScore {
				break
			}
			rank++
		}
		gitsRanks[i] = rank
	}

	return gitsRanks
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter the number of players:")
	scanner.Scan()
	player, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter player scores (space separated):")
	scanner.Scan()
	scoreStr := strings.Fields(scanner.Text())
	playerScore := make([]int, player)

	for i := 0; i < player && i < len(scoreStr); i++ {
		playerScore[i], _ = strconv.Atoi(scoreStr[i])
	}

	fmt.Println("Enter the number of GITS attempts:")
	scanner.Scan()
	gitsAttempt, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Enter GITS scores (space separated):")
	scanner.Scan()
	gitsScoreStr := strings.Fields(scanner.Text())
	gitsScore := make([]int, gitsAttempt)

	for i := 0; i < gitsAttempt && i < len(gitsScoreStr); i++ {
		gitsScore[i], _ = strconv.Atoi(gitsScoreStr[i])
	}

	result := denseRanking(player, playerScore, gitsAttempt, gitsScore)

	for i, rank := range result {
		fmt.Print(rank)
		if i < len(result)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
