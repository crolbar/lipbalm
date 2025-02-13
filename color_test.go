package lipbalm

import (
	"testing"
	"github.com/crolbar/lipbalm/assert"
)

func TestColor(t *testing.T) {
	s := "hi"

	balm := SetColor(Color(245),
		SetColor(ColorBg(233), s))

	assert.Equal(t, "[38;5;245;48;5;233mhi[0m", balm)
}
