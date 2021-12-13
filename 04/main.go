package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type Board struct {
	Rows    [][]int
	Numbers map[int]bool
}

func BoardFromRows(rows [][]int) Board {
	return Board{Rows: rows, Numbers: make(map[int]bool)}
}

func (b *Board) HasBingo() bool {
	results := make(chan bool, 2)
	var wg sync.WaitGroup
	go func() {
		defer wg.Done()
		results <- b.HasFullRow()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		results <- b.HasFullColumn()
	}()
	wg.Add(1)
	go func() {
		wg.Wait()
		close(results)
	}()
	for res := range results {
		if res {
			return true
		}
	}
	return false
}

func (b *Board) HasFullColumn() bool {
	for col := 0; col < len(b.Rows[0]); col++ {
		has_col := true
		for _, row := range b.Rows {
			if !b.Numbers[row[col]] {
				has_col = false
			}
		}
		if has_col {
			return true
		}
	}
	return false
}

func (b *Board) HasFullRow() bool {
	for _, row := range b.Rows {
		has_row := true
		for _, col := range row {
			if !b.Numbers[col] {
				has_row = false
				break
			}
		}
		if has_row {
			return true
		}
	}
	return false
}

func (b *Board) GetUnmarkedSum() int {
	sum := 0
	for _, row := range b.Rows {
		for _, num := range row {
			if !b.Numbers[num] {
				sum += num
			}
		}
	}
	return sum
}

func (b *Board) NumberDrawn(number int) {
	b.Numbers[number] = true
}

func main() {
	lines, _ := readInput("input")
	var numbers []int
	for _, c := range strings.Split(lines[0], ",") {
		num, _ := strconv.Atoi(c)
		numbers = append(numbers, num)
	}
	var boards []Board
	i := 2
	for i < len(lines) {
		var rows [][]int
		for n := 0; n < 5; n++ {
			var row []int
			for _, c := range strings.Fields(lines[i+n]) {
				number, _ := strconv.Atoi(c)
				row = append(row, number)
			}
			rows = append(rows, row)
		}
		board := BoardFromRows(rows)
		boards = append(boards, board)
		i += 6
	}
	lastScore := 0
	for _, number := range numbers {
		for pos := len(boards) - 1; pos >= 0; pos-- {
			board := boards[pos]
			board.NumberDrawn(number)
			if board.HasBingo() {
				lastScore = number * board.GetUnmarkedSum()
				boards = append(boards[:pos], boards[pos+1:]...)
			}
		}
	}
	fmt.Println(lastScore)
}
