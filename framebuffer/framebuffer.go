package framebuffer

import (
	"strings"

	"github.com/crolbar/lipbalm/layout"
)

type FrameBuffer struct {
	height uint16 // height of the frame
	width  uint16 // width of the frame

	// the string that we will output
	//
	// []string because we store the lines
	// and then join them on final output
	frame []string
}

func NewFrameBuffer(width, height uint16) FrameBuffer {
	var (
		line  = strings.Repeat(" ", int(width))
		frame = make([]string, height)
	)

	for i := range frame {
		frame[i] = line
	}

	return FrameBuffer{
		height: height,
		width:  width,
		frame:  frame,
	}
}

func (f *FrameBuffer) RenderString(
	str string,
	rect layout.Rect,
) {
	if int(rect.Y) > len(f.frame) {
		panic("rect.y out of bounds")
	}

	str = ensureSize(str, rect.Width, rect.Height)

	var (
		lines = strings.Split(str, "\n")
	)

	for i, line := range lines {
		var (
			frameLineIdx = rect.Y + uint16(i)
			frameLine    = f.frame[frameLineIdx]
		)
		if int(rect.X) > len(frameLine) {
			panic("rect.x out of bounds")
		}

		f.frame[frameLineIdx] =
			frameLine[:rect.X] + // everything before the rect's x, its excluding so x is free
				line +
				frameLine[rect.X+rect.Width:] // everything after x + width
	}
}

func (f FrameBuffer) View() string {
	return strings.Join(f.frame, "\n")
}
