package layout

func (l Layout) Split(rect Rect) []Rect {
	var (
		numOfSplits = len(l.constrains)
		size        uint16
	)

	switch l.direction {
	// when spiting horizontally we split the height
	case Horizontal:
		size = rect.width

	// when spiting vertically we split the height
	case Vertical:
		size = rect.height
	}


	var (
		alocatedSizes = make([]uint16, numOfSplits)
		remainingSize = size

		// type => [idx, value] (idx in l.constrains list)
		constrainsMap = make(map[ConstrainType][][2]uint16, numOfSplits)
	)

	// fill map
	for i, c := range l.constrains {
		constrainsMap[c.t] = append(constrainsMap[c.t], [2]uint16{
			uint16(i),
			c.v,
		})
	}

	// Length
	takeFixedSize(&alocatedSizes, &remainingSize, constrainsMap[Length])
	// Min
	takeFixedSize(&alocatedSizes, &remainingSize, constrainsMap[Min])

	// Percent
	takePercentOfRem(&alocatedSizes, &remainingSize, size, constrainsMap[Percent])

	// Max
	takeMaxOfRem(&alocatedSizes, &remainingSize, constrainsMap[Max])

	// remainder goes Min
	takeRem(&alocatedSizes, &remainingSize, constrainsMap[Min])

	var (
		splits = make([]Rect, numOfSplits)
		x      = rect.x
		y      = rect.y
		width  = rect.width
		height = rect.height
	)

	for i, size := range alocatedSizes {
		var rect Rect
		switch l.direction {
		case Vertical:
			rect = Rect{
				x:      x,
				y:      y,
				width:  width,
				height: size,
			}
			y += size
		case Horizontal:
			rect = Rect{
				x:      x,
				y:      y,
				width:  size,
				height: height,
			}
			x += size
		}
		splits[i] = rect
	}

	return splits
}

func takeRem(
	alocatedSizes *[]uint16,
	remainingSize *uint16,
	minConstrains [][2]uint16,
) {
	if len(minConstrains) == 0 {
		return
	}

	var (
		numMin      = uint16(len(minConstrains))
		remSize     = *remainingSize
		spaceToTake = remSize / numMin
	)

	for _, c := range minConstrains {
		idx := c[0]

		// take fixed min + remainder
		(*alocatedSizes)[idx] += spaceToTake
		*remainingSize -= spaceToTake
	}
}

func takeMaxOfRem(
	alocatedSizes *[]uint16,
	remainingSize *uint16,
	maxConstrains [][2]uint16,
) {
	if len(maxConstrains) == 0 {
		return
	}

	var (
		numMax  = uint16(len(maxConstrains))
		remSize = *remainingSize

		// TODO: check if works correct
		spaceToTake = remSize / numMax // take full remainder
	)

	for _, c := range maxConstrains {
		var (
			idx  = c[0]
			_max = c[1]
		)

		// clamp to max
		if spaceToTake > _max {
			spaceToTake = _max
		}

		(*alocatedSizes)[idx] = spaceToTake
		*remainingSize -= spaceToTake
	}
}

func takePercentOfRem(
	alocatedSizes *[]uint16,
	remainingSize *uint16,
	fullSize uint16,
	percentConstrains [][2]uint16,
) {
	for _, c := range percentConstrains {
		var (
			idx     = c[0]
			remSize = *remainingSize

			percent     = c[1]
			spaceToTake = uint16((float32(percent) / 100 * float32(fullSize))) // take % from the full size
		)

		if spaceToTake > remSize {
			spaceToTake = remSize
		}

		(*alocatedSizes)[idx] = spaceToTake
		*remainingSize -= spaceToTake
	}
}

func takeFixedSize(
	alocatedSizes *[]uint16,
	remainingSize *uint16,
	fixedConstrains [][2]uint16,
) {
	for _, c := range fixedConstrains {
		var (
			remSize     = *remainingSize
			idx         = c[0]
			spaceToTake = c[1]
		)

		if spaceToTake > remSize {
			spaceToTake = remSize
		}

		(*alocatedSizes)[idx] = spaceToTake
		*remainingSize -= spaceToTake
	}
}
