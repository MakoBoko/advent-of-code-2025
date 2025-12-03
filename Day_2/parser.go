package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

type IDRange struct {
	Lower int
	Upper int
}

func parseFile(filePath string) ([]IDRange, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var idRanges []IDRange
	for scanner.Scan() {
		ranges, err := parseLine(scanner.Text())

		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		idRanges = ranges
	}

	if err := scanner.Err(); err != nil {
		log.Fatal()
		return nil, err
	}

	return idRanges, nil

}

func parseLine(idRangesRaw string) ([]IDRange, error) {
	rangesStr := strings.Split(idRangesRaw, ",")

	var idRanges []IDRange
	for _, rangeStr := range rangesStr {
		splitRange := strings.Split(rangeStr, "-")
		if len(splitRange) < 2 {
			return nil, errors.New("error while parsing id ranges, invalid range: " + rangeStr)
		}

		lowerRange, errLw := strconv.Atoi(splitRange[0])
		upperRange, errUp := strconv.Atoi(splitRange[1])

		if errUp != nil || errLw != nil {
			return nil, errors.New("error while parsing id ranges, invalid number in range: " + rangeStr)
		}

		idRanges = append(idRanges, IDRange{Lower: lowerRange, Upper: upperRange})
	}

	return idRanges, nil
}
