// Created by: Charlotte Jhu
// Created on: March 2023
//
// This program calculates the area of a triangle

package main

import "fmt"

func main() {
	// This function calculates the area of a triangle
	// variables
	var base int
	var height int

	// input
	fmt.Println("This program calculates the area of a triangle.")
	fmt.Println()
	fmt.Print("Enter the base (cm): ")
	fmt.Scanln(&base)
	fmt.Print("Enter the height (cm): ")
	fmt.Scanln(&height)

	// process
	area := (base * height) / 2

	// output
	fmt.Println()
	fmt.Println("The area is", area, "cm².")
	fmt.Println("\nDone.")
}
