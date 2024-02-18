package piscine

import "fmt"

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if (row == 1 || row == y) && (col == 1 || col == x) {
				fmt.Print("o")
			} else if row == 1 || row == y {
				fmt.Print("-")
			} else if col == 1 || col == x {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
