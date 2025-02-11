package lipbalm

import (
	"fmt"
	"os"
	"testing"
	"time"

	// "github.com/charmbracelet/lipgloss"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	str1 := "hello\nworld\nline\nthhreetoetuhoe\nyes"
	str2 := "bye\nworld\nwidth\nsucks"

	// fmt.Println("str1:\n", str1)
	// fmt.Println()
	// fmt.Println("str2:\n", str2)
	// fmt.Println()

	start := time.Now()
	fmt.Printf("%q\n", JoinVertical(Center,
		str1, str2,
	))
	fmt.Println("time for main", time.Since(start))

	// fmt.Println()

	// fmt.Println(JoinHorizontal(Top,
	// 	str1, str2,
	// ))
}

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
