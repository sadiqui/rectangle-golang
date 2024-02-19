package piscine

import "fmt"

func QuadE(x, y int) {
	if x <= 0 || y <= 0 {
		return // exit
	}

	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if (row == 1 && col == 1) || (row == y && col == x && row != 1 && col != 1) {
				fmt.Print("A")
			} else if (row == 1 && col == x) || (row == y && col == 1) {
				fmt.Print("C")
			} else if row == 1 || row == y || col == 1 || col == x {
				fmt.Print("B")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// ABBBC
// B   B
// CBBBA
// ABBBC
// A
// A
// B
// B
// B
// C
