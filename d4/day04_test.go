package day04

import (
	"strconv"
	"testing"
)

func Test_isValid(t *testing.T) {
	tests := []struct {
		args int
		want bool
	}{
		{111111, false},
		{123444, false},
		{111122, true},
		{122345, true},
		{677789, false},
		{678988, false},
		{777889, true},
		{778889, true},
		{112233, true},
		{212233, false},
		{112333, true},
		{223450, false},
		{123789, false},
		{11111, false},
		{1111111, false},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.args), func(t *testing.T) {
			if got := isValid(tt.args); got != tt.want {
				t.Errorf("%d, %v => %v", tt.args, got, tt.want)
			}
		})
	}
}

func Test_calcNumPasswords(t *testing.T) {

	got := calcNumPasswords(273025, 767253)

	if got <= 0 {
		t.Error("Got 0 or less passwords")
	}

	t.Logf("Num password founds: %d", got)

}
