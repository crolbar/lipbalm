package framebuffer

import (
	"github.com/crolbar/lipbalm"
	"math"
	"strings"
)

func ensureSize(str string, width, height uint16) string {
	var (
		// TODO
		position = lipbalm.Right
	)

	lines, widths, _ := lipbalm.GetLines(str)

	var b strings.Builder
	lastLineIdx := len(lines) - 1
	for i, line := range lines {
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

		// shring height
		if i >= int(height)-1 || i == lastLineIdx {
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
