package main

import "fmt"

func main() {
	vents, _ := readInput("input")
	seafloor := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		seafloor[i] = make([]int, 1000)
		for j := 0; j < 1000; j++ {
			seafloor[i][j] = 0
		}
	}

	for _, vent := range vents {
		points, error := vent.GetLinePoints()
		if error == nil {
			for _, point := range points {
				seafloor[point.Y][point.X] += 1
			}
		}
	}
	count := 0
	for _, line := range seafloor {
		for _, point := range line {
			if point > 1 {
				count += 1
			}
		}
	}
	fmt.Println(count)
}
