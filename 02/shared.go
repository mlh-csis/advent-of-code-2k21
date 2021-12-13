package main

import (
	"bufio"
	"os"
	"strings"
)

func readInput(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		lines = append(lines, fields)
	}
	return lines, scanner.Err()
}
