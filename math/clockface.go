package clockface

import (
	"math"
	"time"
)

const (
	secondHandLength = 90
	clockCenterX     = 150
	clockCenterY     = 150
)

type Point struct {
	X float64
	Y float64
}

func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength} // scale
	p = Point{p.X, -p.Y}                                      //flip
	p = Point{p.X + clockCenterX, p.Y + clockCenterY}         //translate
	return p
}

func secondsInRadius(t time.Time) float64 {
	return (math.Pi / (30 / (float64(t.Second()))))
}

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadius(t)
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{X: x, Y: y}
}
