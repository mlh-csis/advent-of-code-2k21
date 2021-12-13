package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	fish_str := strings.Split(line, ",")
	var fish []int
	for _, f := range fish_str {
		f_n, err := strconv.Atoi(f)
		if err != nil {
			return nil, err
		}
		fish = append(fish, f_n)
	}
	return fish, scanner.Err()
}

func calculatePopulationAt(fish []int, days int) int64 {
	todaysFish := [9]int64{}

	// count existing fish by days remaining to reproduce
	for f := 0; f < len(fish); f++ {
		todaysFish[fish[f]]++
	}
	fmt.Println(todaysFish)

	for day := 0; day < days; day++ {
		reproduction := todaysFish[0]
		for i := 0; i < 8; i++ {
			todaysFish[i] = todaysFish[i+1]
		}
		todaysFish[6] += reproduction
		todaysFish[8] = reproduction
	}
	var sum int64
	for i := 0; i < 9; i++ {
		sum += todaysFish[i]
	}
	return sum
}
