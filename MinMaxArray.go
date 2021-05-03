// This program generates 10 random numbers and finds the min/max values.
// Created By: Marlon Poddalgoda
// Created on: 2021-05-03

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func maxValue(array [10]int) int {
	// this function finds the max value in an array

	var maxValue int = 0
	var counter int = 0

	for counter < len(array) {
		// loop goes through each value
		if array[counter] > maxValue {
			maxValue = array[counter]
		} else {
			maxValue += 0
		}
		// add 1 to counter
		counter++
	}
	return (maxValue)
}

func minValue(array [10]int) int {
	// this function finds the min value in an array

	var minValue int = 99
	var counter int = 0

	for counter < len(array) {
		// loop goes through each value
		if array[counter] < minValue {
			minValue = array[counter]
		} else {
			minValue += 0
		}
		// add 1 to counter
		counter++
	}
	return (minValue)
}

func main() {
	// this function takes in user input

	fmt.Println("This program generates 10 random numbers and locates the min/max value")
	fmt.Println("")

	// constants
	const arrayLength int = 10
	const randomRange int = 99

	// counter variable
	var counter int = 0
	var highestValue int
	var lowestValue int

	// creating array
	var randomArray [arrayLength]int

	// generates random number
	rand1 := rand.NewSource(time.Now().UnixNano())
	random := rand.New(rand1)

	// for loop to place random numbers in array
	for counter < arrayLength {
		// generating random numbers
		randomArray[counter] = random.Intn(randomRange)
		// print random value
		fmt.Println(randomArray[counter])
		// increase counter
		counter++
	}

	// call max function
	highestValue = maxValue(randomArray)
	// call min function
	lowestValue = minValue(randomArray)

	fmt.Println("")
	fmt.Println(highestValue)
	fmt.Println(lowestValue)

	// program closes
	fmt.Println("")
	fmt.Println("Done.")
}
