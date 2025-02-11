package main


type Position float64

const (
	Top    Position = 0.0
	Bottom Position = 1.0
	Center Position = 0.5
	Left   Position = 0.0
	Right  Position = 1.0
)

func (p Position) value() float64 {
	return clamp(float64(p), 0, 1)
}
