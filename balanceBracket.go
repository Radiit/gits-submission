package main

import (
	"bufio"
	"fmt"
	"os"
)

func balanceBracket(n string) string {
	stack := []string{}
	for i := 0; i < len(n); i++ {
		current := string(n[i])

		if current == " " || current == "\t" || current == "\n" {
			continue
		}

		if current == "(" || current == "{" || current == "[" {
			stack = append(stack, current)
		} else if current == ")" || current == "}" || current == "]" {

			if len(stack) == 0 {
				return "NO"
			}

			last := stack[len(stack)-1]

			stack = stack[:len(stack)-1]

			if current == ")" && last != "(" {
				return "NO"
			}
			if current == "}" && last != "{" {
				return "NO"
			}
			if current == "]" && last != "[" {
				return "NO"
			}
		}
	}

	if len(stack) == 0 {
		return "YES"
	}

	return "NO"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Input: ")
	scanner.Scan()
	input := scanner.Text()
	result := balanceBracket(input)
	fmt.Println("Output:", result)
}
