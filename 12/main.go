package main

import (
	"fmt"
	"strings"
)

func main() {
	input, _ := ReadInput("input")
	parsed := ParseInput(input)
	part1 := Navigate(parsed, can_visit)
	fmt.Printf("Part 1: %d\n", part1)
	part2 := Navigate(parsed, can_visit2)
	fmt.Printf("Part 2: %d\n", part2)
}

func ParseInput(input []string) map[string]map[string]struct{}{
	output := make(map[string]map[string]struct{})

	for _, line := range input {
		nodes := strings.Split(line, "-")
		if output[nodes[0]] == nil {
			output[nodes[0]] = make(map[string]struct{})
		} 
		if output[nodes[1]] == nil {
			output[nodes[1]] = make(map[string]struct{})
		} 
		output[nodes[0]][nodes[1]] = struct{}{}
		output[nodes[1]][nodes[0]] = struct{}{}
	}

	return output
}

func Navigate(input map[string]map[string]struct{}, MayVisit func(string, string) bool ) int {
	paths := make(map[string]struct{})

	var visit func(current string, prev string, path string)
    visit = func (current string, prev string, path string) {

		if !MayVisit(path, current){
			return
		}
		
		path += fmt.Sprintf(",%s", current)
		if current == "end" {
			paths[path] = struct{}{}
			return
		}
		
		for cave := range input[current] {
			visit(cave, current, path)
		}
		visit(prev, current, path)
	}
	for cave := range input["start"] {
		visit(cave, "start", "start")
	}
	
	return len(paths)
}

func can_visit(path string, current string) bool {
	if current == "start" {
		return false
	}
	prev := strings.Split(path, ",")
	lower := strings.ToLower(current)

	for _, cave := range prev {
		if lower == current && current == cave {
			return false
		}
	}
	return true
}

func can_visit2(path string, current string) bool {
	if current == "start" {
		return false
	}

	lower := strings.ToLower(current)

	if lower == current{
		twos := 0
		seen := make(map[string]int)
		seen[current]=1
		for _, s := range strings.Split(path, ",") {
			if strings.ToLower(s) == s {
				seen[s]++
				if seen[s] > 1 {
					twos++
				}

				if twos == 2 {
					return false
				}
			}
		}
	}
	return true
}