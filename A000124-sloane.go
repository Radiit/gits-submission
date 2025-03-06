package main

import (
	"fmt"
	"strings"
)

func generateSloane(n int) string {
	sequence := []int{}

	currentNum := 1
	sequence = append(sequence, currentNum)

	for i := 1; i < n; i++ {
		currentNum += i
		sequence = append(sequence, currentNum)
	}

	var result []string
	for _, num := range sequence {
		result = append(result, fmt.Sprintf("%d", num))
	}

	return strings.Join(result, "-")
}

func generateSloaneRecursive(n int, currentNum int, sequence []int) []int {
	if len(sequence) == n {
		return sequence
	}

	sequence = append(sequence, currentNum)

	return generateSloaneRecursive(n, currentNum+len(sequence), sequence)
}

func main() {
	var n int
	fmt.Println("Masukkan input n:")
	fmt.Scan(&n)

	iterativeFuncOutput := generateSloane(n)
	fmt.Println("output iterative: ", iterativeFuncOutput)

	recursiveSeq := generateSloaneRecursive(n, 1, []int{})
	var recursiveResult []string

	for _, num := range recursiveSeq {
		recursiveResult = append(recursiveResult, fmt.Sprintf("%d", num))
	}
	fmt.Println("output rekursif:", strings.Join(recursiveResult, "-"))
}
