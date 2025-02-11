package lipbalm

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/crolbar/lipbalm/assert"
)

func TestVJoin(t *testing.T) {
	assert.Equal(t,
		getDump("vertJoinExpect"),
		JoinVertical(Left,
			getDump("bar1"),
			getDump("bar2"),
		),
	)
}

func TestHJoin(t *testing.T) {
	assert.Equal(t,
		getDump("horiJoinExpect"),
		JoinHorizontal(Left,
			getDump("bar1"),
			getDump("bar2"),
		),
	)
}

func TestBothJoin(t *testing.T) {
	start := time.Now()
	assert.Equal(t,
		getDump("go29_screen"),
		JoinHorizontal(Top,
			JoinVertical(Right,
				getDump("go29_wheel"),
				getDump("go29_bp"),
			),
			getDump("go29_slider"),
		),
	)
	fmt.Println("time for both", time.Since(start))
}

func getDump(name string) string {
	data, err := os.ReadFile(fmt.Sprintf("tests/%s.dump", name))
	if err != nil {
		panic(err)
	}
	return string(data)
}
