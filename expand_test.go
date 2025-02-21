package lipbalm

import (
	"github.com/crolbar/lipbalm/assert"
	"testing"
)

func TestExpand(t *testing.T) {
	assert.Equal(t,
		"          \n          \n          \n          \n          \n          \n          \n          \n          \n        hi",
		Expand(10, 10, "hi"),
	)
}

func TestExpand2(t *testing.T) {
	assert.Equal(t,
		"          \n          \n          \n          \n          \n   text   \n          \n          \n          \n          \n          ",
		Expand(11, 10, "text", Center, Center),
	)
}

func TestExpand3(t *testing.T) {
	assert.Equal(t,
		"      text\n          ",
		Expand(2, 10, "text", Left, Bottom),
	)
}

func TestExpandBoth(t *testing.T) {
	s := "hi\noaeshu\naoeu"
	assert.Equal(t,
		"          \n          \n          \n          \n    hi    \n  oaeshu  \n   aoeu   \n          \n          \n          \n          ",
		ExpandHorizontal(10, Center,
			ExpandVertical(11, Center, s)),
	)
}

func TestExpandV(t *testing.T) {
	s := getDump("bar2")
	assert.Equal(t,
		"               \n               \n               \n               \n               \n               \n               \n\x1b[97m├─────────────┤\x1b[0m\n\x1b[97m│\x1b[0m\x1b[38;5;57m             \x1b[0m\x1b[97m│\x1b[0m\n\x1b[97m│\x1b[0m\x1b[38;5;57m             \x1b[0m\x1b[97m│\x1b[0m\n\x1b[97m│\x1b[0m\x1b[38;5;57m             \x1b[0m\x1b[97m│\x1b[0m\n\x1b[97m│\x1b[0m\x1b[38;5;57m             \x1b[0m\x1b[97m│\x1b[0m\n\x1b[97m│\x1b[0m\x1b[38;5;57m             \x1b[0m\x1b[97m│\x1b[0m\n\x1b[97m│\x1b[0m\x1b[38;5;57m             \x1b[0m\x1b[97m│\x1b[0m\n\x1b[97m│\x1b[0m\x1b[38;5;57m             \x1b[0m\x1b[97m│\x1b[0m\n\x1b[97m│\x1b[0m\x1b[38;5;57m             \x1b[0m\x1b[97m│\x1b[0m\n\x1b[97m│\x1b[0m\x1b[38;5;57m             \x1b[0m\x1b[97m│\x1b[0m\n\x1b[97m│\x1b[0m\x1b[38;5;57m             \x1b[0m\x1b[97m│\x1b[0m\n\x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m\n\x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m\n\x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m\n\x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m\n\x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m\n\x1b[97m└─────────────┘\x1b[0m\n               \n               \n               \n               \n               \n               \n               ",
		MakeSquare(Center, ExpandVertical(31, Center, s)),
	)

}

func TestExpandH(t *testing.T) {
	s := getDump("bar2")
	assert.Equal(t,
		"                                              \x1b[97m├─────────────┤\x1b[0m\n                                              \x1b[97m│\x1b[0m\x1b[38;5;57m             \x1b[0m\x1b[97m│\x1b[0m\n                                              \x1b[97m│\x1b[0m\x1b[38;5;57m             \x1b[0m\x1b[97m│\x1b[0m\n                                              \x1b[97m│\x1b[0m\x1b[38;5;57m             \x1b[0m\x1b[97m│\x1b[0m\n                                              \x1b[97m│\x1b[0m\x1b[38;5;57m             \x1b[0m\x1b[97m│\x1b[0m\n                                              \x1b[97m│\x1b[0m\x1b[38;5;57m             \x1b[0m\x1b[97m│\x1b[0m\n                                              \x1b[97m│\x1b[0m\x1b[38;5;57m             \x1b[0m\x1b[97m│\x1b[0m\n                                              \x1b[97m│\x1b[0m\x1b[38;5;57m             \x1b[0m\x1b[97m│\x1b[0m\n                                              \x1b[97m│\x1b[0m\x1b[38;5;57m             \x1b[0m\x1b[97m│\x1b[0m\n                                              \x1b[97m│\x1b[0m\x1b[38;5;57m             \x1b[0m\x1b[97m│\x1b[0m\n                                              \x1b[97m│\x1b[0m\x1b[38;5;57m             \x1b[0m\x1b[97m│\x1b[0m\n                                              \x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m\n                                              \x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m\n                                              \x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m\n                                              \x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m\n                                              \x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m\n                                              \x1b[97m└─────────────┘\x1b[0m",
		ExpandHorizontal(61, Right, s),
	)
}
