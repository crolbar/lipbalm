package lipbalm

import (
	"math"
	"strings"
)

func Expand(height, width int, str string, alingments ...Position) string {
	vAlignment, hAlignment := getAlignments(alingments)
	return ExpandHorizontal(width, hAlignment,
		ExpandVertical(height, vAlignment, str),
	)
}

// expands the string with adding new lines to the Position specifiend
// Top, Bottom, Center, to reach the specified `height`
func ExpandVertical(height int, alignment Position, str string) string {
	if len(str) == 0 {
		return strings.Repeat("\n", height-1)
	}

	numLines := strings.Count(str, "\n") + 1

	if height < numLines {
		return str
	}

	var (
		padding_height = height - numLines
		padding        = strings.Repeat("\n", padding_height)
	)

	switch alignment {
	case Top:
		return str + padding
	case Bottom:
		return padding + str
	default:
		var (
			split = int(math.Round(float64(padding_height) * alignment.Value()))
			right = padding_height - split
			left  = padding_height - right
		)

		return padding[0:left] + str + padding[0:right]
	}
}

// expands all lines of a multiline string to `width` width
// or if n < maxWidth of the lines, maxWidth will be used
//
// can be use as alignment function if used with n=0
func ExpandHorizontal(width int, alignment Position, str string) string {
	var (
		lines, widths, maxWidth = GetLines(str)
		numLines                = len(lines)

		sb strings.Builder
	)

	if width < maxWidth {
		width = maxWidth
	}

	for i, line := range lines {
		var (
			atLastLine = i == numLines-1

			padding_width = width - widths[i]
			padding       = strings.Repeat(" ", padding_width)
		)

		switch alignment {
		case Left:
			sb.WriteString(line)
			sb.WriteString(padding)

		case Right:
			sb.WriteString(padding)
			sb.WriteString(line)

		default:
			if padding_width < 1 {
				sb.WriteString(line)
				break
			}

			split := int(math.Round(float64(padding_width) * alignment.Value()))
			right := padding_width - split
			left := padding_width - right

			sb.WriteString(padding[0:left])
			sb.WriteString(line)
			sb.WriteString(padding[0:right])
		}

		if atLastLine {
			break
		}

		sb.WriteByte('\n')

	}

	return sb.String()
}
