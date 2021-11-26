package models

import "math"

type Coordinate struct {
    x,y float64
}

func NewCoordinate(x float64, y float64) *Coordinate {
    return &Coordinate{x: x, y: y}
}

func (c Coordinate) GetCoordinates() []float64 {
    return []float64{c.x, c.y}
}

func (c Coordinate) FindDistanceFrom(source Coordinate) float64  {
    return math.Sqrt(math.Pow(c.x-source.x, 2)+ math.Pow(c.y-source.y, 2))
}