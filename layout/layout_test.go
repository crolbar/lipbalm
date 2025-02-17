package layout

import (
	"fmt"
	"testing"

	"github.com/crolbar/lipbalm/assert"
)

func TestLayout(t *testing.T) {
	frame := NewRect(0, 10, 100, 100)

	layout := DefaultLayout().
		Vercital().
		Constrains([]Constrain{
			NewConstrain(Length, 2),
			NewConstrain(Length, 96),
			NewConstrain(Length, 2),
		})

	assert.Equal(t, int(layout.direction), int(Vertical))

	splits := layout.Split(frame)
	fmt.Println(splits)

	assert.Equal(t, splits[0].Y, 10)
	assert.Equal(t, splits[0].Height, 2)

	assert.Equal(t, splits[1].Y, 12)
	assert.Equal(t, splits[1].Height, 96)

	assert.Equal(t, splits[2].Y, 108)
	assert.Equal(t, splits[2].Height, 2)
}

func TestRect(t *testing.T) {
	frame := NewRect(0, 0, 100, 100)

	assert.Equal(t, frame.X, 0)
	assert.Equal(t, frame.Y, 0)
	assert.Equal(t, frame.Width, 100)
	assert.Equal(t, frame.Height, 100)
}

func TestConstrain(t *testing.T) {
	constrain := NewConstrain(Min, 5)

	assert.Equal(t, int(constrain.t), int(Min))
	assert.Equal(t, int(constrain.v), int(5))
}

func TestSplit1(t *testing.T) {
	rect := NewRect(50, 50, 100, 100)

	layout := DefaultLayout().
		Horizontal().
		Constrains([]Constrain{
			NewConstrain(Min, 2),
			NewConstrain(Max, 5),
			NewConstrain(Length, 90),
		})

	splits := layout.Split(rect)
	fmt.Println(splits)

	assert.Equal(t, 50, splits[0].X)
	assert.Equal(t, 2+3, splits[0].Width) // 3 is remainder of 100 - 2 + 5 + 90

	assert.Equal(t, 55, splits[1].X)
	assert.Equal(t, 5, splits[1].Width)

	assert.Equal(t, 60, splits[2].X)
	assert.Equal(t, 90, splits[2].Width)
}

func TestSplit2(t *testing.T) {
	rect := NewRect(0, 0, 100, 100)

	layout := DefaultLayout().
		Vercital().
		Constrains([]Constrain{
			NewConstrain(Percent, 50),
			NewConstrain(Percent, 50),
		})

	splits := layout.Split(rect)
	fmt.Println(splits)

	assert.Equal(t, 0, splits[0].Y)
	assert.Equal(t, 50, splits[0].Height)

	assert.Equal(t, 50, splits[1].Y)
	assert.Equal(t, 50, splits[1].Height)
}

func TestSplit3(t *testing.T) {
	rect := NewRect(100, 0, 500, 0)

	layout := DefaultLayout().
		Horizontal().
		Constrains([]Constrain{
			NewConstrain(Min, 5),
			NewConstrain(Percent, 45),
			NewConstrain(Length, 10),
			NewConstrain(Percent, 50),
			NewConstrain(Max, 5),
		})

	splits := layout.Split(rect)
	fmt.Println(splits)

	assert.Equal(t, 100, splits[0].X)
	assert.Equal(t, 10, splits[0].Width)

	assert.Equal(t, 110, splits[1].X)
	assert.Equal(t, 0.45*500, splits[1].Width)

	assert.Equal(t, 110+0.45*500, splits[2].X)
	assert.Equal(t, 10, splits[2].Width)

	assert.Equal(t, 345, splits[3].X)
	assert.Equal(t, 0.5 * 500, splits[3].Width)

	assert.Equal(t, 100 + 500 - 5, splits[4].X)
	assert.Equal(t, 5, splits[4].Width)
}

func TestSplit4(t *testing.T) {
	rect := NewRect(0, 0, 100, 100)

	layout := DefaultLayout().
		Vercital().
		Constrains([]Constrain{
			NewConstrain(Min, 50),
			NewConstrain(Percent, 50),
			NewConstrain(Max, 5),
		})

	vert := layout.Split(rect)
	fmt.Println(vert)

	assert.Equal(t, 0, vert[0].Y)
	assert.Equal(t, 50, vert[0].Height)

	assert.Equal(t, 50, vert[1].Y)
	assert.Equal(t, 50, vert[1].Height)

	assert.Equal(t, 100, vert[2].Y)
	assert.Equal(t, 0, vert[2].Height)

	hori := DefaultLayout().
		Horizontal().
		Constrains([]Constrain{
			NewConstrain(Min, 5),
			NewConstrain(Length, 50),
			NewConstrain(Max, 5),
		}).Split(vert[1])

	assert.Equal(t, 0, hori[0].X)
	assert.Equal(t, 50, hori[0].Y)
	assert.Equal(t, 45, hori[0].Width)

	assert.Equal(t, 45, hori[1].X)
	assert.Equal(t, 50, hori[1].Width)

	assert.Equal(t, 95, hori[2].X)
	assert.Equal(t, 5, hori[2].Width)

}
