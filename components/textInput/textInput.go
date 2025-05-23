package textInput

// simple text input ui component
// just init, call ti.Update(key) in Update and optionally render it

// example init
// ti := NewTextInput("Test", 40, 5,
// 	WithBorder(),
// 	WithCursorColor(lb.ColorBg(2)),
// 	WithInitText("hello"),
// )

// Update:
//
// switch msg := msg.(type) {
// case tea.KeyMsg:
//      v.ti.Update(msg.String())
// .....

// and View returns just a width by height box that optionally has a border
// with the text that's aligned based on VAlignment and HAlignment

import (
	"strconv"
	"strings"

	lb "github.com/crolbar/lipbalm"
	lbl "github.com/crolbar/lipbalm/layout"
)

type TextInput struct {
	Title string

	// called when Update updates the Text
	Trigger         func(any) error
	TriggerArgument any

	// height, width and Rect count in the border
	// uses Height & Width if both non zero else uses Rect for size
	Height int
	Width  int
	Rect   lbl.Rect

	Text      *strings.Builder
	CursorPos int

	HasBorder bool
	Border    lb.BorderType

	VAlignment lb.Position
	HAlignment lb.Position

	TextColor string

	// sets color of border when hasborder
	// else color the text
	FocusedColor string

	// Cursor at the end is a whitespace an
	// fg color wont be visable
	CursorColor string

	NoTopBorder    bool
	NoRightBorder  bool
	NoBottomBorder bool
	NoLeftBorder   bool

	Focused bool
}

type Opts func(*TextInput)

// the first optional item would be used as the border
func WithBorder(border ...lb.BorderType) Opts {
	return func(ti *TextInput) {
		ti.HasBorder = true
		if len(border) > 0 {
			ti.Border = border[0]
		}
	}
}

// first item in arg would be used as the argument to the trigger
func WithTrigger(t func(any) error, arg ...any) Opts {
	return func(ti *TextInput) {
		ti.Trigger = t
		if len(arg) > 0 {
			ti.TriggerArgument = arg[0]
		}
	}
}

func WithFocus() Opts {
	return func(ti *TextInput) {
		ti.Focused = true
	}
}

func WithInitText(text string) Opts {
	return func(ti *TextInput) {
		ti.Text.WriteString(text)
		ti.CursorPos = len(text)
	}
}

func WithFocusedColor(color string) Opts {
	return func(ti *TextInput) {
		ti.FocusedColor = color
	}
}

func WithTextColor(color string) Opts {
	return func(ti *TextInput) {
		ti.TextColor = color
	}
}

func WithCursorColor(color string) Opts {
	return func(ti *TextInput) {
		ti.CursorColor = color
	}
}

func WithNoTopBorder() Opts {
	return func(ti *TextInput) {
		ti.NoTopBorder = true
	}
}

func WithNoRightBorder() Opts {
	return func(ti *TextInput) {
		ti.NoRightBorder = true
	}
}

func WithNoBottomBorder() Opts {
	return func(ti *TextInput) {
		ti.NoBottomBorder = true
	}
}

func WithNoLeftBorder() Opts {
	return func(ti *TextInput) {
		ti.NoLeftBorder = true
	}
}

func WithVAlignment(alignment lb.Position) Opts {
	return func(ti *TextInput) {
		ti.VAlignment = alignment
	}
}

func WithHAlignment(alignment lb.Position) Opts {
	return func(ti *TextInput) {
		ti.HAlignment = alignment
	}
}

var DefaultTextInput = TextInput{
	FocusedColor: lb.Color(54),
	VAlignment:   lb.Top,
	HAlignment:   lb.Left,
	CursorColor:  lb.ColorBg(1),
}

func Init(
	title string,
	opts ...Opts,
) *TextInput {
	ti := DefaultTextInput
	ti.Text = &strings.Builder{}
	ti.Title = title
	ti.Border = lb.NormalBorder(lb.WithTextTop(title, lb.Left))

	for _, o := range opts {
		o(&ti)
	}

	return &ti
}

