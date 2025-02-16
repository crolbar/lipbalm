package layout

func (l Layout) split(rect Rect) []Rect {
	splits := make([]Rect, 0)

	switch l.direction {
	case Vertical:
	case Horizontal:
		nextY := rect.y
		for _, c := range l.constrains {
			hSplit := splitHorizontal(rect, nextY, c)
			nextY = hSplit.y + hSplit.height
			splits = append(splits, hSplit)
		}
	}

	return splits
}

func splitHorizontal(
	rect Rect,
	nextY uint16,
	constrain Constrain,
) Rect {
	r := Rect{
		y: nextY,
		x: rect.x,
		width: rect.width,
	}

	switch constrain.t {
	case Lenght:
		r.height = constrain.v
	}

	return r
}
