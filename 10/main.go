package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, _ := ReadInput("input")
	corrupted, incomplete := BothParts(input)
	fmt.Printf("Part 1: %d\n", corrupted)
	fmt.Printf("Part 2: %d\n", incomplete)
}

func BothParts(input []string) (int, int) {
	corrupted, incomplete := 0, 0
	for _, line := range input {
		corr, incompl := checkLine(line)
		corrupted += corr
		incomplete += incompl
	}
	return corrupted, incomplete
}

func ReadInput(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func checkLine(line string) (int, int) {
	scores := map[rune]int{
		'{': 1197,
		'(': 3,
		'<': 25137,
		'[': 57,
	}

	leftover_scores := map[rune]int{
		'{': 3,
		'(': 1,
		'<': 4,
		'[': 2,
	}
	d := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
		'>': '<',
	}
	s := ""
	for _, c := range line {
		switch c {
		case '{', '(', '<', '[':
			s += string(c)
		default:
			remove := string(d[c])
			if string(s[len(s)-1]) != remove {
				return scores[d[c]], 0
			}
			s = s[0 : len(s)-1]
		}
	}
	sum := 0
	for _, c := range Reverse(s) {
		sum = 5*sum + leftover_scores[c]
	}
	return 0, sum
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
