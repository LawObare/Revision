package main

import (
	"github.com/01-edu/z01"
)

func main() {
	PrintComb()
}

func PrintComb() {
	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			for k := 0; k <= 9; k++ {
				if !(i == 0 && j == 0 && k == 0) && !(i == 9 && j == 9 && k == 9) {
					z01.PrintRune(rune(i + '0'))
					z01.PrintRune(rune(j + '0'))
					z01.PrintRune(rune(k + '0'))
					if !(i == 9 && j == 9 && k == 8) {
						z01.PrintRune(' ')
						z01.PrintRune(',')
					}
				}
			}
		}
	}
}
