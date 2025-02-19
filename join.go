package lipbalm

import (
	"math"
	"strings"
)

// joins multiline strings horizontally
//
// Example:
//
//   str1 = " string \n one "
//   str2 = " string \n two "
//
//   we will get:
//   res = " string  string \n one     two    "
//
//   so the first lines of the strings join, the second lines join and so on.
//   also adds padding to each line so it becomes as wide as the widest
//   line is that string, so we get a square
func JoinHorizontal(pos Position, strs ...string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	var (
		blocks    = make([][]string, len(strs))
		maxWidths = make([]int, len(strs))
		widths    = make([][]int, len(strs))
		maxHeight int
	)

	for i, str := range strs {
		blocks[i], widths[i], maxWidths[i] = GetLines(str)

		lines := len(blocks[i])
		if lines > maxHeight {
			maxHeight = lines
		}
	}

	// add padding(lines) to blocks withs less than maxHeight lines
	// so we get a nice square
	for i := range blocks {
		lines := len(blocks[i])

		if lines >= maxHeight {
			continue
		}

		var (
			extraLines  = make([]string, maxHeight-lines)
			extraWidths = make([]int, maxHeight-lines)
		)

		switch pos {
		case Top:
			blocks[i] = append(blocks[i], extraLines...)
			widths[i] = append(widths[i], extraWidths...)

		case Bottom:
			blocks[i] = append(extraLines, blocks[i]...)
			widths[i] = append(extraWidths, widths[i]...)

		default:
			var (
				n      = len(extraLines)
				split  = int(math.Round(pos.Value() * float64(n)))
				top    = n - split
				bottom = n - top
			)

			blocks[i] = append(extraLines[top:], blocks[i]...)
			widths[i] = append(extraWidths[top:], widths[i]...)

			blocks[i] = append(blocks[i], extraLines[bottom:]...)
			widths[i] = append(widths[i], extraWidths[bottom:]...)
		}
	}

	lastBlockIdx := len(blocks[0]) - 1

	var b strings.Builder
	for i := range blocks[0] {
		for j, block := range blocks {
			b.WriteString(block[i])

			var (
				blockMaxWidth = maxWidths[j]
				lineWidth     = widths[j][i]

				diff = blockMaxWidth - lineWidth
			)

			if diff > 0 {
				b.WriteString(strings.Repeat(" ", diff))
			}
		}

		if i == lastBlockIdx {
			continue
		}

		b.WriteByte('\n')
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
		blocks[i], widths[i], w = GetLines(strs[i])

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
				split := int(math.Round(float64(padding_width) * pos.Value()))
				right := padding_width - split
				left := padding_width - right

				b.WriteString(padding[0:left])
				b.WriteString(line)
				b.WriteString(padding[0:right])
			}

			if atLastBlock && atLastLine {
				break
			}

			b.WriteByte('\n')
		}
	}

	return b.String()
}
