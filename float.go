package main

import (
	"fmt"
	"math"
)

func main() {
	third := 1.0 / 3.0
	fmt.Println(third + third + third)

	piggyBank := 0.1
	piggyBank += 0.2
	fmt.Println(piggyBank)

	fmt.Println(math.Abs(piggyBank-0.3) < 0.0001)

	piggyBankCopy := 0.0
	for i := 0; i < 11; i++ {
		piggyBankCopy += 0.1
	}
	fmt.Println(piggyBankCopy)
}
