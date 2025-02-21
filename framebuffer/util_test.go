package framebuffer

import (
	"github.com/crolbar/lipbalm/assert"
	"strings"
	"testing"
)

func TestGetWithoutAnsi(t *testing.T) {
	assert.Equal(t, 3, getWithoutAnsi(3, []rune("Hello")))
	assert.Equal(t, 8, getWithoutAnsi(3, []rune("\x1b[31mHello")))
	// counts in the trailing escape sequence
	assert.Equal(t, 10, getWithoutAnsi(4, []rune("01234\x1b[31m56789")))
	assert.Equal(t, 12, getWithoutAnsi(7, []rune("Hello\x1b[31m, World")))
	assert.Equal(t, 4, getWithoutAnsi(4, []rune("Hello\x1b[31m")))
	assert.Equal(t, 10, getWithoutAnsi(4, []rune("Hello\x1b[31m5")))
	assert.Equal(t, 0, getWithoutAnsi(1, []rune("\x1b[31m")))
	assert.Equal(t, 14, getWithoutAnsi(9, []rune("Hello\x1b[31mWorld\x1b[31m")))
	assert.Equal(t, 20, getWithoutAnsi(9, []rune("Hello\x1b[31mWorld\x1b[31m!")))
}

func TestEnsureSize(t *testing.T) {
	var (
		str    = "0128347091207840712378478127384781237478912378478078aosntehu\naohuathaoeu\naoentuh\ntisheu\n83nteud"
		width  = 50
		height = 20
	)

	str = ensureSize(str, uint16(width), uint16(height))

	lines := strings.Split(str, "\n")

	assert.Equal(t, height, len(lines))
	assert.Equal(t, 60, len(lines[0]))
	for _, line := range lines[1:] {
		assert.Equal(t, 50, len(line))
	}

	str = ""
	width = 34
	height = 72

	str = ensureSize(str, uint16(width), uint16(height))

	lines = strings.Split(str, "\n")

	assert.Equal(t, height, len(lines))
	for _, line := range lines {
		assert.Equal(t, width, len(line))
	}

	str = "eu\nhelllaanltoheu"
	width = 5
	height = 1

	str = ensureSize(str, uint16(width), uint16(height))

	lines = strings.Split(str, "\n")

	assert.Equal(t, 2, len(lines))
	assert.Equal(t, width, len(lines[0]))
	assert.Equal(t, 14, len(lines[1]))
}
