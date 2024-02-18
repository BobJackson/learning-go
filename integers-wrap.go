package main

import "fmt"

func main() {
	var red uint8 = 255
	red++
	fmt.Println(red)

	red = 255
	red += 2
	fmt.Println(red)
	fmt.Printf("%T of red\n ", red)

	red = 0
	red--
	fmt.Printf("red is %v\n", red)

	var number int8 = 127
	number++
	fmt.Println(number)

	number = -128
	number--
	fmt.Println(number)

	var green uint8 = 3
	fmt.Printf("%08b\n", green)
	green++
	fmt.Printf("%08b\n", green)

	var blue uint16 = 65535
	blue++
	fmt.Println(blue)
}
