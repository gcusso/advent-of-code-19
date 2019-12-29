package day03

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func Test_lineToPoints(t *testing.T) {
	type args struct {
		pos point
		dir point
		mag int
	}
	tests := []struct {
		name       string
		args       args
		wantPoints []point
	}{
		{"0 to R2", args{point{0, 0}, point{1, 0}, 2}, []point{
			point{1, 0},
			point{2, 0},
		}},
		{"0 to U2", args{point{0, 0}, point{0, 1}, 2}, []point{
			point{0, 1},
			point{0, 2},
		}},
		{"0 to D2", args{point{0, 0}, point{0, -1}, 2}, []point{
			point{0, -1},
			point{0, -2},
		}},
		{"0 to L2", args{point{0, 0}, point{-1, 0}, 2}, []point{
			point{-1, 0},
			point{-2, 0},
		}},
		{"10 to R3", args{point{10, 10}, point{1, 0}, 3}, []point{
			point{11, 10},
			point{12, 10},
			point{13, 10},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPoints := lineToPoints(tt.args.pos, tt.args.dir, tt.args.mag); !reflect.DeepEqual(gotPoints, tt.wantPoints) {
				t.Errorf("lineToPoints() = %v, want %v", gotPoints, tt.wantPoints)
			}
		})
	}
}

func Test_calcNextPoints(t *testing.T) {
	type args struct {
		pos  point
		sDir string
	}
	tests := []struct {
		name    string
		args    args
		want    []point
		wantErr bool
	}{
		{"0 to R2", args{point{0, 0}, "R2"}, []point{
			point{1, 0},
			point{2, 0},
		}, false},
		{"0 to U2", args{point{0, 0}, "U2"}, []point{
			point{0, 1},
			point{0, 2},
		}, false},
		{"0 to D2", args{point{0, 0}, "D2"}, []point{
			point{0, -1},
			point{0, -2},
		}, false},
		{"0 to L2", args{point{0, 0}, "L2"}, []point{
			point{-1, 0},
			point{-2, 0},
		}, false},
		{"10 to U2", args{point{10, 10}, "U2"}, []point{
			point{10, 11},
			point{10, 12},
		}, false},
		{"Dir Error", args{point{0, 0}, "X212"}, nil, true},
		{"Parse Error", args{point{0, 0}, "D3f2"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calcNextPoints(tt.args.pos, tt.args.sDir)
			if (err != nil) != tt.wantErr {
				t.Errorf("calcNextPoints() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcNextPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculatePath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name       string
		args       args
		wantPoints []point
	}{
		{"R2,U2,L3", args{"R2,U2,L3"}, []point{
			point{1, 0},
			point{2, 0},
			point{2, 1},
			point{2, 2},
			point{1, 2},
			point{0, 2},
			point{-1, 2},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPoints := calculatePath(tt.args.path); !reflect.DeepEqual(gotPoints, tt.wantPoints) {
				t.Errorf("calculatePath() = %v, want %v", gotPoints, tt.wantPoints)
			}
		})
	}
}

// part one
func Test_manhattanDistance(t *testing.T) {
	tests := []struct {
		input        string
		wantDistance int
	}{
		{"R8,U5,L5,D3\nU7,R6,D4,L4", 6},
		{"R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83", 159},
		{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7", 135},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if gotDistance := manhattanDistance(tt.input); gotDistance != tt.wantDistance {
				t.Errorf("manhattanDistance() = %v, want %v", gotDistance, tt.wantDistance)
			}
		})
	}
}

func Test_manhattanDistanceInput(t *testing.T) {
	dat, err := ioutil.ReadFile("input.txt")

	if err != nil {
		t.Fatal(err)
	}

	result := manhattanDistance(string(dat))

	t.Logf("Result: %d", result)

	if result == 0 {
		t.Errorf("Result was 0")
	}
}

// part two
func Test_stepDistance(t *testing.T) {
	tests := []struct {
		input     string
		wantSteps int
	}{
		{"R8,U5,L5,D3\nU7,R6,D4,L4", 30},
		{"R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83", 610},
		{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7", 410},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if gotSteps := stepDistance(tt.input); gotSteps != tt.wantSteps {
				t.Errorf("stepDistance() = %v, want %v", gotSteps, tt.wantSteps)
			}
		})
	}
}

func Test_stepDistanceInput(t *testing.T) {
	dat, err := ioutil.ReadFile("input.txt")

	if err != nil {
		t.Fatal(err)
	}

	result := stepDistance(string(dat))

	t.Logf("Result: %d", result)

	if result == 0 {
		t.Errorf("Result was 0")
	}
}
