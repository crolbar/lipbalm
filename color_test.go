package lipbalm

import (
	"github.com/crolbar/lipbalm/assert"
	"testing"
)

func TestColor(t *testing.T) {
	s := "hi"

	balm := SetColor(Color(245),
		SetColor(ColorBg(233), s))

	assert.Equal(t, "[38;5;245;48;5;233mhi[0m", balm)
}

func TestColorMultiline(t *testing.T) {
	s := "hiyo\noeuo\nlnes\nmore\nline\nyess"
	assert.Equal(t,
		"\x1b[38;5;245;48;5;233mhiyo\x1b[0m\n\x1b[38;5;245;48;5;233moeuo\x1b[0m\n\x1b[38;5;245;48;5;233mlnes\x1b[0m\n\x1b[38;5;245;48;5;233mmore\x1b[0m\n\x1b[38;5;245;48;5;233mline\x1b[0m\n\x1b[38;5;245;48;5;233myess\x1b[0m",
		SetColor(Color(245),
			SetColor(ColorBg(233),
				MakeSquare(Right, s),
			),
		),
	)
}
