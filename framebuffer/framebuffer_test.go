package framebuffer

import (
	"fmt"
	"strings"
	"testing"

	"github.com/crolbar/lipbalm"
	"github.com/crolbar/lipbalm/assert"
	"github.com/crolbar/lipbalm/layout"
)

func TestM(t *testing.T) {
	fb := NewFrameBuffer(10, 10)

	fb.RenderString(
		"yeahh\nhelllaanltoheu\nthree\nfour",
		layout.NewRect(3, 5, 5, 4),
	)

	fb.RenderString(
		"yeahh\nhelllaanltoheu",
		layout.NewRect(0, 7, 10, 1),
	)

	frame := fb.View()

	fmt.Println(lipbalm.Border(lipbalm.NormalBorder(), frame))

	// assert.Equal(t,
	// 	"          \n          \n          \n          \n          \n          \n          \n          \n          \n          ",
	// 	frame)
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
	for _, line := range lines {
		assert.Equal(t, width, len(line))
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

	str = "yeahh\nhelllaanltoheu"
	width = 5
	height = 1

	str = ensureSize(str, uint16(width), uint16(height))

	lines = strings.Split(str, "\n")

	assert.Equal(t, height, len(lines))
	for _, line := range lines {
		assert.Equal(t, width, len(line))
	}
}
