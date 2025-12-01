package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
)

func parseFile() ([]int, error) {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var safeInstrictions []int

	for scanner.Scan() {
		instruction, err := parseLine(scanner.Text())

		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		safeInstrictions = append(safeInstrictions, instruction)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal()
		return nil, err
	}

	return safeInstrictions, nil

}

func parseLine(instructionText string) (int, error) {
	if len(instructionText) <= 1 {
		return 0, errors.New("invalid instruction")
	}

	moveAmount, err := strconv.Atoi(instructionText[1:])

	if err != nil {
		return 0, errors.New("invalid instruction: move amount must be a number")
	}

	if instructionText[0] == 'L' {
		moveAmount = -moveAmount
	} else if instructionText[0] != 'R' {
		return 0, errors.New("invalid instruction: must be L or R")
	}

	return moveAmount, nil
}
