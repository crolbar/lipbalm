package layout


type Rect struct {
	x      uint16
	y      uint16
	width  uint16
	height uint16
}

func NewRect(x, y, width, height uint16) Rect {
	return Rect{
		x:      x,
		y:      y,
		width:  width,
		height: height,
	}
}

