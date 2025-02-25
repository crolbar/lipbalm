package framebuffer

import (
	"github.com/crolbar/lipbalm"
	"github.com/crolbar/lipbalm/layout"
	"strings"
)

type FrameBuffer struct {
	height uint16 // height of the frame
	width  uint16 // width of the frame
	frame  [][]rune
}

func NewFrameBuffer(width, height uint16) FrameBuffer {
	return FrameBuffer{
		height: height,
		width:  width,
		frame:  genBuffer(int(width), int(height)),
	}
}

func (f *FrameBuffer) Resize(width, height int) {
	f.width = uint16(width)
	f.height = uint16(height)
	f.frame = genBuffer(int(width), int(height))
}

func (f *FrameBuffer) Clear() {
	f.frame = genBuffer(int(f.width), int(f.height))
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
// alignments[0] alignment of string horizontally
// alignments[1] alignment of string vertically
func (f *FrameBuffer) RenderString(
	str string,
	rect layout.Rect,
	alignments ...lipbalm.Position,
) {
	if rect.Width <= 0 || rect.Height <= 0 {
		return
	}

	// make sure that the string is expanded to `rect.Width` width and `rect.Height` height
	str = ensureSize(str, rect.Width, rect.Height, alignments...)

	for i, line := range strings.Split(str, "\n") {
		if int(rect.Y)+i >= len(f.frame) {
			continue
		}

		var (
			frameLineIdx   = rect.Y + uint16(i)
			frameLineRunes = f.frame[frameLineIdx]

			// beforeX position on the frameLine with skipped ansi codes
			beforeX = getWithoutAnsi(int(rect.X), f.frame[frameLineIdx])
			afterX  = min(beforeX+int(rect.Width), len(frameLineRunes))

			// everything after x + width
			after = frameLineRunes[afterX:]

			line          = []rune(line)
			lineSize      = len(line)
			availableSize = afterX - beforeX
		)

		// grow the line buffer if needed
		if lineSize > availableSize {
			var (
				sizeNeeded = lineSize - availableSize
				newLine    = make([]rune, len(f.frame[frameLineIdx])+sizeNeeded)
			)

			copy(newLine, f.frame[frameLineIdx])
			f.frame[frameLineIdx] = newLine

			afterX = afterX + sizeNeeded
		}

		copy(f.frame[frameLineIdx][beforeX:afterX], line)
		copy(f.frame[frameLineIdx][afterX:], after)
	}
}

func (f FrameBuffer) View() string {
	var (
		sb   strings.Builder
		size = 0
	)
	for _, row := range f.frame {
		size += len(row) + 1
	}
	sb.Grow(size)

	for i, row := range f.frame {
		for _, r := range row {
			sb.WriteRune(r)
		}

		if i == len(f.frame)-1 {
			break
		}

		sb.WriteByte('\n')
	}

	return sb.String()
}
