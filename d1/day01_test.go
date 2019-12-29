package day01

import (
	"testing"
)

func Test_calcFuel(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := calcFuel(tt.input); got != tt.want {
				t.Errorf("calcFuel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcInputFuelReqs(t *testing.T) {
	sum, err := calcInputFuelReqs()
	if err != nil {
		t.Errorf("calcInputFuelReqs() error = %v", err)
	}
	t.Log(sum)
}

func Test_calcTotalFuel(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{12, 2},
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := calcTotalFuel(tt.input); got != tt.want {
				t.Errorf("calcFuel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcTotalInputFuelReqs(t *testing.T) {

	sum, err := calcInputFuelReqs()
	if err != nil {
		t.Errorf("calcInputFuelReqs() error = %v", err)
	}

	tsum, err := calcTotalInputFuelReqs()
	if err != nil {
		t.Errorf("calcTotalInputFuelReqs() error = %v", err)
	}

	if tsum <= sum {
		t.Errorf("calcTotalInputFuelReqs() total = %d <= sum = %d", tsum, sum)
	}

	t.Log(tsum)
}
