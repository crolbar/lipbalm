package lipbalm

import "strings"

// returns the string split on \n (so the lines)
// the width of each line
// and the largest width
func getLines(s string) (lines []string, widths []int, widest int) {
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
