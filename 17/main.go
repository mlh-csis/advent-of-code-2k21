package main

import "fmt"

func main() {
	input, _ := ReadInput("input")
	min_x, max_x, min_y, max_y:= 0, 0, 0, 0
	fmt.Sscanf(input[0], "target area: x=%d..%d, y=%d..%d", &min_x, &max_x, &min_y, &max_y)

	maxes := make(map[[2]int]bool)
	var maxys []int
	
	for vx := 0; vx < 200; vx++ {
		for vy := -300; vy < 300; vy++ {
			step := -1
			y := 0
			x := 0
			dvx := vx
			dvy := vy
			temp_max_y := 0
			for {
				step++
				

				y += dvy
				x += dvx

				if y > temp_max_y {
					temp_max_y = y
				}

				if x > max_x || y < min_y {
					break
				}

				if x >= min_x && x <= max_x && y >= min_y && y <= max_y {
					maxes[[2]int{vx, vy}] = true
					maxys = append(maxys, temp_max_y)
					break
				}

				if dvx > 0 {
					dvx--
				} else if dvx < 0 {
					dvx++
				}
				dvy--
				
			}
		}
	}
	_, max := MinMax(maxys)
	fmt.Println(max)
	fmt.Println(len(maxes))
}

func Part1(input []string) int {
	return 0
}

func Part2(input []string) int {
	return 0
}
