package main

import "fmt"

func main() {
	in := [][]string{{"1", "0", "1", "0", "0"}, {"1", "0", "1", "1", "1"}, {"1", "1", "1", "1", "1"}, {"1", "0", "0", "1", "0"}}
	out := maximalRectangle(in)
	fmt.Println(out)
}
func maximalRectangle(matrix [][]string) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	m := len(matrix)
	n := len(matrix[0])

	// use the column length
	left := make([]int, n)
	right := make([]int, n)
	height := make([]int, n)

	maxArea := 0
	var curr_left int
	var curr_right int

	for i := range right {
		right[i] = n - 1
	}

	fmt.Println(right)

	for row := 0; row < m; row++ {
		curr_left = 0
		curr_right = n - 1

		// update height
		for j := 0; j < n; j++ {
			if matrix[row][j] == "1" {
				height[j]++
			} else {
				height[j] = 0
			}
		}

		// fmt.Println(height)

		// update left
		for j := 0; j < n; j++ {
			if matrix[row][j] == "1" {
				left[j] = max(left[j], curr_left)
			} else {
				left[j] = 0
				curr_left = j + 1
			}
		}

		// update right
		for j := n - 1; j >= 0; j-- {
			if matrix[row][j] == "1" {
				right[j] = min(right[j], curr_right)
			} else {
				right[j] = n - 1
				curr_right = j - 1
			}
		}

		fmt.Println("****")
		fmt.Printf("height: %v \n", height)
		fmt.Printf("right: %v \n", right)
		fmt.Printf("left: %v \n", left)

		// area
		for j := 0; j < n; j++ {
			current := (right[j] - left[j] + 1) * height[j]
			maxArea = max(maxArea, current)
			fmt.Printf("current:max %v:%v \n", current, maxArea)

		}
	}
	return maxArea
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
