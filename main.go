package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func main() {
	points := [][2]float64{
		{0, 0},
		{0, 1},
		{0, 2},
		{1, 0},
		{1, 1},
		{1, 2},
		{2, 1},
		{2, 2},
	}
	fmt.Print(len(CountRectangles(points)))
}

func CountRectangles(points [][2]float64) map[string]bool {
	rectangles := make(map[string]bool)
	hashOfPoints := InitPointsHash(points)
	for i := 0; i < len(points); i++ {
		point1 := points[i]
		for j := i + 1; j < len(points); j++ {
			point2 := points[j]
			for k := j + 1; k < len(points); k++ {
				point3 := points[k]
				a, b, c, isRight := Pythagorain(point1, point2, point3)
				if !isRight {
					continue
				}
				point4 := [2]float64{
					c[0] + (b[0] - a[0]),
					c[1] + (b[1] - a[1]),
				}
				index, ok1 := hashOfPoints[point4]
				if !ok1 {
					continue
				}
				pointsIndices := []int{i, j, k, index}
				sort.Ints(pointsIndices)
				rectangleId := FormId(pointsIndices)
				_, ok2 := rectangles[rectangleId]
				if !ok2 {
					rectangles[rectangleId] = true
				}
			}
		}
	}
	return rectangles
}

func DistanceSquared(point1, point2 [2]float64) float64 {
	return RoundFloat(math.Pow(point1[0]-point2[0], 2)+math.Pow(point1[1]-point2[1], 2), 3)
}

func InitPointsHash(points [][2]float64) map[[2]float64]int {
	hashOfPoints := make(map[[2]float64]int)
	for index, point := range points {
		hashOfPoints[point] = index
	}
	return hashOfPoints
}

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func Pythagorain(a, b, c [2]float64) ([2]float64, [2]float64, [2]float64, bool) {
	if DistanceSquared(a, b)+
		DistanceSquared(b, c) ==
		DistanceSquared(a, c) {
		return b, a, c, true
	}
	if DistanceSquared(b, c)+
		DistanceSquared(a, c) ==
		DistanceSquared(a, b) {
		return c, a, b, true
	}
	if DistanceSquared(a, c)+
		DistanceSquared(b, a) ==
		DistanceSquared(b, c) {
		return a, b, c, true
	}
	return [2]float64{}, [2]float64{}, [2]float64{}, false
}

func FormId(pointsIndices []int) string {
	res := ""
	length := len(pointsIndices)
	for i := 0; i < length; i++ {
		res += strconv.Itoa(pointsIndices[i]) + "."
	}
	return res
}
