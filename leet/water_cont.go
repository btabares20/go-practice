package leet

import "math"

type Coords struct {
	X int
	Y int
}

type Container struct {
	left  Coords
	right Coords
}
func maxArea(height []int) int {
	length := len(height)
	container := Container{
		left:  Coords{0, height[0]},
		right: Coords{length - 1, height[length-1]},
	}

	area := 0

	for container.left.X < container.right.X {
		left_X := container.left.X
		left_Y := container.left.Y
		right_X := container.right.X
		right_Y := container.right.Y

		max_height := math.Min(float64(left_Y), float64(right_Y))
		width := math.Abs(float64(left_X) - float64(right_X))
		eval_area := int(max_height) * int(width)

		if eval_area > area {
			area = eval_area
		}

		if left_Y < right_Y {
			container.left = Coords{left_X + 1, height[left_X+1]}
		} else {
			container.right = Coords{right_X - 1, height[right_X-1]}
		}
	}

	return area
}

