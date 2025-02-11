package lipbalm

import (
	"strings"

	"github.com/charmbracelet/x/ansi/parser"
	"github.com/mattn/go-runewidth"
	"github.com/rivo/uniseg"
)

func getLines(s string) (lines []string, widest int) {
	lines = strings.Split(s, "\n")

	for _, l := range lines {
		w := StringWidth(l)
		if widest < w {
			widest = w
		}
	}

	return lines, widest
}


type Method uint8

const (
	WcWidth Method = iota
	GraphemeWidth
)

func StringWidth(s string) int {
	return stringWidth(GraphemeWidth, s)
}

func StringWidthWc(s string) int {
	return stringWidth(WcWidth, s)
}

func stringWidth(m Method, s string) int {
	if s == "" {
		return 0
	}

	var (
		pstate  = parser.GroundState // initial state
		cluster string
		width   int
	)

	for i := 0; i < len(s); i++ {
		state, action := parser.Table.Transition(pstate, s[i])
		if state == parser.Utf8State {
			var w int
			cluster, _, w, _ = uniseg.FirstGraphemeClusterInString(s[i:], -1)
			if m == WcWidth {
				w = runewidth.StringWidth(cluster)
			}
			width += w
			i += len(cluster) - 1
			pstate = parser.GroundState
			continue
		}

		if action == parser.PrintAction {
			width++
		}

		pstate = state
	}

	return width
}
