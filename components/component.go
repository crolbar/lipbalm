package component

import lbl "github.com/crolbar/lipbalm/layout"

type Component interface {
	View() string
	Update(key string) (change bool, err error)

	SetTrigger(func(any) error)
	GetTrigger() func(any) error
	SetTriggerArgument(any)

	HasFocus() bool
	Focus()
	DeFocus()

	GetRect() lbl.Rect
	SetRect(lbl.Rect)
}
