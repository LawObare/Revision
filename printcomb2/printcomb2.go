package main

import (
	"github.com/01-edu/z01"
)

func main() {
	PrintComb2()
}

func PrintComb2() {
	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			for k := 0; k <= 9; k++ {
				for l := 0; l <= 9; l++ {
					z01.PrintRune(rune(i + '0'))
					z01.PrintRune(rune(j + '0'))
					z01.PrintRune(' ')
					z01.PrintRune(rune(k + '0'))
					z01.PrintRune(rune(l + '0'))
					if !(i == 9 && j == 9 && k == 9 && l == 9) {
						z01.PrintRune(',')
					}
				}
			}
		}
	}
}
