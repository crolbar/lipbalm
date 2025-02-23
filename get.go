package lipbalm

import "strings"

// returns the string split on \n (so the lines)
// the width of each line
// and the largest width
func GetLines(s string) (lines []string, widths []int, widest int) {
	lines = strings.Split(s, "\n")
	widths = make([]int, len(lines))

	for i, l := range lines {
		w := GetStringWidth(l)
		widths[i] = w
		if widest < w {
			widest = w
		}
	}

	return lines, widths, widest
}

// returns the height of a multiline string
func GetHeight(s string) int {
	return strings.Count(s, "\n") + 1
}

// returns the max width of a multiline string
func GetWidth(str string) (maxWidth int) {
	for _, l := range strings.Split(str, "\n") {
		if w := GetStringWidth(l); w > maxWidth {
			maxWidth = w
		}
	}

	return maxWidth
}

// returns the width of the strings
// without ansi codes & chars < 32
func GetStringWidth(s string) int {
	var (
		width   = 0
		skiping = false
	)

	for _, r := range s {
		if r == '\x1b' {
			skiping = true
			continue
		}

		if skiping {
			if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
				skiping = false
			}
			continue
		}

		if r < 32 || r == 127 {
			continue
		}

		width++
	}

	return width
}
