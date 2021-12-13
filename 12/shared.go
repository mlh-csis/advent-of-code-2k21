package main

import (
	"bufio"
	"os"
)

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

func sum(input []int) int {
	sum := 0
	for i := range input {
		sum += input[i]
	}
	return sum}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func intersection(s1 []string, s2 []string) []string {
	s_intersection := make(map[string]bool)
	for _, s := range s1 {
		s_intersection[s] = true
	}

	for _, s := range s2 {
		if _, ok := s_intersection[s]; ok {
			s_intersection[s] = false
		}
	}

	var ret []string
	for k, val := range s_intersection {
		if val {
			ret = append(ret, k)
		}
	}
	return ret
}

func union(s1 []string, s2 []string) []string {
	s_union := make(map[string]bool)
	for _, s := range s1 {
		s_union[s] = true
	}
	for _, s := range s2 {
		s_union[s] = true
	}

	var ret []string
	for k := range s_union {
		ret = append(ret, k)
	}
	return ret
}