// pressed key from the tea.KeyMsg.String() type
func (ti *TextInput) Update(key string) (change bool, err error) {
	if !ti.Focused {
		return
	}

	onChange := func() {
		if ti.Trigger != nil {
			ti.Trigger(ti.TriggerArgument)
		}
		change = true
	}

	// char / symbol
	if len(key) == 1 {
		err = ti.InsertText(rune(key[0]))
		onChange()
		return
	}

	switch key {
	case "left", "ctrl+b":
		ti.MoveCursorLeft()
	case "right", "ctrl+f":
		ti.MoveCursorRight()
	case "ctrl+left", "alt+b":
		ti.MoveCursorLeftWord()
	case "ctrl+right", "alt+f":
		ti.MoveCursorRightWord()
	case "backspace":
		_, err = ti.DeleteBeforeCursor()
		onChange()
	case "delete":
		_, err = ti.DeleteAfterCursor()
		onChange()
	case "alt+delete":
		_, err = ti.DeleteWordAfterCursor()
		onChange()
	case "ctrl+backspace", "ctrl+w", "ctrl+h":
		_, err = ti.DeleteWordBeforeCursor()
		onChange()
	}

	return
}

func (ti *TextInput) InsertText(ch rune) (err error) {
	// at end of buffer
	if ti.CursorPos == ti.Text.Len() {
		_, err = ti.Text.WriteRune(ch)
	} else {
		var (
			text           = ti.Text.String()
			pos            = ti.CursorPos
			preCursorText  = text[:pos]
			postCursorText = text[pos:]
		)

		ti.Text.Reset()
		_, err = ti.Text.WriteString(preCursorText + string(ch) + postCursorText)
	}

	ti.MoveCursorRight()
	return
}

func (ti *TextInput) DeleteBeforeCursor() (out rune, err error) {
	pos := ti.CursorPos - 1

	if pos == -1 {
		return
	}

	text := ti.Text.String()
	out = rune(text[pos])

	text = text[:pos] + text[min(pos+1, len(text)):]

	ti.Text.Reset()
	_, err = ti.Text.WriteString(text)
	ti.MoveCursorLeft()

	return
}

func (ti *TextInput) DeleteAfterCursor() (out rune, err error) {
	pos := ti.CursorPos

	if pos == ti.Text.Len() {
		return
	}

	text := ti.Text.String()
	out = rune(text[pos])

	text = text[:pos] + text[min(pos+1, len(text)):]

	ti.Text.Reset()
	_, err = ti.Text.WriteString(text)

	return
}

func (ti *TextInput) DeleteWordBeforeCursor() (out string, err error) {
	s := ti.Text.String()
	pos := ti.CursorPos

	// idx after the whitespace behind our cursor
	whiteSpaceIdx := getPrevWhitespaceIdx(s, pos) + 1

	out = s[whiteSpaceIdx:pos]
	s = s[:whiteSpaceIdx] + s[pos:]

	ti.Text.Reset()
	_, err = ti.Text.WriteString(s)
	ti.MoveCursorTo(whiteSpaceIdx)

	return
}

func (ti *TextInput) DeleteWordAfterCursor() (out string, err error) {
	s := ti.Text.String()
	pos := ti.CursorPos

	// idx after the whitespace after our cursor
	whiteSpaceIdx := getNextWhitespaceIdx(s, pos)

	out = s[pos:whiteSpaceIdx]
	s = s[:pos] + s[whiteSpaceIdx:]

	ti.Text.Reset()
	_, err = ti.Text.WriteString(s)

	return
}

func getPrevWhitespaceIdx(s string, pos int) int {
	for i := pos - 1; i >= 0; i-- {
		if i >= len(s) {
			continue
		}

		// if the char right behind the cursor is a whitespace
		if i == pos-1 && s[i] == ' ' {
			continue
		}

		if s[i] == ' ' {
			return i
		}
	}
	return -1
}

func getNextWhitespaceIdx(s string, pos int) int {
	for i := pos + 1; i < len(s); i++ {
		if i >= len(s) {
			continue
		}

		// if the char right after the cursor is a whitespace
		if i == pos+1 && s[i] == ' ' {
			continue
		}

		if s[i] == ' ' {
			return i
		}
	}
	return len(s)
}

