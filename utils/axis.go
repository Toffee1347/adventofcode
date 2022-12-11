package utils

import "math"

func AxisGetDistanceBetweenPoints(pointA Coordinate, pointB Coordinate) float64 {
	length := Coordinate{
		X: pointA.X - pointB.X,
		Y: pointA.Y - pointB.Y,
	}

	lengthSquared := length.X*length.X + length.Y*length.Y
	return math.Sqrt(float64(lengthSquared))
}
