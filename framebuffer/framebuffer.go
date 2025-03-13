package framebuffer

import (
	"github.com/crolbar/lipbalm"
	"github.com/crolbar/lipbalm/layout"
	"strings"
)

type FrameBuffer struct {
	height uint16 // height of the frame
	width  uint16 // width of the frame
	frame  [][]cell
}

type colorMode uint8

const (
	noColor colorMode = iota
	fg256
	fgTC
	bg256
	bgTC
)

const ansi_reset string = "\x1b[0m"

type color struct {
	mode colorMode

	// 0 => 256 color / true color R
	// 1 => true color G
	// 2 => true color B
	vals [3]uint8
}

type cell struct {
	rune rune
	fg   color
	bg   color
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

func (f *FrameBuffer) RenderString(
	str string,
	rect layout.Rect,
	alignments ...lipbalm.Position,
) {
	if rect.Width <= 0 ||
		rect.Height <= 0 ||
		rect.Y >= f.height ||
		rect.X >= f.width ||
		rect.X+rect.Width > f.width {
		return
	}

	// make sure that the string is expanded to `rect.Width` width and `rect.Height` height
	str = ensureSize(str, rect.Width, rect.Height, alignments...)

	for i, line := range strings.Split(str, "\n") {
		frameLineIdx := rect.Y + uint16(i)

		if frameLineIdx >= f.height {
			break
		}

		copy(
			f.frame[frameLineIdx][rect.X:rect.X+rect.Width],
			convertLineToCells([]rune(line), rect.Width),
		)
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
		var (
			curBgColor color = emptyColor
			curFgColor color = emptyColor

			resetAppliedColor = func() {
				if curBgColor.mode != noColor || curFgColor.mode != noColor {
					sb.WriteString(ansi_reset)
				}
			}

			handleColorChange = func(c cell) {
				// if there was a prev color applied, reset it
				resetAppliedColor()

				curBgColor = c.bg
				curFgColor = c.fg

				writeColor(&sb, curFgColor, curBgColor)
			}
		)

		for _, cell := range row {
			// if there is a change in color
			if curBgColor != cell.bg || curFgColor != cell.fg {
				handleColorChange(cell)
			}

			sb.WriteRune(cell.rune)
		}

		// reset applied color if there is one
		resetAppliedColor()

		if i == len(f.frame)-1 {
			break
		}

		sb.WriteByte('\n')
	}

	return sb.String()
}
