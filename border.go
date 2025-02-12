package lipbalm

import "strings"

type BorderType struct {
	Top         string
	Bottom      string
	Left        string
	Right       string
	TopLeft     string
	TopRight    string
	BottomLeft  string
	BottomRight string
}

var (
	normalBorder = BorderType{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "┌",
		TopRight:    "┐",
		BottomLeft:  "└",
		BottomRight: "┘",
	}

	roundedBorder = BorderType{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "╰",
		BottomRight: "╯",
	}
)

func NormalBorder() BorderType {
	return normalBorder
}

func RoundedBorder() BorderType {
	return roundedBorder
}

func getDisabled(disabled []bool) (top bool, right bool, bottom bool, left bool) {
	switch len(disabled) {
	case 4:
		left = disabled[3]
		fallthrough
	case 3:
		bottom = disabled[2]
		fallthrough
	case 2:
		right = disabled[1]
		fallthrough
	case 1:
		top = disabled[0]
	}

	return !top, !right, !bottom, !left
}

func Border(b BorderType, str string, disabled ...bool) string {
	str = MakeSquare(str, Right)

	var (
		top, right, bottom, left = getDisabled(disabled)

		lines    = strings.Split(str, "\n")
		width    = GetStringWidth(lines[0])
		numLines = len(lines)

		topBorder    = iff(left, b.TopLeft, "") + strings.Repeat(b.Top, width) + iff(right, b.TopRight, "")
		bottomBorder = iff(left, b.BottomLeft, "") + strings.Repeat(b.Bottom, width) + iff(right, b.BottomRight, "")

		sb strings.Builder

		lastLineIdx = numLines - 1

		sbSize = len(str) +
			iff(top, len(topBorder)+1, 0) +
			iff(bottom, len(bottomBorder)+1, 0) +
			iff(left, numLines*len(b.Left), 0) +
			iff(right, numLines*len(b.Right), 0)
	)

	sb.Grow(sbSize)

	if top {
		sb.WriteString(topBorder)
		sb.WriteByte('\n')
	}

	for i, line := range lines {
		if left {
			sb.WriteString(b.Left)
		}

		sb.WriteString(line)

		if right {
			sb.WriteString(b.Right)
		}

		if i == lastLineIdx {
			break
		}

		sb.WriteByte('\n')
	}

	if bottom {
		sb.WriteByte('\n')
		sb.WriteString(bottomBorder)
	}

	return sb.String()
}
