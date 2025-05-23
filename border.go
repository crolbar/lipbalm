package lipbalm

import (
	"math"
	"strings"
)

type BorderTextPos int

const (
	btTop BorderTextPos = iota
	btLeft
	btRight
	btBottom
)

// ColorFg/ColorBg:
//
//	ansi color code (use Color/ColorRGB)
type BorderType struct {
	Top         string
	Bottom      string
	Left        string
	Right       string
	TopLeft     string
	TopRight    string
	BottomLeft  string
	BottomRight string
	ColorFg     string
	ColorBg     string
	Text        string // shown only if it can fit in the border, if it is longer it is ignored
	TextPos     BorderTextPos
	TextAlign   Position // only on horizontal border
}

type BorderOpts func(bt *BorderType)

func WithFgColor(code uint8) BorderOpts {
	return func(bt *BorderType) {
		bt.ColorFg = Color(code)
	}
}

func WithBgColor(code uint8) BorderOpts {
	return func(bt *BorderType) {
		bt.ColorBg = ColorBg(code)
	}
}

func WithFgColorRGB(R, G, B uint8) BorderOpts {
	return func(bt *BorderType) {
		bt.ColorFg = ColorRGB(R, G, B)
	}
}

func WithBgColorRGB(R, G, B uint8) BorderOpts {
	return func(bt *BorderType) {
		bt.ColorBg = ColorBgRGB(R, G, B)
	}
}

func WithText(text string, align ...Position) BorderOpts {
	return func(bt *BorderType) {
		withTextHelper(bt, text, btTop, align)
	}
}

func WithTextTop(text string, align ...Position) BorderOpts {
	return func(bt *BorderType) {
		withTextHelper(bt, text, btTop, align)
	}
}

func WithTextBottom(text string, align ...Position) BorderOpts {
	return func(bt *BorderType) {
		withTextHelper(bt, text, btBottom, align)
	}
}

func WithTextLeft(text string, align ...Position) BorderOpts {
	return func(bt *BorderType) {
		withTextHelper(bt, text, btLeft, align)
	}
}

func WithTextRight(text string, align ...Position) BorderOpts {
	return func(bt *BorderType) {
		withTextHelper(bt, text, btRight, align)
	}
}

func withTextHelper(bt *BorderType, text string, pos BorderTextPos, align []Position) {
	bt.Text = text
	bt.TextPos = pos
	if len(align) > 0 {
		bt.TextAlign = align[0]
	}
}

// color:
//
//	ansi color code (use Color/ColorRGB)
//	[0] for foreground, [1] for background
func NormalBorder(opts ...BorderOpts) BorderType {
	bt := BorderType{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "┌",
		TopRight:    "┐",
		BottomLeft:  "└",
		BottomRight: "┘",
	}

	for _, o := range opts {
		o(&bt)
	}

	return bt
}

func RoundedBorder(opts ...BorderOpts) BorderType {
	bt := BorderType{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "╰",
		BottomRight: "╯",
	}

	for _, o := range opts {
		o(&bt)
	}

	return bt
}

func BorderN(
	str string,
	disabled ...bool,
) string {
	return Border(NormalBorder(), str, disabled...)
}

func BorderNF(
	code uint8,
	str string,
	disabled ...bool,
) string {
	return Border(
		NormalBorder(
			WithFgColor(code),
		), str, disabled...)
}

func BorderR(
	str string,
	disabled ...bool,
) string {
	return Border(RoundedBorder(), str, disabled...)
}

func BorderRF(
	code uint8,
	str string,
	disabled ...bool,
) string {
	return Border(
		RoundedBorder(
			WithFgColor(code),
		), str, disabled...)
}

