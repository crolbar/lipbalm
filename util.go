package lipbalm

import (
	"math"
	"strings"
)

// adds padding to a multiline string to make it square-ish
// so every line is equal width
// position specifIes where to add the padding: Left, Right, Center
func MakeSquare(position Position, str string) string {
	if len(str) == 0 {
		return ""
	}

	lines, widths, maxWidth := GetLines(str)

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

			split := int(math.Round(float64(padding_width) * position.Value()))
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

func iff[T int | string | any](b bool, f, s T) T {
	if b {
		return f
	}
	return s
}

func writeStringToSb(sb *strings.Builder) func(s string) {
	return func(s string) {
		sb.WriteString(s)
	}
}

func doIf[T string](b T, c func(s T)) {
	if b != "" {
		c(b)
	}
}

func doIfp[B bool | string, T string](b B, w T, c func(s T)) {
	switch v := any(b).(type) {
	case string:
		if v != "" {
			c(w)
		}
	case bool:
		if v {
			c(w)
		}
	}
}

func getAlignments(alignments []Position) (halignment, valignment Position) {
	halignment = Right
	valignment = Right
	if len(alignments) > 0 {
		halignment = alignments[0]
	}

	if len(alignments) > 1 {
		valignment = alignments[1]
	}

	return halignment, valignment
}
