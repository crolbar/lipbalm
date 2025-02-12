package lipbalm

import "strings"

type marginPos int

const (
	top marginPos = iota
	bottom
	left
	right
)

// Adds margin to top,bottom,left,right
func Margin(margin int, str string) string {
	return applyMargin(margin, str, top, bottom, left, right)
}

// Adds margin to top,bottom
func MarginVertical(margin int, str string) string {
	return applyMargin(margin, str, top, bottom)
}

// Adds margin to left,right
func MarginHorizontal(margin int, str string) string {
	return applyMargin(margin, str, left, right)
}

// Adds margin to left
func MarginLeft(margin int, str string) string {
	return applyMargin(margin, str, left)
}

// Adds margin to right
func MarginRight(margin int, str string) string {
	return applyMargin(margin, str, right)
}

// Adds margin to top
func MarginTop(margin int, str string) string {
	return applyMargin(margin, str, top)
}

// Adds margin to bottom
func MarginBottom(margin int, str string) string {
	return applyMargin(margin, str, bottom)
}

func applyMargin(margin int, str string, pos ...marginPos) string {
	var (
		hasBottom = false
		hasTop    = false
		hasLeft   = false
		hasRight  = false
	)

	for _, p := range pos {
		switch p {
		case top:
			hasTop = true
		case bottom:
			hasBottom = true
		case left:
			hasLeft = true
		case right:
			hasRight = true
		}
	}

	lines, _, maxWidth := getLines(str)

	// horizontal
	if hasLeft || hasRight {
		var (
			b strings.Builder

			numLines         = len(lines)
			horizontalMargin = strings.Repeat(" ", margin)

			numAddedChars = margin * numLines

			lastLineIdx = numLines - 1
		)

		b.Grow(len(str) +
			iff(hasLeft, numAddedChars, 0) +
			iff(hasRight, numAddedChars, 0))

		for i, line := range lines {
			if hasLeft {
				b.WriteString(horizontalMargin)
			}

			b.WriteString(line)

			if hasRight {
				b.WriteString(horizontalMargin)
			}

			if i == lastLineIdx {
				break
			}

			b.WriteByte('\n')
		}

		str = b.String()
	}

	// vertical
	maxWidth = maxWidth +
		iff(hasLeft, margin, 0) +
		iff(hasRight, margin, 0)

	padding := strings.Repeat(" ", maxWidth)

	if hasTop {
		str = strings.Repeat(padding+"\n", margin) + str
	}

	if hasBottom {
		str = str + strings.Repeat("\n"+padding, margin)
	}

	return str
}
