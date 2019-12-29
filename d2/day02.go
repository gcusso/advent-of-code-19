package day02

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func intCodes(input []int) []int {

	count := len(input)
	r := make([]int, count)
	copy(r, input)

	for i := 0; i < count; i += 4 {
		op := r[i]
		if op != 1 && op != 2 {
			return r
		}
		s1 := r[r[i+1]%count]
		s2 := r[r[i+2]%count]
		offIndex := r[i+3] % count
		if op == 1 {
			r[offIndex] = s1 + s2
		} else {
			r[offIndex] = s1 * s2
		}
	}

	return r
}

func gravityAssistProgram() ([]int, error) {

	dat, err := ioutil.ReadFile("input.txt")

	if err != nil {
		return nil, err
	}

	sDat := string(dat)
	sDat = strings.Replace(sDat, "\n", "", -1)
	sInput := strings.Split(sDat, ",")

	inp := make([]int, len(sInput))

	for i, s := range sInput {
		inp[i], err = strconv.Atoi(s)

		if err != nil {
			return nil, err
		}
	}

	inp[1] = 12
	inp[2] = 2

	r := intCodes(inp)

	return r, nil
}

func findGravityAssistProgram(target int) (int, error) {

	dat, err := ioutil.ReadFile("input.txt")

	if err != nil {
		return -1, err
	}

	sDat := string(dat)
	sDat = strings.Replace(sDat, "\n", "", -1)
	sInput := strings.Split(sDat, ",")

	inp := make([]int, len(sInput))

	for i, s := range sInput {
		inp[i], err = strconv.Atoi(s)

		if err != nil {
			return -1, err
		}
	}
	tInp := make([]int, len(inp))

	for n := 0; n < 100; n++ {
		for v := 0; v < 100; v++ {
			copy(tInp, inp)
			inp[1], inp[2] = n, v
			r := intCodes(inp)
			if r[0] == target {
				return 100*n + v, nil
			}
		}
	}

	return -1, nil
}
