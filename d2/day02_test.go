package day02

import (
	"reflect"
	"testing"
)

func Test_opcodes(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
		{[]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
		{[]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
		{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := intCodes(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intCodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gravityAssistProgram(t *testing.T) {
	got, err := gravityAssistProgram()
	if err != nil {
		t.Errorf("gravityAssistProgram() error = %v", err)
		return
	}
	t.Logf("Gravity Assist result: %v", got[0])
}

func Test_findGravityAssistProgram(t *testing.T) {
	r, err := findGravityAssistProgram(19690720)
	if err != nil {
		t.Errorf("findGravityAssistProgram() error = %v", err)
		return
	}
	t.Logf("Find Gravity Assist result: %d", r)
}
