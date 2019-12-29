package day01

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func calcFuel(input int) int {
	return input/3 - 2
}

// Part 1
func calcInputFuelReqs() (int, error) {

	dat, err := ioutil.ReadFile("input.txt")

	if err != nil {
		return -1, err
	}

	ar := strings.Split(string(dat), "\n")
	sum := 0

	for _, snum := range ar {

		if snum == "" {
			continue
		}

		input, err := strconv.Atoi(snum)
		if err != nil {
			return -1, err
		}
		sum += calcFuel(input)
	}

	return sum, nil
}

// Part2
func calcTotalFuel(input int) int {
	output := 0
	lastFuelStep := calcFuel(input)
	for lastFuelStep > 0 {
		output += lastFuelStep
		lastFuelStep = calcFuel(lastFuelStep)
	}
	return output
}

func calcTotalInputFuelReqs() (int, error) {

	dat, err := ioutil.ReadFile("input.txt")

	if err != nil {
		return -1, err
	}

	ar := strings.Split(string(dat), "\n")
	sum := 0

	for _, snum := range ar {

		if snum == "" {
			continue
		}

		input, err := strconv.Atoi(snum)
		if err != nil {
			return -1, err
		}
		sum += calcTotalFuel(input)
	}

	return sum, nil
}
