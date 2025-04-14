package button

import (
	lb "github.com/crolbar/lipbalm"
	lbl "github.com/crolbar/lipbalm/layout"
)

type Button struct {
	// title used in the border
	Title string

	// called when Update updates Pressed
	Trigger         func(any) error
	TriggerArgument any

	// height, width and Rect count in the border
	// uses Height & Width if both non zero else uses Rect for size
	Width  int
	Height int
	Rect   lbl.Rect

	HasBorder bool
	Border    lb.BorderType

	// text in the button
	Text        string
	PressedText string

	// alignments of the text
	VAlignment lb.Position
	HAlignment lb.Position

	// sets color of border when HasBorder
	// else sets box color
	FocusedColor string

	Pressed bool

	PressedFgColor       string // fg text color on pressed
	PressedBgColor       string // bg text color on pressed
	PressedBorderFgColor string // fg border color on pressed
	PressedBorderBgColor string // bg border coler on pressed

	NoTopBorder    bool
	NoRightBorder  bool
	NoBottomBorder bool
	NoLeftBorder   bool

	Focused bool
}

var PressKeys []string = []string{" "}

type Opts func(*Button)

func WithBorder(border ...lb.BorderType) Opts {
	return func(b *Button) {
		b.HasBorder = true
		if len(border) > 0 {
			b.Border = border[0]
		}
	}
}

func WithInitState(pressed bool) Opts {
	return func(b *Button) {
		b.Pressed = pressed
	}
}

func WithFocus() Opts {
	return func(b *Button) {
		b.Focused = true
	}
}

func WithText(text string) Opts {
	return func(b *Button) {
		b.Text = text
	}
}

func WithPressedText(text string) Opts {
	return func(b *Button) {
		b.PressedText = text
	}
}

func WithFocusedColor(color string) Opts {
	return func(b *Button) {
		b.FocusedColor = color
	}
}

func WithPressedFgColor(color string) Opts {
	return func(b *Button) {
		b.PressedFgColor = color
	}
}

func WithPressedBgColor(color string) Opts {
	return func(b *Button) {
		b.PressedBgColor = color
	}
}

func WithPressedBorderFgColor(color string) Opts {
	return func(b *Button) {
		b.PressedBorderFgColor = color
	}
}

func WithPressedBorderBgColor(color string) Opts {
	return func(b *Button) {
		b.PressedBorderBgColor = color
	}
}

func WithNoTopBorder() Opts {
	return func(b *Button) {
		b.NoTopBorder = true
	}
}

func WithNoRightBorder() Opts {
	return func(b *Button) {
		b.NoRightBorder = true
	}
}

func WithNoBottomBorder() Opts {
	return func(b *Button) {
		b.NoBottomBorder = true
	}
}

func WithNoLeftBorder() Opts {
	return func(b *Button) {
		b.NoLeftBorder = true
	}
}

func WithVAlignment(alignment lb.Position) Opts {
	return func(b *Button) {
		b.VAlignment = alignment
	}
}

func WithHAlignment(alignment lb.Position) Opts {
	return func(b *Button) {
		b.HAlignment = alignment
	}
}

var DefaultButton = Button{
	FocusedColor:   lb.Color(54),
	PressedBgColor: lb.ColorBg(54),
	VAlignment:     lb.Center,
	HAlignment:     lb.Center,
}

func Init(
	title string,
	opts ...Opts,
) Button {
	b := DefaultButton
	b.Title = title
	b.Border = lb.NormalBorder(lb.WithTextTop(title, lb.Left))

	for _, o := range opts {
		o(&b)
	}

	return b
}

func (b *Button) Update(key string) (change bool, err error) {
	if !b.Focused {
		return
	}

	switch {
	case matchKey(key, PressKeys):
		b.Press()

		if b.Trigger != nil {
			err = b.Trigger(b.TriggerArgument)
		}
		change = true
	}
	return
}

func (b Button) View() string {
	var (
		text   = b.Text
		border = b.Border

		h = b.GetHeight()
		w = b.GetWidth()
	)

	if b.Pressed {
		if has(b.PressedText) {
			text = b.PressedText
		}

		if has(b.PressedBorderFgColor) {
			border.ColorFg = b.PressedBorderFgColor
		}

		if has(b.PressedBorderBgColor) {
			border.ColorBg = b.PressedBorderBgColor
		}
	}

	// h and w are used for expantion of the internal box in the border
	// since width & height include the border exclude in from the box exapantion
	if b.HasBorder {
		h -= 2
		w -= 2
	}

	out := lb.Expand(h, w, text, b.VAlignment, b.HAlignment)

	if !b.HasBorder && b.Focused && has(b.FocusedColor) {
		out = lb.SetColor(b.FocusedColor, out)
	}

	if b.Pressed && has(b.PressedBgColor) {
		out = lb.SetColor(b.PressedBgColor, out)
	}

	if b.Pressed && has(b.PressedFgColor) {
		out = lb.SetColor(b.PressedFgColor, out)
	}

	if b.HasBorder {
		if b.Focused && has(b.FocusedColor) {
			border.ColorFg = b.FocusedColor
		}

		out = lb.Border(border, out,
			b.NoTopBorder,
			b.NoRightBorder,
			b.NoBottomBorder,
			b.NoLeftBorder)
	}

	return out
}

func matchKey(key string, keys []string) bool {
	for _, k := range keys {
		if key == k {
			return true
		}
	}
	return false
}

func has(s string) bool {
	return s != ""
}

func (b *Button) GetRect() lbl.Rect {
	return b.Rect
}

func (b *Button) GetHeight() int {
	if b.Height == 0 {
		return int(b.Rect.Height)
	}

	return b.Height
}

func (b *Button) GetWidth() int {
	if b.Width == 0 {
		return int(b.Rect.Width)
	}

	return b.Width
}

func (b *Button) Press() {
	b.Pressed = true
}

func (b *Button) PressToggle() {
	b.Pressed = !b.Pressed
}

func (b *Button) Depress() {
	b.Pressed = false
}

func (b *Button) IsPressed() bool {
	return b.Pressed
}

func (b *Button) HasFocus() bool {
	return b.Focused
}

func (b *Button) FocusToggle() {
	b.Focused = !b.Focused
}

func (b *Button) Focus() {
	b.Focused = true
}

func (b *Button) DeFocus() {
	b.Focused = false
}

func (b *Button) SetTrigger(t func(any) error) {
	b.Trigger = t
}

func (b *Button) GetTrigger() func(any) error {
	return b.Trigger
}

func (b *Button) SetTriggerArgument(a any) {
	b.TriggerArgument = a
}
