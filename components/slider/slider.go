package slider

// slider bar / progress bar
//
// ┌test─┐
// │     │
// │     │
// │     │
// │     │
// │     │
// │█████│
// │█████│
// │█████│
// │█████│
// │█████│
// └─────┘
// ┌0.392──────────┐
// │█████          │
// │█████          │
// └───────────────┘

// Init ex:

// s := NewSlider("test", 5, 10,
// 	  WithBorder(),
// 	  WithReversed(),
// 	  WithVertical(),
// 	  WithFocus(),
// 	  WithProgressColor(lb.Color(1)),

// Update:
//
// switch msg := msg.(type) {
// case tea.KeyMsg:
//     s.Update(msg.String())
// .....

// Mouse Update:
// rect is lipbalm.layout.Rect
// the bubbletea program should be started with mouse support:
//   tea.NewProgram(model{}, tea.WithMouseCellMotion())
//
// switch msg := msg.(type) {
// case tea.MouseMsg:
//     m.s.UpdateMouseClick(msg.String(), msg.X, msg.Y, rect)
// ...

// View: s.View()

import (
	"fmt"
	"strings"

	lb "github.com/crolbar/lipbalm"
	lbl "github.com/crolbar/lipbalm/layout"
)

type Slider struct {
	// title used in the border
	// set to _RATIO to display the ratio in the border
	Title string

	// called when Update updates the progress
	Trigger         func(any) error
	TriggerArgument any

	// 0 for 0% and 255 for 100%
	Progress     uint8
	ProgressRate uint8

	// height, width and Rect count in the border
	// uses Height & Width if both non zero else uses Rect for size
	Height int
	Width  int
	Rect   lbl.Rect

	// make 0% be the other side
	Reverse  bool
	Vertical bool

	HasBorder bool
	Border    lb.BorderType

	NoTopBorder    bool
	NoRightBorder  bool
	NoBottomBorder bool
	NoLeftBorder   bool

	// unset these if don't want
	FocusedColor   string
	UnfocusedColor string

	// used when no border & no focused & unfocused colors
	// or when we have border
	ProgressColor string

	Focused bool
}

var DecreaseKeys []string = []string{
	"left", "h", "ctrl+b", "j", "down",
}

var IncreaseKeys []string = []string{
	"right", "l", "ctrl+f", "k", "up",
}

var MouseKeys []string = []string{
	"left press", "left release", "left motion",
}

type Opts func(*Slider)

// the first optional item would be used as the border
func WithBorder(border ...lb.BorderType) Opts {
	return func(s *Slider) {
		s.HasBorder = true
		if len(border) > 0 {
			s.Border = border[0]
		}
	}
}

// first item in arg would be used as the argument to the trigger
func WithTrigger(t func(any) error, arg ...any) Opts {
	return func(s *Slider) {
		s.Trigger = t
		if len(arg) > 0 {
			s.TriggerArgument = arg[0]
		}
	}
}

func WithInitProgress(progress uint8) Opts {
	return func(s *Slider) {
		s.Progress = progress
	}
}

func WithFocus() Opts {
	return func(s *Slider) {
		s.Focused = true
	}
}

func WithReversed() Opts {
	return func(s *Slider) {
		s.Reverse = true
	}
}

func WithVertical() Opts {
	return func(s *Slider) {
		s.Vertical = true
	}
}

func WithFocusedColor(color string) Opts {
	return func(s *Slider) {
		s.FocusedColor = color
	}
}

func WithUnfocusedColor(color string) Opts {
	return func(s *Slider) {
		s.UnfocusedColor = color
	}
}

func WithProgressColor(color string) Opts {
	return func(s *Slider) {
		s.ProgressColor = color
	}
}

func WithNoTopBorder() Opts {
	return func(s *Slider) {
		s.NoTopBorder = true
	}
}

func WithNoRightBorder() Opts {
	return func(s *Slider) {
		s.NoRightBorder = true
	}
}

func WithNoBottomBorder() Opts {
	return func(s *Slider) {
		s.NoBottomBorder = true
	}
}

func WithNoLeftBorder() Opts {
	return func(s *Slider) {
		s.NoLeftBorder = true
	}
}

var DefaultSlider Slider = Slider{
	FocusedColor:   lb.Color(54),
	ProgressColor:  lb.Color(54),
	UnfocusedColor: lb.Color(15),
	Progress:       0,
	ProgressRate:   5,
}

func Init(
	title string,
	opts ...Opts,
) *Slider {
	s := DefaultSlider
	s.Title = title
	s.Border = lb.NormalBorder(lb.WithTextTop(title, lb.Left))

	for _, o := range opts {
		o(&s)
	}

	return &s
}

func (s *Slider) Update(key string) (change bool, err error) {
	if !s.Focused {
		return
	}
	var (
		incKeys = IncreaseKeys
		decKeys = DecreaseKeys

		onChange = func() {
			if s.Trigger != nil {
				s.Trigger(s.TriggerArgument)
			}
			change = true
		}
	)

	if s.Vertical != s.Reverse {
		incKeys = DecreaseKeys
		decKeys = IncreaseKeys
	}

	switch {
	case matchKey(key, decKeys):
		s.DecreaseProgress()
		onChange()
	case matchKey(key, incKeys):
		s.IncreaseProgress()
		onChange()
	}
	return
}

