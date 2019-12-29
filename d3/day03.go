package day03

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type point struct {
	X, Y int
}

func (s point) String() string {
	return fmt.Sprintf("P(%d, %d)", s.X, s.Y)
}

func (s point) Add(v point) point {
	return point{s.X + v.X, s.Y + v.Y}
}

func (s point) Mul(v int) point {
	return point{s.X * v, s.Y * v}
}

func (s point) AddV(v int) point {
	return point{s.X, s.Y + v}
}

func (s point) AddH(v int) point {
	return point{s.X + v, s.Y}
}

func (s point) Sub(v point) point {
	return point{s.X - v.X, s.Y - v.Y}
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func (s point) Mag() int {
	return abs(s.X) + abs(s.Y)
}

func pointEquals(a, b point) bool {
	return a.X == b.X && a.Y == b.Y
}

func lineToPoints(pos point, dir point, mag int) (points []point) {
	for i := 1; i <= mag; i++ {
		points = append(points, pos.Add(dir.Mul(i)))
	}
	return
}

func calcNextPoints(pos point, sDir string) ([]point, error) {
	mag, err := strconv.Atoi(sDir[1:])
	if err != nil {
		return nil, err
	}

	switch sDir[0] {
	case 'R':
		return lineToPoints(pos, point{1, 0}, mag), nil
	case 'L':
		return lineToPoints(pos, point{-1, 0}, mag), nil
	case 'U':
		return lineToPoints(pos, point{0, 1}, mag), nil
	case 'D':
		return lineToPoints(pos, point{0, -1}, mag), nil
	default:
		return nil, fmt.Errorf("Unknown direction: %c", sDir[0])
	}
}

func calculatePath(path string) (points []point) {

	sDirs := strings.Split(path, ",")

	p := point{0, 0}
	// points = append(points, p)
	for _, sDir := range sDirs {
		nextPoints, err := calcNextPoints(p, sDir)

		if err != nil {
			fmt.Printf("Error when printing %s\n", sDir)
			continue
		}

		points = append(points, nextPoints...)
		p = points[len(points)-1]
	}

	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func manhattanDistance(input string) (distance int) {
	// R75,D30,R83\nU62,R66,U55
	sPaths := strings.Split(strings.Replace(input, "\r", "", 0), "\n")

	path1 := calculatePath(sPaths[0])
	path2 := calculatePath(sPaths[1])

	pointSet := make(map[point]bool, len(path1))

	for _, p := range path1 {
		pointSet[p] = true
	}

	minD := math.MaxInt32

	for _, p := range path2 {
		if _, ok := pointSet[p]; ok {
			minD = min(minD, p.Mag())
		}
	}

	if minD == math.MaxInt32 {
		return -1
	}

	return minD
}

func stepDistance(input string) (steps int) {

	sPaths := strings.Split(strings.Replace(input, "\r", "", 0), "\n")

	path1 := calculatePath(sPaths[0])
	path2 := calculatePath(sPaths[1])

	pointSet := make(map[point]int, len(path1))

	for i, p := range path1 {
		pointSet[p] = i
	}

	minD := math.MaxInt32

	for i, p := range path2 {
		if v, ok := pointSet[p]; ok {
			minD = min(minD, i+v+2)
		}
	}

	if minD == math.MaxInt32 {
		return -1
	}

	return minD
}
