package framebuffer

import (
	"github.com/crolbar/lipbalm"
	"math"
	"strings"
)

func ensureSize(str string, width, height uint16, alignments ...lipbalm.Position) string {
	var (
		halignment, valignment = getAlignments(alignments)

		lines, widths, _ = lipbalm.GetLines(str)

		lastLineIdx = len(lines) - 1

		paddingHeight = int(height) - len(lines)
		paddingLine   = strings.Repeat(" ", int(width))
		paddingSplit  = int(math.Round(float64(paddingHeight) * valignment.Value()))
		paddingTop    = paddingHeight - paddingSplit
		paddingBottom = paddingHeight - paddingTop

		b strings.Builder
	)

	// top padding
	if len(lines) < int(height) && valignment < lipbalm.Bottom {
		b.WriteString(strings.Repeat("\n"+paddingLine, paddingTop))
	}

	for i, line := range lines {
		lineWidth := widths[i]

		// grow width
		if uint16(lineWidth) < width {
			applyPadding(&b, line, halignment, int(width), lineWidth)
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

	// bottom padding
	if len(lines) < int(height) && valignment > lipbalm.Top {
		b.WriteString(strings.Repeat("\n"+paddingLine, paddingBottom))
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

// returns idx of the nth visable character in the string
// without treating ansi codes as visable characters
func getWithoutAnsi(n int, s string) int {
	if n == 0 {
		return n
	}

	var (
		lastVisibleIdx = 0
		visableCount   = 0
		skiping        = false
		runes          = []rune(s)
	)

	for i, r := range runes {
		// start ansi escape
		if r == '\x1b' {
			skiping = true
			continue
		}

		// continue/stop ansi escape
		if skiping {
			if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
				skiping = false
			}
			continue
		}

		// other invisible ascii chars
		if r < 32 || r == 127 {
			continue
		}

		visableCount++
		lastVisibleIdx = i

		// visableCount start with 1,
		// and n is an index so it start with 0, so n+1
		if visableCount >= n+1 {
			// if we have a trailing ansi sequence.
			// not the best approach
			// but im expeting the area after n to be empty
			if i < len(runes)-1 && runes[i+1] == '\x1b' {
				continue
			}

			return i
		}
	}

	return lastVisibleIdx
}

func getAlignments(alignments []lipbalm.Position) (halignment, valignment lipbalm.Position) {
	halignment = lipbalm.Right
	valignment = lipbalm.Right
	if len(alignments) > 0 {
		halignment = alignments[0]
	}

	if len(alignments) > 1 {
		valignment = alignments[1]
	}

	return halignment, valignment
}
