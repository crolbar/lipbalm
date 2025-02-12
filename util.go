package lipbalm

import (
	"math"
	"strings"
)

// adds padding to a multiline string to make it square-ish
// so every line is equal width
// position specifIes where to add the padding: Left, Right, Center
func MakeSquare(str string, position Position) string {
	if len(str) == 0 {
		return ""
	}

	lines, widths, maxWidth := getLines(str)

	if len(lines) == 1 {
		return str
	}

	var b strings.Builder
	lastLineIdx := len(lines) - 1
	for i, line := range lines {
		var (
			padding_width = maxWidth - widths[i]
			padding       = strings.Repeat(" ", padding_width)
		)

		switch position {
		case Left:
			b.WriteString(padding)
			b.WriteString(line)

		case Right:
			b.WriteString(line)
			b.WriteString(padding)

		default:
			if padding_width < 1 {
				b.WriteString(line)
				break
			}

			split := int(math.Round(float64(padding_width) * position.value()))
			right := padding_width - split
			left := padding_width - right

			b.WriteString(padding[0:left])
			b.WriteString(line)
			b.WriteString(padding[0:right])
		}

		if i == lastLineIdx {
			break
		}

		b.WriteByte('\n')
	}

	return b.String()
}

func clamp[T int | float64](v, low, high T) T {
	return min(max(v, low), high)
}

func iff[T int | string](b bool, f, s T) T {
	if b {
		return f
	}
	return s
}
