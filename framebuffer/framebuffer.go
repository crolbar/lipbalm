package framebuffer

import (
	"github.com/crolbar/lipbalm"
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

func (f *FrameBuffer) Resize(width, height int) {
	f.width = uint16(width)
	f.height = uint16(height)
	f.Clear()
}

func (f *FrameBuffer) Clear() {
	var (
		line  = strings.Repeat(" ", int(f.width))
		frame = make([]string, f.height)
	)

	for i := range frame {
		frame[i] = line
	}

	f.frame = frame
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
//
// will add padding the string if its smaller than the rect
// alignments[0] position of horizontal padding
// alignments[1] position of vertical padding
func (f *FrameBuffer) RenderString(
	str string,
	rect layout.Rect,
	alignments ...lipbalm.Position,
) error {
	if rect.Width <= 0 || rect.Height <= 0 {
		return nil
	}

	// make sure that the string is expanded to `rect.Width` width and `rect.Height` height
	str = ensureSize(str, rect.Width, rect.Height, alignments...)

	lines := strings.Split(str, "\n")
	for i, line := range lines {
		if int(rect.Y)+i >= len(f.frame) {
			continue
		}

		var (
			frameLineIdx   = rect.Y + uint16(i)
			frameLineRunes = []rune(f.frame[frameLineIdx])

			// beforeX position on the frameLine with skipped ansi codes
			beforeX = getWithoutAnsi(int(rect.X), f.frame[frameLineIdx])
			afterX  = min(beforeX+int(rect.Width), len(frameLineRunes)-1)

			// everything before the rect's x. its excluding, so x is free
			before = string(frameLineRunes[:beforeX])
			// everything after x + width
			after = string(frameLineRunes[afterX:])
		)

		f.frame[frameLineIdx] = before + line + after
	}

	return nil
}

func (f FrameBuffer) View() string {
	return strings.Join(f.frame, "\n")
}
