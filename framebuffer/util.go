package framebuffer

import (
	"github.com/crolbar/lipbalm"
	"math"
	"strings"
)

func ensureSize(str string, width, height uint16) string {
	var (
		// TODO
		position = lipbalm.Left
	)

	lines, widths, _ := lipbalm.GetLines(str)

	var b strings.Builder
	lastLineIdx := len(lines) - 1
	for i, line := range lines {

		// shring height
		if i >= int(height) {
			break
		}

		lineWidth := widths[i]

		// grow width
		if uint16(lineWidth) < width {
			applyPadding(&b, line, position, int(width), lineWidth)
		} else

		// shrink width
		if uint16(lineWidth) > width {
			b.WriteString(line[:width])
		} else

		// exact width
		{
			b.WriteString(line)
		}

		if i == lastLineIdx {
			break
		}

		b.WriteByte('\n')
	}

	// grow height
	if len(lines) < int(height) {
		paddingLine := strings.Repeat(" ", int(width))
		padding := strings.Repeat("\n"+paddingLine, int(height)-len(lines))
		b.WriteString(padding)
	}

	return b.String()
}

func applyPadding(
	b *strings.Builder,
	line string,
	position lipbalm.Position,
	maxWidth, width int,
) {
	var (
		padding_width = maxWidth - width
		padding       = strings.Repeat(" ", padding_width)
	)

	switch position {
	case lipbalm.Left:
		b.WriteString(padding)
		b.WriteString(line)

	case lipbalm.Right:
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

}
