// My weight loss program
package main

import "fmt"

// main is the function where it all begins.
func main() {
	fmt.Print("My weight on the surface of Mars is ")
	fmt.Print(160 * 0.3783)
	fmt.Print(" lbs, and I would be ")
	fmt.Print(32 * 365 / 687)
	fmt.Print(" years old.")
	fmt.Println()

	fmt.Println("My weight on the surface of Mars is", 160*0.3783, "lbs, and I would be", 32*365/687, "years old.")

	fmt.Printf("My weight on the surface of Mars is %v lbs,", 160*0.3783)
	fmt.Printf(" and I would be %v years old.\n", 32*365/687)
	fmt.Printf("My weight on the surface of %v is %v lbs.\n", "Earth", 160)

	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)

}
