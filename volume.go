// Copyright (c) 2021 Daniel Pawelko All rights reserved
//
// Created by: Daniel Pawelko
// Created on: May 2021
// This file contains GO program that volume of a sphere

package main

import (
	"fmt"
	"math"
)

// This main function will ask user for dimensions and output calculations
func main() {
	// Defining variables
	var radius float64

	fmt.Println("Volume of Sphere")
	fmt.Println("V = 4/3πr³")
	fmt.Println()

	// User Input
	fmt.Println("Please enter Radius:")
	fmt.Println("Radius: ")
	fmt.Scanln(&radius)
	fmt.Println()

	// calculations
	var volume float64 = (4.0 / 3.0 * math.Pi * math.Pow(radius, 3))

	// Print out volume
	fmt.Println("Volume is: ", math.Round(volume*100)/100)
}