func (ti *TextInput) MoveCursorLeft() {
	ti.CursorPos = max(ti.CursorPos-1, 0)
}

func (ti *TextInput) MoveCursorLeftWord() {
	ti.CursorPos = max(getPrevWhitespaceIdx(ti.GetText(), ti.CursorPos)+1, 0)
}

func (ti *TextInput) MoveCursorTo(n int) {
	ti.CursorPos = min(ti.Text.Len(), max(0, n))
}

func (ti *TextInput) MoveCursorRight() {
	ti.CursorPos = min(ti.CursorPos+1, ti.Text.Len())
}

func (ti *TextInput) MoveCursorRightWord() {
	ti.CursorPos = min(ti.Text.Len(), max(0, getNextWhitespaceIdx(ti.GetText(), ti.CursorPos)))
}

func (ti TextInput) View() string {
	var (
		text           = ti.Text.String()
		pos            = ti.CursorPos
		preCursorText  = text[:pos]
		postCursorText = text[min(pos+1, len(text)):]
		cursorChar     = " "

		h = ti.GetHeight()
		w = ti.GetWidth()
	)

	// if cursor is on char in text
	if pos < len(text) {
		cursorChar = string(text[pos])
	}

	if ti.Focused {
		cursorChar = lb.SetColor(ti.CursorColor, cursorChar)
	}

	text = preCursorText + cursorChar + postCursorText

	if len(ti.TextColor) > 0 {
		text = lb.SetColor(ti.TextColor, text)
	}

	if !ti.HasBorder && ti.Focused {
		text = lb.SetColor(ti.FocusedColor, text)
	}

	if ti.HasBorder {
		h -= 2
		w -= 2
	}

	out := lb.Expand(h, w,
		text,
		ti.VAlignment, ti.HAlignment,
	)

	if ti.HasBorder {
		if ti.Focused {
			ti.Border.ColorFg = ti.FocusedColor
		}

		out = lb.Border(ti.Border, out,
			ti.NoTopBorder,
			ti.NoRightBorder,
			ti.NoBottomBorder,
			ti.NoLeftBorder)
	}

	return out
}

func (ti *TextInput) SetText(text string) {
	ti.Text.Reset()
	ti.Text.WriteString(text)
	ti.CursorPos = max(0, min(ti.Text.Len(), ti.CursorPos))
}

func (ti *TextInput) GetRect() lbl.Rect {
	return ti.Rect
}

func (ti *TextInput) SetRect(r lbl.Rect) {
	ti.Rect = r
}

func (ti *TextInput) GetHeight() int {
	if ti.Height == 0 {
		return int(ti.Rect.Height)
	}

	return ti.Height
}

func (ti *TextInput) GetWidth() int {
	if ti.Width == 0 {
		return int(ti.Rect.Width)
	}

	return ti.Width
}

func (ti *TextInput) HasFocus() bool {
	return ti.Focused
}

func (ti *TextInput) FocusToggle() {
	ti.Focused = !ti.Focused
}

func (ti *TextInput) Focus() {
	ti.Focused = true
}

func (ti *TextInput) DeFocus() {
	ti.Focused = false
}

func (ti *TextInput) SetTrigger(t func(any) error) {
	ti.Trigger = t
}

func (ti *TextInput) GetTrigger() func(any) error {
	return ti.Trigger
}

func (ti *TextInput) SetTriggerArgument(a any) {
	ti.TriggerArgument = a
}

func (ti *TextInput) GetTriggerArgument() any {
	return ti.TriggerArgument
}

func (ti *TextInput) GetText() string {
	return ti.Text.String()
}

func (ti *TextInput) GetTextAsInt() (int, error) {
	num, err := strconv.ParseInt(ti.Text.String(), 10, 32)
	if err != nil {
		return 0, err
	}
	return int(num), nil
}

func (ti *TextInput) GetTextAsFloat() (float64, error) {
	num, err := strconv.ParseFloat(ti.Text.String(), 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}
