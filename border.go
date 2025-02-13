package lipbalm

import "strings"

// ColorFg/ColorBg:
//
//	ansi color code (use Color/ColorRGB)
type BorderType struct {
	Top         string
	Bottom      string
	Left        string
	Right       string
	TopLeft     string
	TopRight    string
	BottomLeft  string
	BottomRight string
	ColorFg     string
	ColorBg     string
}

// color:
//
//	ansi color code (use Color/ColorRGB)
//	[0] for foreground, [1] for background
func NormalBorder(color ...string) BorderType {
	colorFg, colorBg := getColor(color)

	return BorderType{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "┌",
		TopRight:    "┐",
		BottomLeft:  "└",
		BottomRight: "┘",
		ColorFg:     colorFg,
		ColorBg:     colorBg,
	}
}

func RoundedBorder(color ...string) BorderType {
	colorFg, colorBg := getColor(color)

	return BorderType{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "╰",
		BottomRight: "╯",
		ColorFg:     colorFg,
		ColorBg:     colorBg,
	}
}

func getColor(color []string) (fg, bg string) {
	var (
		colorFg = ""
		colorBg = ""
	)

	switch len(color) {
	case 2:
		colorBg = color[1]
		fallthrough
	case 1:
		colorFg = color[0]
	}

	return colorFg, colorBg
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

func Border(
	b BorderType,
	str string,
	disabled ...bool,
) string {
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

		fullColorLen = len(b.ColorFg) + len(b.ColorBg) +
			iff(b.ColorBg != "" || b.ColorFg != "", len(ansi_reset), 0)

		sbSize = len(str) +
			iff(top, len(topBorder)+1+fullColorLen, 0) + // top border
			iff(bottom, len(bottomBorder)+1+fullColorLen, 0) + // bottom border
			iff(left, numLines*(len(b.Left)+fullColorLen), 0) + // left border
			iff(right, numLines*(len(b.Right)+fullColorLen), 0) // right border

		writeToSb = writeStringToSb(&sb)

		applyColor = func() {
			if b.ColorBg != "" && b.ColorFg != "" {
				sb.WriteString(
					b.ColorFg[:len(b.ColorFg)-1] + ";" +
						b.ColorBg[2:])
			} else {
				doIf(b.ColorFg, writeToSb)
				doIf(b.ColorBg, writeToSb)
			}
		}
		resetColor = func() {
			doIfp(b.ColorFg != "" || b.ColorBg != "",
				ansi_reset, writeToSb)
		}
	)

	sb.Grow(sbSize)

	if top {
		applyColor()
		sb.WriteString(topBorder)
		resetColor()

		sb.WriteByte('\n')
	}

	for i, line := range lines {
		if left {
			applyColor()
			sb.WriteString(b.Left)
			resetColor()
		}

		sb.WriteString(line)

		if right {
			applyColor()
			sb.WriteString(b.Right)
			resetColor()
		}

		if i == lastLineIdx {
			break
		}

		sb.WriteByte('\n')
	}

	if bottom {
		sb.WriteByte('\n')

		applyColor()
		sb.WriteString(bottomBorder)
		resetColor()
	}

	return sb.String()
}
