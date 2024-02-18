package piscine

import "fmt"

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i == 0 || i == y-1 || j == 0 || j == x-1 {
				fmt.Print("o")
			} else if i == 1 && (j > 0 && j < x-1) {
				fmt.Print("-")
			} else if i > 1 && i < y-1 && (j == 1 || j == x-2) {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