func (s *Slider) UpdateMouseClick(
	key string,
	mx int,
	my int,
	rect lbl.Rect,
) (change bool) {
	if !s.Focused {
		return
	}

	hasCollision, ratioInRect := s.CheckMouseCollision(mx, my, rect)
	if !hasCollision {
		return
	}

	switch {
	case matchKey(key, MouseKeys):
		if s.Reverse {
			s.Progress = 255 - ratioInRect
		} else {
			s.Progress = ratioInRect
		}

		if s.Trigger != nil {
			s.Trigger(s.TriggerArgument)
		}
		change = true
	}
	return
}

func (s *Slider) CheckMouseCollision(
	mx int,
	my int,
	rect lbl.Rect,
) (bool, uint8) {
	var (
		rx = int(rect.X)
		ry = int(rect.Y)
		w  = int(rect.Width)
		h  = int(rect.Height)
	)

	if mx >= rx && mx <= rx+w-1 && my >= ry && my <= ry+h-1 {
		if s.HasBorder {
			// no click on border
			if mx == rx || my == ry || mx == rx+w-1 || my == ry+h-1 {
				return false, 0
			}

			// if click in start
			if !s.Vertical && mx == rx+1 {
				return true, 0
			}
			if s.Vertical && my == ry+1 {
				return true, 0
			}

			// skip border for ratio calc
			h -= 2
			w -= 2
		}

		var ratio uint8
		if s.Vertical {
			ratio = uint8(255.0 * (float64(my-ry) / float64(h)))
		} else {
			ratio = uint8(255.0 * (float64(mx-rx) / float64(w)))
		}

		return true, ratio
	}

	return false, 0
}

func (s *Slider) IncreaseProgress() {
	if int(s.Progress)+int(s.ProgressRate) >= 255 {
		s.Progress = 255
		return
	}

	s.Progress = s.Progress + s.ProgressRate
}

func (s *Slider) DecreaseProgress() {
	if int(s.Progress)-int(s.ProgressRate) <= 0 {
		s.Progress = 0
		return
	}

	s.Progress = s.Progress - s.ProgressRate
}

func (s Slider) View() string {
	var (
		sb  strings.Builder
		end = s.GetHeight()
		w   = s.GetWidth()

		progress = s.GetRatio()

		fullBar  string
		emptyBar string
	)

	if s.HasBorder {
		end -= 2
		w -= 2
	}

	if s.Vertical {
		fullBar = strings.Repeat("█", w)
		emptyBar = strings.Repeat(" ", w)
	} else {
		progWidth := float64(w) * progress
		fullBar = strings.Repeat("█", int(progWidth))
		emptyBar = strings.Repeat(" ", w-int(progWidth))
	}

	if !s.HasBorder {
		if s.Focused && s.FocusedColor != "" {
			fullBar = lb.SetColor(s.FocusedColor, fullBar)
		}

		if !s.Focused && s.UnfocusedColor != "" {
			fullBar = lb.SetColor(s.UnfocusedColor, fullBar)
		}

		// if we don't focus and unfocused colors set to progress color
		if s.UnfocusedColor == "" && s.FocusedColor == "" && s.ProgressColor != "" {
			fullBar = lb.SetColor(s.ProgressColor, fullBar)
		}
	}

	if s.HasBorder {
		if s.ProgressColor != "" {
			fullBar = lb.SetColor(s.ProgressColor, fullBar)
		}
	}

	for i := range end {
		if s.Vertical {
			if s.Reverse {
				if i < int((1-progress)*float64(end)) {
					sb.WriteString(emptyBar)
				} else {
					sb.WriteString(fullBar)
				}
			} else {
				if i < int(progress*float64(end)) {
					sb.WriteString(fullBar)
				} else {
					sb.WriteString(emptyBar)
				}
			}
		} else {
			if s.Reverse {
				sb.WriteString(emptyBar)
				sb.WriteString(fullBar)
			} else {
				sb.WriteString(fullBar)
				sb.WriteString(emptyBar)
			}
		}

		if i == end-1 {
			break
		}

		sb.WriteByte('\n')
	}

	out := sb.String()

	if s.HasBorder {
		if s.Title == "_RATIO" {
			s.Border.Text = fmt.Sprintf("%.3f", s.GetRatio())
		} else {
			s.Border.Text = s.Title
		}

		if s.Focused && s.FocusedColor != "" {
			s.Border.ColorFg = s.FocusedColor
		}
		if !s.Focused && s.UnfocusedColor != "" {
			s.Border.ColorFg = s.UnfocusedColor
		}

		out = lb.Border(s.Border, out,
			s.NoTopBorder,
			s.NoRightBorder,
			s.NoBottomBorder,
			s.NoLeftBorder)
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

func (s *Slider) GetRatio() float64 {
	return float64(s.Progress) / 255.0
}

func (s *Slider) SetRatio(ratio float64) {
	s.Progress = uint8(255.0 * ratio)
}

func (s *Slider) GetRect() lbl.Rect {
	return s.Rect
}

func (s *Slider) GetHeight() int {
	if s.Height == 0 {
		return int(s.Rect.Height)
	}

	return s.Height
}

func (s *Slider) GetWidth() int {
	if s.Width == 0 {
		return int(s.Rect.Width)
	}

	return s.Width
}

func (s *Slider) HasFocus() bool {
	return s.Focused
}

func (s *Slider) FocusToggle() {
	s.Focused = !s.Focused
}

func (s *Slider) Focus() {
	s.Focused = true
}

func (s *Slider) DeFocus() {
	s.Focused = false
}

func (s *Slider) SetTrigger(t func(any) error) {
	s.Trigger = t
}

func (s *Slider) GetTrigger() func(any) error {
	return s.Trigger
}

func (s *Slider) SetTriggerArgument(a any) {
	s.TriggerArgument = a
}
