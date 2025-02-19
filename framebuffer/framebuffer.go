package framebuffer

import (
	"fmt"
	"github.com/crolbar/lipbalm/layout"
	"strings"
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

func (f FrameBuffer) Size() layout.Rect {
	return layout.Rect{
		X:      0,
		Y:      0,
		Width:  f.width,
		Height: f.height,
	}
}

// write a string to the specified rect of the framebuffer.
// doesn't support writing to the same place two different times
func (f *FrameBuffer) RenderString(
	str string,
	rect layout.Rect,
) {
	var (
		endY = (rect.Y + rect.Height) - 1
		endX = (rect.X + rect.Width) - 1
	)

	if endX >= f.width || endY >= f.height {
		panic(fmt.Sprintf(
			"rect: %v went out of bounds of FrameBuffer: %v", rect, f,
		))
	}

	if rect.Width <= 0 || rect.Height <= 0 {
		return
	}

	// make sure that the string is exacly `rect.Width` width and `rect.Height` height
	str = ensureSize(str, rect.Width, rect.Height)

	lines := strings.Split(str, "\n")
	for i, line := range lines {
		var (
			frameLineIdx   = rect.Y + uint16(i)
			frameLineRunes = []rune(f.frame[frameLineIdx])

			// x position on the frameLine with skipped ansi codes
			x = getWithoutAnsi(int(rect.X), f.frame[frameLineIdx])

			// everything before the rect's x. its excluding, so x is free
			before = string(frameLineRunes[:x])
			// everything after x + width
			after = string(frameLineRunes[uint16(x)+rect.Width:])
		)

		f.frame[frameLineIdx] = before + line + after
	}
}

func (f FrameBuffer) View() string {
	return strings.Join(f.frame, "\n")
}
