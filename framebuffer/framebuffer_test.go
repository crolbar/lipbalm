package framebuffer

import (
	"fmt"
	lb "github.com/crolbar/lipbalm"
	"github.com/crolbar/lipbalm/assert"
	"github.com/crolbar/lipbalm/layout"
	"os"
	"testing"
)

var l = layout.DefaultLayout()

func TestOverlay(t *testing.T) {
	var (
		fb = NewFrameBuffer(50, 27)

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

	w := uint16(float32(hsplits1[0].Width) * 0.8)
	h := uint16(float32(hsplits1[0].Height) * 0.8)
	r := layout.Rect{
		X:      hsplits1[0].Width - w/2,
		Y:      hsplits1[0].Height - h/2,
		Width:  w,
		Height: h,
	}

	for i, s := range append(append(hsplits1, hsplits2...), r) {
		fb.RenderString(
			lb.Border(lb.NormalBorder(lb.WithFgColor(1), lb.WithBgColor(101)),
				lb.Expand(int(s.Height-2), int(s.Width-2),
					lb.SetColor(lb.ColorRGB(120, 0, 120),
						lb.SetColor(lb.ColorBgRGB(0, 170, 170),
							fmt.Sprintf("%d", i),
						)),
					lb.Center, lb.Center),
			),
			s,
		)
	}

	frame := fb.View()
	// fmt.Println(frame)

	assert.Equal(t,
		"\x1b[38;5;1;48;5;101m┌───────────────────────┐┌───────────────────────┐\x1b[0m\n\x1b[38;5;1;48;5;101m│\x1b[0m                       \x1b[38;5;1;48;5;101m││\x1b[0m                       \x1b[38;5;1;48;5;101m│\x1b[0m\n\x1b[38;5;1;48;5;101m│\x1b[0m                       \x1b[38;5;1;48;5;101m││\x1b[0m                       \x1b[38;5;1;48;5;101m│\x1b[0m\n\x1b[38;5;1;48;5;101m│\x1b[0m                       \x1b[38;5;1;48;5;101m││\x1b[0m                       \x1b[38;5;1;48;5;101m│\x1b[0m\n\x1b[38;5;1;48;5;101m│\x1b[0m                       \x1b[38;5;1;48;5;101m││\x1b[0m                       \x1b[38;5;1;48;5;101m│\x1b[0m\n\x1b[38;5;1;48;5;101m│\x1b[0m                       \x1b[38;5;1;48;5;101m││\x1b[0m                       \x1b[38;5;1;48;5;101m│\x1b[0m\n\x1b[38;5;1;48;5;101m│\x1b[0m           \x1b[38;2;120;0;120;48;2;0;170;170m0\x1b[0m           \x1b[38;5;1;48;5;101m││\x1b[0m           \x1b[38;2;120;0;120;48;2;0;170;170m1\x1b[0m           \x1b[38;5;1;48;5;101m│\x1b[0m\n\x1b[38;5;1;48;5;101m│\x1b[0m                       \x1b[38;5;1;48;5;101m││\x1b[0m                       \x1b[38;5;1;48;5;101m│\x1b[0m\n\x1b[38;5;1;48;5;101m│\x1b[0m              \x1b[38;5;1;48;5;101m┌──────────────────┐\x1b[0m              \x1b[38;5;1;48;5;101m│\x1b[0m\n\x1b[38;5;1;48;5;101m│\x1b[0m              \x1b[38;5;1;48;5;101m│\x1b[0m                  \x1b[38;5;1;48;5;101m│\x1b[0m              \x1b[38;5;1;48;5;101m│\x1b[0m\n\x1b[38;5;1;48;5;101m│\x1b[0m              \x1b[38;5;1;48;5;101m│\x1b[0m                  \x1b[38;5;1;48;5;101m│\x1b[0m              \x1b[38;5;1;48;5;101m│\x1b[0m\n\x1b[38;5;1;48;5;101m│\x1b[0m              \x1b[38;5;1;48;5;101m│\x1b[0m                  \x1b[38;5;1;48;5;101m│\x1b[0m              \x1b[38;5;1;48;5;101m│\x1b[0m\n\x1b[38;5;1;48;5;101m└──────────────│\x1b[0m                  \x1b[38;5;1;48;5;101m│──────────────┘\x1b[0m\n\x1b[38;5;1;48;5;101m┌──────────────│\x1b[0m         \x1b[38;2;120;0;120;48;2;0;170;170m4\x1b[0m        \x1b[38;5;1;48;5;101m│──────────────┐\x1b[0m\n\x1b[38;5;1;48;5;101m│\x1b[0m              \x1b[38;5;1;48;5;101m│\x1b[0m                  \x1b[38;5;1;48;5;101m│\x1b[0m              \x1b[38;5;1;48;5;101m│\x1b[0m\n\x1b[38;5;1;48;5;101m│\x1b[0m              \x1b[38;5;1;48;5;101m│\x1b[0m                  \x1b[38;5;1;48;5;101m│\x1b[0m              \x1b[38;5;1;48;5;101m│\x1b[0m\n\x1b[38;5;1;48;5;101m│\x1b[0m              \x1b[38;5;1;48;5;101m│\x1b[0m                  \x1b[38;5;1;48;5;101m│\x1b[0m              \x1b[38;5;1;48;5;101m│\x1b[0m\n\x1b[38;5;1;48;5;101m│\x1b[0m              \x1b[38;5;1;48;5;101m└──────────────────┘\x1b[0m              \x1b[38;5;1;48;5;101m│\x1b[0m\n\x1b[38;5;1;48;5;101m│\x1b[0m                       \x1b[38;5;1;48;5;101m││\x1b[0m                       \x1b[38;5;1;48;5;101m│\x1b[0m\n\x1b[38;5;1;48;5;101m│\x1b[0m           \x1b[38;2;120;0;120;48;2;0;170;170m2\x1b[0m           \x1b[38;5;1;48;5;101m││\x1b[0m           \x1b[38;2;120;0;120;48;2;0;170;170m3\x1b[0m           \x1b[38;5;1;48;5;101m│\x1b[0m\n\x1b[38;5;1;48;5;101m│\x1b[0m                       \x1b[38;5;1;48;5;101m││\x1b[0m                       \x1b[38;5;1;48;5;101m│\x1b[0m\n\x1b[38;5;1;48;5;101m│\x1b[0m                       \x1b[38;5;1;48;5;101m││\x1b[0m                       \x1b[38;5;1;48;5;101m│\x1b[0m\n\x1b[38;5;1;48;5;101m│\x1b[0m                       \x1b[38;5;1;48;5;101m││\x1b[0m                       \x1b[38;5;1;48;5;101m│\x1b[0m\n\x1b[38;5;1;48;5;101m│\x1b[0m                       \x1b[38;5;1;48;5;101m││\x1b[0m                       \x1b[38;5;1;48;5;101m│\x1b[0m\n\x1b[38;5;1;48;5;101m│\x1b[0m                       \x1b[38;5;1;48;5;101m││\x1b[0m                       \x1b[38;5;1;48;5;101m│\x1b[0m\n\x1b[38;5;1;48;5;101m└───────────────────────┘└───────────────────────┘\x1b[0m\n                                                  ",
		frame,
	)
}

func TestBgColor(t *testing.T) {
	fb := NewFrameBuffer(20, 6)

	s := l.Vercital().
		Constrains(
			layout.NewConstrain(layout.Length, 3),
			layout.NewConstrain(layout.Length, 3),
		).
		Split(fb.Size())

	fb.RenderString(
		lb.Border(lb.NormalBorder(lb.WithBgColorRGB(70, 20, 33)),
			lb.SetColor(
				lb.Color(1),
				lb.SetColor(
					lb.ColorBg(52),
					lb.ExpandHorizontal(20-2, lb.Right, "first, 1"),
				),
			),
		),
		s[0],
	)

	fb.RenderString(
		lb.Border(lb.NormalBorder(lb.WithBgColorRGB(20, 0, 78)),
			lb.SetColor(
				lb.ColorBg(74),
				lb.ExpandHorizontal(20-2, lb.Center, "second, 1"),
			),
		),
		s[1],
	)

	frame := fb.View()

	// fmt.Println(lb.BorderN(frame))
	// fmt.Printf("%q\n", lb.BorderN(frame))

	assert.Equal(t,
		"┌────────────────────┐\n│\x1b[48;2;70;20;33m┌──────────────────┐\x1b[0m│\n│\x1b[48;2;70;20;33m│\x1b[0m\x1b[38;5;1;48;5;52m          first, 1\x1b[0m\x1b[48;2;70;20;33m│\x1b[0m│\n│\x1b[48;2;70;20;33m└──────────────────┘\x1b[0m│\n│\x1b[48;2;20;0;78m┌──────────────────┐\x1b[0m│\n│\x1b[48;2;20;0;78m│\x1b[0m\x1b[48;5;74m     second, 1    \x1b[0m\x1b[48;2;20;0;78m│\x1b[0m│\n│\x1b[48;2;20;0;78m└──────────────────┘\x1b[0m│\n└────────────────────┘",
		lb.BorderN(frame))
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
			lb.Border(lb.NormalBorder(lb.WithFgColor(uint8(i+100))),
				lb.ExpandVertical(int(s.Height)-2, lb.Center,
					lb.ExpandHorizontal(int(s.Width)-2, lb.Center,
						lb.SetColor(lb.Color(9),
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
			lb.Border(lb.NormalBorder(),
				lb.ExpandVertical(int(s.Height)-2, lb.Bottom,
					lb.ExpandHorizontal(int(s.Width)-2, lb.Bottom, ""),
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

	border := lb.NormalBorder(lb.WithFgColor(70))

	fb.RenderString(
		lb.Border(border,
			lb.ExpandVertical(int(splits[0].Height)-2, lb.Bottom,
				lb.ExpandHorizontal(int(splits[0].Width)-2, lb.Bottom, ""),
			)),
		splits[0],
	)

	fb.RenderString(
		lb.Border(border,
			lb.ExpandVertical(int(splits[1].Height)-2, lb.Bottom,
				lb.ExpandHorizontal(int(splits[1].Width)-2, lb.Bottom, ""),
			)),
		splits[1],
	)

	frame := fb.View()

	// fmt.Println(frame)

	assert.Equal(t,
		"┌──────────────────────────────────────────────────┐\n│\x1b[38;5;70m┌───────────────────────┐┌───────────────────────┐\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m│\x1b[0m                       \x1b[38;5;70m││\x1b[0m                       \x1b[38;5;70m│\x1b[0m│\n│\x1b[38;5;70m└───────────────────────┘└───────────────────────┘\x1b[0m│\n└──────────────────────────────────────────────────┘",
		lb.Border(lb.NormalBorder(), frame),
	)
}

func getDump(name string) string {
	data, err := os.ReadFile(fmt.Sprintf("%s.dump", name))
	if err != nil {
		panic(err)
	}
	return string(data)
}
