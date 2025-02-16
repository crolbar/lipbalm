package layout

import (
	"fmt"
	"testing"

	"github.com/crolbar/lipbalm/assert"
)

func TestLayout(t *testing.T) {
	frame := NewRect(0, 10, 100, 100)

	layout := DefaultLayout().
		Horizontal().
		Constrains([]Constrain{
			NewConstrain(Lenght, 2),
			NewConstrain(Lenght, 96),
			NewConstrain(Lenght, 2),
		})

	assert.Equal(t, int(layout.direction), int(Horizontal))

	splits := layout.split(frame)
	fmt.Println(splits)

	assert.Equal(t, splits[0].y, 10)
	assert.Equal(t, splits[0].height, 2)

	assert.Equal(t, splits[1].y, 12)
	assert.Equal(t, splits[1].height, 96)

	assert.Equal(t, splits[2].y, 108)
	assert.Equal(t, splits[2].height, 2)
}

func TestRect(t *testing.T) {
	frame := NewRect(0, 0, 100, 100)

	assert.Equal(t, frame.x, 0)
	assert.Equal(t, frame.y, 0)
	assert.Equal(t, frame.width, 100)
	assert.Equal(t, frame.height, 100)
}

func TestConstrain(t *testing.T) {
	constrain := NewConstrain(Min, 5)

	assert.Equal(t, int(constrain.t), int(Min))
	assert.Equal(t, int(constrain.v), int(5))
}
