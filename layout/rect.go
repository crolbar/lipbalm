package layout

type Rect struct {
	X      uint16
	Y      uint16
	Width  uint16
	Height uint16
}

func NewRect(x, y, width, height uint16) Rect {
	return Rect{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
}
