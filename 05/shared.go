package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Vent struct {
	Start Point
	End   Point
}

func (v *Vent) GetLinePoints() ([]Point, error) {
	var points []Point

	if Abs(v.Start.X-v.End.X) == Abs(v.Start.Y-v.End.Y) || v.Start.Y == v.End.Y || v.Start.X == v.End.X {
		var xs []int
		var ys []int

		if v.Start.X < v.End.X {
			for i := v.Start.X; i <= v.End.X; i++ {
				xs = append(xs, i)
			}
		} else {
			for i := v.Start.X; i >= v.End.X; i-- {
				xs = append(xs, i)
			}
		}

		if v.Start.Y < v.End.Y {
			for i := v.Start.Y; i <= v.End.Y; i++ {
				ys = append(ys, i)
			}
		} else {
			for i := v.Start.Y; i >= v.End.Y; i-- {
				ys = append(ys, i)
			}
		}
		max := len(xs)
		if max == 1 {
			max = len(ys)
		}
		for i := 0; i < max; i++ {
			x := xs[0]
			if len(xs) > 1 {
				x = xs[i]
			}
			y := ys[0]
			if len(ys) > 1 {
				y = ys[i]
			}
			points = append(points, Point{X: x, Y: y})
		}

		return points, nil
	}

	return points, errors.New("not a considered line")
}



func readInput(path string) ([]Vent, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var vents []Vent
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		points := strings.Split(line, " -> ")

		start_str := strings.Split(points[0], ",")
		end_str := strings.Split(points[1], ",")

		var start [2]int
		var end [2]int

		for pos, s := range start_str {
			n, _ := strconv.Atoi(s)
			start[pos] = n
		}

		for pos, s := range end_str {
			n, _ := strconv.Atoi(s)
			end[pos] = n
		}

		start_p := Point{X: start[0], Y: start[1]}
		end_p := Point{X: end[0], Y: end[1]}
		vents = append(vents, Vent{Start: start_p, End: end_p})
	}

	return vents, scanner.Err()
}
