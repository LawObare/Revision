package main

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	var isNegative bool
	if n < 0 {
		n = -n
		isNegative = true
	}
	if isNegative {
		z01.PrintRune('-')
	}
	if n >= 10 {
		PrintNbr(n / 10)
	}
	z01.PrintRune(rune('0' + n%10))

}

func main() {
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
	z01.PrintRune('\n')
}
