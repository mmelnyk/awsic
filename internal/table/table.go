package table

import "fmt"

func printWidth(msg string, width int) {
	fmt.Print(msg)
	for i := len(msg); i < width; i++ {
		fmt.Print(" ")
	}
}

func calculateWidth(column []string, data [][]string) []int {
	width := make([]int, len(column))

	for i := 0; i < len(width)-1; i++ {
		width[i] = len(column[i])
	}

	for _, r := range data {
		for i := 0; i < len(width)-1; i++ {
			if len(r[i]) > width[i] {
				width[i] = len(r[i])
			}
		}
	}

	for i := 0; i < len(width)-1; i++ {
		width[i] += 2
	}
	return width
}

func Print(column []string, data [][]string) {
	width := calculateWidth(column, data)

	for i, n := range column {
		printWidth(n, width[i])
	}
	fmt.Println()

	for _, r := range data {
		for i, n := range r {
			printWidth(n, width[i])
		}
		fmt.Println()
	}
}
