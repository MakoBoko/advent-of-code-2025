package main

import (
	"fmt"
)

func main() {
	instructions, err := parseFile()

	if err != nil {
		return
	}
	password := computePassword(instructions)

	fmt.Println("The password is: ", password)
}

func computePassword(instructions []int) int {
	password := 0
	dial := 50
	max_dial := 100 //99+1

	for _, instruction := range instructions {

		dial += instruction

		dial = (dial%max_dial + max_dial) % max_dial

		if dial == 0 {
			password++
		}
	}

	return password
}
