package main

import (
	"fmt"
	"strings"
)


func main() {
	input, _ := ReadInput("input")
	template := input[0]
	rules := make(map[string]string, len(input)-2)
	for i := 2; i < len(input); i++ {
		parts := strings.Split(input[i], " -> ")
		rules[parts[0]] = parts[1]
	}
	part1 := Solve(template, rules, 10)
	fmt.Printf("Part 1: %d\n", part1)
	part2 := Solve(template, rules, 40)
	fmt.Printf("Part 2: %d\n", part2)
}

func Solve(template string, rules map[string]string, iterations int) int {
	pairs := make(map[string]int)

	for i := 0; i+1 < len(template); i++ {
		pairs[template[i:i+2]]++
	}

	for steps := 0; steps < iterations; steps++ {
		new_pair := make(map[string]int)
		for pair, how_many := range pairs {
			if insertion, ok := rules[pair]; ok {
				new_pair[string(pair[0]) + insertion] += how_many
				new_pair[insertion + string(pair[1])] += how_many
			}
		}
		pairs = new_pair
	}

	counts := make(map[string]int, 10)

	for pair, cnt := range pairs {
		counts[string(pair[0])] += cnt
	}
	counts[string(template[len(template)-1])]++

	var cnts []int
	for _, v:= range counts {
		cnts = append(cnts, v)
	}
	min, max := MinMax(cnts)
	return max-min
}
