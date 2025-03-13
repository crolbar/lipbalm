package framebuffer

import (
	"fmt"
	"os"
	"testing"

	"github.com/crolbar/lipbalm"
	"github.com/crolbar/lipbalm/assert"
	"github.com/crolbar/lipbalm/layout"
)

var l = layout.DefaultLayout()

func TestM(t *testing.T) {
	var (
		fb = NewFrameBuffer(23, 15)

		vsplits = l.Vercital().
			Constrains(
				layout.NewConstrain(layout.Percent, 50),
				layout.NewConstrain(layout.Percent, 50),
			).Split(fb.Size())

		hsplits1 = l.Horizontal().
				Constrains(
				layout.NewConstrain(layout.Percent, 50),
				layout.NewConstrain(layout.Percent, 50),
			).Split(vsplits[0])

		hsplits2 = l.Horizontal().
				Constrains(
				layout.NewConstrain(layout.Percent, 50),
				layout.NewConstrain(layout.Percent, 50),
			).Split(vsplits[1])
	)
	r := layout.Rect{
		X:      hsplits1[0].Width / 2,
		Y:      hsplits1[0].Height / 2,
		Width:  hsplits1[0].Width,
		Height: hsplits1[0].Height,
	}

	for i, s := range append(append(hsplits1, hsplits2...), r) {
		fb.RenderString(
			lipbalm.BorderNF(1,
				lipbalm.Expand(int(s.Height-2), int(s.Width-2),
					lipbalm.SetColor(lipbalm.ColorRGB(120, 0, 120),
						lipbalm.SetColor(lipbalm.ColorBgRGB(0, 170, 170),
							fmt.Sprintf("%d", i),
						)),
					lipbalm.Center, lipbalm.Center),
			),
			s,
		)
	}

	fmt.Println(fb.View())

	fmt.Println(fb.frame)
}

func TestSplit2Color(t *testing.T) {
	fb := NewFrameBuffer(100, 26)
	splitsA := make([]layout.Rect, 0)

	splits := l.Vercital().
		Constrains(
			layout.NewConstrain(layout.Percent, 100),
			layout.NewConstrain(layout.Min, 4),
		).Split(fb.Size())
	splitsA = append(splitsA, splits[1])

	hsplits := l.Horizontal().
		Constrains(
			layout.NewConstrain(layout.Length, 5),
			layout.NewConstrain(layout.Percent, 20),
			layout.NewConstrain(layout.Percent, 20),
			layout.NewConstrain(layout.Min, 5),
			layout.NewConstrain(layout.Min, 5),
			layout.NewConstrain(layout.Percent, 20),
			layout.NewConstrain(layout.Percent, 20),
			layout.NewConstrain(layout.Length, 5),
		).Split(splits[0])

	splitsA = append(splitsA, hsplits[0])
	splitsA = append(splitsA, hsplits[3:5]...)
	splitsA = append(splitsA, hsplits[7])

	for _, s := range []layout.Rect{
		hsplits[1], hsplits[2],
		hsplits[5], hsplits[6],
	} {
		splits = l.Vercital().
			Constrains(
				layout.NewConstrain(layout.Percent, 50),
				layout.NewConstrain(layout.Percent, 50),
			).Split(s)
		splitsA = append(splitsA, splits...)
	}

	for i, s := range splitsA {
		fb.RenderString(
			lipbalm.Border(lipbalm.NormalBorder(lipbalm.WithFgColor(uint8(i+100))),
				lipbalm.ExpandVertical(int(s.Height)-2, lipbalm.Center,
					lipbalm.ExpandHorizontal(int(s.Width)-2, lipbalm.Center,
						lipbalm.SetColor(lipbalm.Color(9),
							fmt.Sprintf("%d", i),
						),
					),
				)),
			s,
		)
	}

	frame := fb.View()

	// fmt.Println(frame)

	assert.Equal(t, getDump("splits2"), frame)
}

func TestSplits1(t *testing.T) {
	fb := NewFrameBuffer(100, 26)
	splitsA := make([]layout.Rect, 0)

	hsplits := l.Horizontal().
		Constrains(
			layout.NewConstrain(layout.Percent, 25),
			layout.NewConstrain(layout.Percent, 50),
			layout.NewConstrain(layout.Percent, 25),
		).Split(fb.Size())

	vsplits := l.Vercital().
		Constrains(
			layout.NewConstrain(layout.Percent, 50),
			layout.NewConstrain(layout.Percent, 50),
		).Split(hsplits[0])
	splitsA = append(splitsA, vsplits...)

	vsplits = l.Vercital().
		Constrains(
			layout.NewConstrain(layout.Percent, 50),
			layout.NewConstrain(layout.Percent, 50),
		).Split(hsplits[2])
	splitsA = append(splitsA, vsplits...)

	vsplits = l.Vercital().
		Constrains(
			layout.NewConstrain(layout.Length, 3),
			layout.NewConstrain(layout.Percent, 50),
			layout.NewConstrain(layout.Percent, 50),
			layout.NewConstrain(layout.Length, 3),
		).Split(hsplits[1])
	splitsA = append(splitsA, vsplits...)

	for _, s := range splitsA {
		fb.RenderString(
			lipbalm.Border(lipbalm.NormalBorder(),
				lipbalm.ExpandVertical(int(s.Height)-2, lipbalm.Bottom,
					lipbalm.ExpandHorizontal(int(s.Width)-2, lipbalm.Bottom, ""),
				)),
			s,
		)
	}

	frame := fb.View()

	assert.Equal(t, getDump("splits1"), frame)
}

func TestVertHalfFB(t *testing.T) {
	fb := NewFrameBuffer(50, 50)

	splits := l.Horizontal().
		Constrains(
			layout.NewConstrain(layout.Percent, 50),
			layout.NewConstrain(layout.Percent, 50),
		).Split(fb.Size())

	border := lipbalm.NormalBorder(lipbalm.WithFgColor(70))

	fb.RenderString(
		lipbalm.Border(border,
			lipbalm.ExpandVertical(int(splits[0].Height)-2, lipbalm.Bottom,
				lipbalm.ExpandHorizontal(int(splits[0].Width)-2, lipbalm.Bottom, ""),
			)),
		splits[0],
	)

	fb.RenderString(
		lipbalm.Border(border,
			lipbalm.ExpandVertical(int(splits[1].Height)-2, lipbalm.Bottom,
				lipbalm.ExpandHorizontal(int(splits[1].Width)-2, lipbalm.Bottom, ""),
			)),
		splits[1],
	)

	frame := fb.View()

	// fmt.Println(frame)

	assert.Equal(t,
		"┌──────────────────────────────────────────────────┐\n│\x1b[38;5;70m┌───────────────────────┐\x1b[0m\x1b[38;5;70m┌───────────────────────┐\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m└───────────────────────┘\x1b[0m\x1b[38;5;70m└───────────────────────┘\x1b[0m│\n└──────────────────────────────────────────────────┘",
		lipbalm.Border(lipbalm.NormalBorder(), frame),
	)
}

func getDump(name string) string {
	data, err := os.ReadFile(fmt.Sprintf("%s.dump", name))
	if err != nil {
		panic(err)
	}
	return string(data)
}
