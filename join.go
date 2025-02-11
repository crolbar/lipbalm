package lipbalm

import (
	"math"
	"strings"
)

// JoinHorizontal is a utility function for horizontally joining two
// potentially multi-lined strings along a vertical axis. The first argument is
// the position, with 0 being all the way at the top and 1 being all the way
// at the bottom.
//
// If you just want to align to the top, center or bottom you may as well just
// use the helper constants Top, Center, and Bottom.
//
// Example:
//
//	blockB := "...\n...\n..."
//	blockA := "...\n...\n...\n...\n..."
//
//	// Join 20% from the top
//	str := lipgloss.JoinHorizontal(0.2, blockA, blockB)
//
//	// Join on the top edge
//	str := lipgloss.JoinHorizontal(lipgloss.Top, blockA, blockB)
func JoinHorizontal(pos Position, strs ...string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	var (
		// Groups of strings broken into multiple lines
		blocks = make([][]string, len(strs))

		// Max line widths for the above text blocks
		maxWidths = make([]int, len(strs))

		// Height of the tallest block
		maxHeight int
	)

	// Break text blocks into lines and get max widths for each text block
	for i, str := range strs {
		blocks[i], _, maxWidths[i] = getLines(str)
		if len(blocks[i]) > maxHeight {
			maxHeight = len(blocks[i])
		}
	}

	// Add extra lines to make each side the same height
	for i := range blocks {
		if len(blocks[i]) >= maxHeight {
			continue
		}

		extraLines := make([]string, maxHeight-len(blocks[i]))

		switch pos { //nolint:exhaustive
		case Top:
			blocks[i] = append(blocks[i], extraLines...)

		case Bottom:
			blocks[i] = append(extraLines, blocks[i]...)

		default: // Somewhere in the middle
			n := len(extraLines)
			split := int(math.Round(float64(n) * pos.value()))
			top := n - split
			bottom := n - top

			blocks[i] = append(extraLines[top:], blocks[i]...)
			blocks[i] = append(blocks[i], extraLines[bottom:]...)
		}
	}

	// Merge lines
	var b strings.Builder
	for i := range blocks[0] { // remember, all blocks have the same number of members now
		for j, block := range blocks {
			b.WriteString(block[i])

			// Also make lines the same length
			b.WriteString(strings.Repeat(" ", maxWidths[j]-GetStringWidth(block[i])))
		}
		if i < len(blocks[0])-1 {
			b.WriteRune('\n')
		}
	}

	return b.String()
}

// joins multiline strings and adds padding to each line
// so they have equal widths
func JoinVertical(pos Position, strs ...string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	var (
		blocks   = make([][]string, len(strs))
		widths   = make([][]int, len(strs))
		maxWidth int
		numLines int
	)

	for i := range strs {
		var w int
		blocks[i], widths[i], w = getLines(strs[i])

		numLines += len(blocks[i])
		if w > maxWidth {
			maxWidth = w
		}
	}

	var (
		b         strings.Builder
		blocksLen = len(blocks)
	)

	b.Grow(((maxWidth + 1) * numLines) - 1)
	for i, block := range blocks {
		var (
			atLastBlock = i == blocksLen-1
			blockLen    = len(block)
		)

		for j, line := range block {
			var (
				atLastLine = j == blockLen-1

				padding_width = maxWidth - widths[i][j]
				padding       = strings.Repeat(" ", padding_width)
			)

			switch pos {
			case Left:
				b.WriteString(line)
				b.WriteString(padding)

			case Right:
				b.WriteString(padding)
				b.WriteString(line)

			default:
				if padding_width < 1 {
					b.WriteString(line)
					break
				}

				// split the padding and put it to the start & end of the line
				// so we get a line centered
				split := int(math.Round(float64(padding_width) * pos.value()))
				right := padding_width - split
				left := padding_width - right

				b.WriteString(padding[0:left])
				b.WriteString(line)
				b.WriteString(padding[0:right])
			}

			if atLastBlock && atLastLine {
				break
			}

			b.WriteString("\n")
		}
	}

	return b.String()
}