func Border(
	b BorderType,
	str string,
	disabled ...bool,
) string {
	str = MakeSquare(Right, str)

	var (
		top, right, bottom, left = getDisabled(disabled)

		lines    = strings.Split(str, "\n")
		width    = GetStringWidth(lines[0])
		numLines = len(lines)

		topBorder    = iff(left, b.TopLeft, "") + strings.Repeat(b.Top, width) + iff(right, b.TopRight, "")
		bottomBorder = iff(left, b.BottomLeft, "") + strings.Repeat(b.Bottom, width) + iff(right, b.BottomRight, "")

		textRunes = []rune(b.Text)
		textWidth = GetWidth(b.Text)

		sb strings.Builder

		lastLineIdx = numLines - 1

		fullColorLen = len(b.ColorFg) + len(b.ColorBg) +
			iff(b.ColorBg != "" || b.ColorFg != "", len(ansi_reset), 0)

		sbSize = len(str) +
			iff(top, len(topBorder)+1+fullColorLen, 0) + // top border
			iff(bottom, len(bottomBorder)+1+fullColorLen, 0) + // bottom border
			iff(left, numLines*(len(b.Left)+fullColorLen), 0) + // left border
			iff(right, numLines*(len(b.Right)+fullColorLen), 0) // right border

		writeToSb = writeStringToSb(&sb)

		applyColor = func() {
			if b.ColorBg != "" && b.ColorFg != "" {
				sb.WriteString(
					b.ColorFg[:len(b.ColorFg)-1] + ";" +
						b.ColorBg[2:])
			} else {
				doIf(b.ColorFg, writeToSb)
				doIf(b.ColorBg, writeToSb)
			}
		}
		resetColor = func() {
			doIfp(b.ColorFg != "" || b.ColorBg != "",
				ansi_reset, writeToSb)
		}
	)

	// text in bottom/top border & the text fit in the border
	if b.Text != "" && textWidth <= width {
		if b.TextPos == btTop {
			topBorder = embedTextIntoBorder(b, topBorder, textRunes, textWidth, left, right)
		}

		if b.TextPos == btBottom {
			bottomBorder = embedTextIntoBorder(b, bottomBorder, textRunes, textWidth, left, right)
		}
	}

	sb.Grow(sbSize)

	if top {
		applyColor()
		sb.WriteString(topBorder)
		resetColor()

		sb.WriteByte('\n')
	}

	for i, line := range lines {
		if left {
			if b.TextPos == btLeft && i < textWidth {
				sb.WriteRune(textRunes[i])
			} else {
				applyColor()
				sb.WriteString(b.Left)
				resetColor()
			}
		}

		sb.WriteString(line)

		if right {
			if b.TextPos == btRight && i < textWidth {
				sb.WriteRune(textRunes[i])
			} else {
				applyColor()
				sb.WriteString(b.Right)
				resetColor()
			}
		}

		if i == lastLineIdx {
			break
		}

		sb.WriteByte('\n')
	}

	if bottom {
		sb.WriteByte('\n')

		applyColor()
		sb.WriteString(bottomBorder)
		resetColor()
	}

	return sb.String()
}

// text into horizontal border top/bottom
func embedTextIntoBorder(
	b BorderType,
	border string,
	textRunes []rune,
	textWidth int,
	left bool,
	right bool,
) string {
	var (
		sb          strings.Builder
		borderRunes = []rune(border)
		align       = b.TextAlign

		off       = iff(left, 1, 0) + iff(right, 1, 0)
		textStart = int(math.Round(
			float64(len(borderRunes)-off)*float64(align) - // - the two optional borders left&right
				float64(textWidth)*float64(align)))

		textAnsiOff  = 0
		writeOutAnsi = func(tIdx int, textRune rune) {
			for textRune != 'm' {
				sb.WriteRune(textRune)
				textAnsiOff++
				textRune = textRunes[tIdx-textStart+textAnsiOff]
			}

			// write the 'm'
			sb.WriteRune(textRune)
		}
	)

	for i, r := range borderRunes {
		// left border corner
		if i == 0 && left {
			sb.WriteRune(r)
			continue
		}
		// right border corner
		if i == len(borderRunes)-1 && right {
			sb.WriteRune(r)
			continue
		}

		// text
		tIdx := i - iff(left, 1, 0)
		if tIdx >= textStart && tIdx-textStart < textWidth {
			// if start of tetx & border has color, reset it
			if tIdx == textStart && (b.ColorFg != "" || b.ColorBg != "") {
				sb.WriteString(ansi_reset)
			}

			textRune := textRunes[tIdx-textStart+textAnsiOff]

			// if text has ansi codes write them out first
			if textRune == rune(ansi_esc) {
				writeOutAnsi(tIdx, textRune)
				textAnsiOff++
				textRune = textRunes[tIdx-textStart+textAnsiOff]
			}

			sb.WriteRune(textRune)

			// if text has ansi codes at the end write them out
			if tIdx-textStart+textAnsiOff < len(textRunes)-1 {
				// next rune
				textRune = textRunes[tIdx-textStart+textAnsiOff+1]
				if textRune == rune(ansi_esc) {
					// increase because of the +1
					textAnsiOff++

					writeOutAnsi(tIdx, textRune)
				}
			}

			// set border color back
			{
				// if we are not an the end of the text
				if tIdx-textStart != textWidth-1 {
					continue
				}
				// if we don't have a right border & we are at the end of the border
				// we don't have to apply the border color again
				if i == len(borderRunes)-1 && !right {
					continue
				}

				// if we are at the end of the text and
				// if border has color make sure we don't reset it
				if b.ColorBg != "" && b.ColorFg != "" {
					sb.WriteString(
						b.ColorFg[:len(b.ColorFg)-1] + ";" +
							b.ColorBg[2:])
				} else {
					if b.ColorFg != "" {
						sb.WriteString(b.ColorFg)
					}
					if b.ColorBg != "" {
						sb.WriteString(b.ColorBg)
					}
				}
			}

			continue
		}

		// border char
		sb.WriteRune(r)
	}

	return sb.String()
}

func getDisabled(disabled []bool) (top bool, right bool, bottom bool, left bool) {
	switch len(disabled) {
	case 4:
		left = disabled[3]
		fallthrough
	case 3:
		bottom = disabled[2]
		fallthrough
	case 2:
		right = disabled[1]
		fallthrough
	case 1:
		top = disabled[0]
	}

	return !top, !right, !bottom, !left
}
