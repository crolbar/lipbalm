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
		paddingBottom = paddingHeight - paddingSplit
		paddingTop    = paddingHeight - paddingBottom

		b strings.Builder
	)

	// top padding, bottom alignment
	if len(lines) < int(height) && valignment > lipbalm.Top {
		for range paddingTop {
			b.WriteString(paddingLine)
			b.WriteByte('\n')
		}
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

	// bottom padding, top alignment
	if len(lines) < int(height) && valignment < lipbalm.Bottom {
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
	case lipbalm.Right:
		b.WriteString(padding)
		b.WriteString(line)

	case lipbalm.Left:
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

func getAlignments(alignments []lipbalm.Position) (halignment, valignment lipbalm.Position) {
	halignment = lipbalm.Left
	valignment = lipbalm.Top
	if len(alignments) > 0 {
		halignment = alignments[0]
	}

	if len(alignments) > 1 {
		valignment = alignments[1]
	}

	return halignment, valignment
}

var emptyColor = color{}
var emptyCell = cell{' ', emptyColor, emptyColor}

func genBuffer(width, height int) [][]cell {
	buff := make([][]cell, height)
	line := make([]cell, height*width)

	for i := range buff {
		buff[i] = line[i*width : (i+1)*width]
	}

	for i := range line {
		line[i] = emptyCell
	}

	return buff
}

func strToUint8(s string) (u uint8) {
	for _, c := range s {
		u = u*10 + uint8(c-'0')
	}

	return
}

// string(ansi code) -> color
func parseAnsi(ansi []rune) (fg color, bg color) {
	ansi = ansi[2:]

	var (
		codes = strings.Split(string(ansi), ";")
		i     = 0
	)

	for i < len(codes) {
		var (
			code = codes[i]

			cur *color = nil
		)

		// fg/bg
		switch code {
		case "38": // fg
			cur = &fg
			cur.mode = fg256
		case "48": // bg
			cur = &bg
			cur.mode = bg256
		}

		if cur == nil {
			break
		}

		i++
		// color type
		switch codes[i] {
		case "5": // 256
			i++
			cur.vals[0] = strToUint8(codes[i])
			i++
		case "2": // tc
			cur.mode++

			i++
			for j := range 3 {
				cur.vals[j] = strToUint8(codes[i])
				i++
			}
		}
	}

	return
}

// []rune -> []cell
// with ansi codes extracted for each rune
func convertLineToCells(line []rune, width uint16) []cell {
	var (
		cells = make([]cell, width)
		j     = 0

		ansiStartIdx = -1

		fgcolor color
		bgcolor color
	)

	for i, r := range line {
		// line (without ansi) is longer that the rect's width
		if j >= int(width) {
			break
		}

		// start ansi escape
		if r == '\x1b' {
			ansiStartIdx = i
			continue
		}

		if ansiStartIdx >= 0 {
			if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
				fgcolor, bgcolor = parseAnsi(line[ansiStartIdx:i])
				ansiStartIdx = -1
			}
			continue
		}

		// other invisible ascii chars
		if r < 32 || r == 127 {
			continue
		}

		cells[j] = cell{
			rune: r,
			fg:   fgcolor,
			bg:   bgcolor,
		}
		j++
	}

	return cells
}
