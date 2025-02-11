package lipbalm

import (
	"testing"
	"github.com/crolbar/lipbalm/assert"
)

func TestStringWidthSimple(t *testing.T) {
	s := "\nasotheu\n"
	assert.Equal(t, 7, GetStringWidth(s))
}

func TestStringWidthAnsi(t *testing.T) {
	s := "\033[1;31mBold Red\033[0m \033[3;34mItalic Blue\033[0m\033[4;32mUnderlined Green\033[0m\n\033[1;35mBold Magenta\033[0m Normal Text"
	assert.Equal(t, 60, GetStringWidth(s))
}

func TestStringWidthAnsi2(t *testing.T) {
	s := "[97m│[0m[38;5;57m█████████████[0m[97m│[0m"
	assert.Equal(t, 15, GetStringWidth(s))
}
