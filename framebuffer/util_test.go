package framebuffer

import (
	"github.com/crolbar/lipbalm"
	"github.com/crolbar/lipbalm/assert"
	"strings"
	"testing"
)

func TestEnsureSizeExpand(t *testing.T) {
	var (
		str    = "0128347091207840712378478127384781237478912378478078aosntehu\naohuathaoeu\naoentuh\ntisheu\n83nteud"
		width  = 50
		height = 20
	)

	str = ensureSize(str, uint16(width), uint16(height))

	// fmt.Printf("%q\n", str)

	lines := strings.Split(str, "\n")

	assert.Equal(t, height, len(lines))
	assert.Equal(t, 60, len(lines[0]))
	for _, line := range lines[1:] {
		assert.Equal(t, 50, len(line))
	}
}

func TestEnsureSizeExpandBottomRight(t *testing.T) {
	var (
		str    = ""
		width  = 34
		height = 72
	)

	str = ensureSize(str, uint16(width), uint16(height), lipbalm.Right, lipbalm.Bottom)

	// fmt.Printf("%q\n", str)

	lines := strings.Split(str, "\n")

	assert.Equal(t, height, len(lines))
	for _, line := range lines {
		assert.Equal(t, width, len(line))
	}
}

func TestEnsureSizeNotExpand(t *testing.T) {
	var (
		str    = "eu\nhelllaanltoheu"
		width  = 5
		height = 1
	)

	str = ensureSize(str, uint16(width), uint16(height), lipbalm.Center, lipbalm.Top)

	lines := strings.Split(str, "\n")

	assert.Equal(t, 2, len(lines))
	assert.Equal(t, width, len(lines[0]))
	assert.Equal(t, 14, len(lines[1]))
}
