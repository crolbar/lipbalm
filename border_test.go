package lipbalm

import (
	"github.com/crolbar/lipbalm/assert"
	"testing"
)

func TestBorderTop(t *testing.T) {
	s := "something\nto border\nup"
	assert.Equal(t,
		"─────────\nsomething\nto border\nup       ",
		Border(NormalBorder(), s, false, true, true, true),
	)
}

func TestBorderBottom(t *testing.T) {
	s := getDump("bar1_no_border")
	assert.Equal(t,
		"\x1b[38;5;57m             \x1b[0m\n\x1b[38;5;57m             \x1b[0m\n\x1b[38;5;57m             \x1b[0m\n\x1b[38;5;57m             \x1b[0m\n\x1b[38;5;57m             \x1b[0m\n\x1b[38;5;57m█████████████\x1b[0m\n\x1b[38;5;57m█████████████\x1b[0m\n\x1b[38;5;57m█████████████\x1b[0m\n\x1b[38;5;57m█████████████\x1b[0m\n\x1b[38;5;57m█████████████\x1b[0m\n\x1b[38;5;57m█████████████\x1b[0m\n\x1b[38;5;57m█████████████\x1b[0m\n\x1b[38;5;57m█████████████\x1b[0m\n\x1b[38;5;57m█████████████\x1b[0m\n\x1b[38;5;57m█████████████\x1b[0m\n─────────────",
		Border(NormalBorder(), s, true, true, false, true),
	)
}

func TestBorderHorizontal(t *testing.T) {
	s := getDump("bar1_no_border")
	assert.Equal(t,
		"│\x1b[38;5;57m             \x1b[0m│\n│\x1b[38;5;57m             \x1b[0m│\n│\x1b[38;5;57m             \x1b[0m│\n│\x1b[38;5;57m             \x1b[0m│\n│\x1b[38;5;57m             \x1b[0m│\n│\x1b[38;5;57m█████████████\x1b[0m│\n│\x1b[38;5;57m█████████████\x1b[0m│\n│\x1b[38;5;57m█████████████\x1b[0m│\n│\x1b[38;5;57m█████████████\x1b[0m│\n│\x1b[38;5;57m█████████████\x1b[0m│\n│\x1b[38;5;57m█████████████\x1b[0m│\n│\x1b[38;5;57m█████████████\x1b[0m│\n│\x1b[38;5;57m█████████████\x1b[0m│\n│\x1b[38;5;57m█████████████\x1b[0m│\n│\x1b[38;5;57m█████████████\x1b[0m│",
		Border(NormalBorder(), s, true, false, true, false),
	)
}

func TestBorderAll(t *testing.T) {
	s := getDump("go29_screen")
	assert.Equal(t,
		"┌──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┐\n│\x1b[97m│\x1b[0m              left(17439)                               right(0)                \x1b[97m│\x1b[0m          \x1b[38;5;57m│\x1b[0m               range(900)               \x1b[38;5;57m│\x1b[0m│\n│\x1b[97m├────────────────────────────────────────\x1b[0m\x1b[97m────────────────────────────────────────┤\x1b[0m          \x1b[38;5;57m├────────────────────────────────────────┤\x1b[0m│\n│\x1b[97m│\x1b[0m\x1b[38;5;57m                   █████████████████████\x1b[0m\x1b[38;5;57m                                        \x1b[0m\x1b[97m│\x1b[0m          \x1b[38;5;57m│\x1b[0m\x1b[38;5;57m████████████████████████████████████████\x1b[0m\x1b[38;5;57m│\x1b[0m│\n│\x1b[97m│\x1b[0m\x1b[38;5;57m                   █████████████████████\x1b[0m\x1b[38;5;57m                                        \x1b[0m\x1b[97m│\x1b[0m          \x1b[38;5;57m│\x1b[0m\x1b[38;5;57m████████████████████████████████████████\x1b[0m\x1b[38;5;57m│\x1b[0m│\n│\x1b[97m│\x1b[0m\x1b[38;5;57m                   █████████████████████\x1b[0m\x1b[38;5;57m                                        \x1b[0m\x1b[97m│\x1b[0m          \x1b[38;5;57m│\x1b[0m\x1b[38;5;57m████████████████████████████████████████\x1b[0m\x1b[38;5;57m│\x1b[0m│\n│\x1b[97m└────────────────────────────────────────\x1b[0m\x1b[97m────────────────────────────────────────┘\x1b[0m          \x1b[38;5;57m└────────────────────────────────────────┘\x1b[0m│\n│                                                                                            \x1b[97m│\x1b[0m             autocenter(0)              \x1b[97m│\x1b[0m│\n│                                                                                            \x1b[97m├────────────────────────────────────────┤\x1b[0m│\n│                                                                                            \x1b[97m│\x1b[0m\x1b[38;5;57m                                        \x1b[0m\x1b[97m│\x1b[0m│\n│                                                                                            \x1b[97m│\x1b[0m\x1b[38;5;57m                                        \x1b[0m\x1b[97m│\x1b[0m│\n│                                                                                            \x1b[97m│\x1b[0m\x1b[38;5;57m                                        \x1b[0m\x1b[97m│\x1b[0m│\n│                                                                                            \x1b[97m└────────────────────────────────────────┘\x1b[0m│\n│                                                                                                                                      │\n│                                                                                                                                      │\n│                                                                                                                                      │\n│                                                                                                                                      │\n│                                                                                                                                      │\n│                                                                                                                                      │\n│                                                                                                                                      │\n│                                                                                                                                      │\n│                                                                                                                                      │\n│                                                                                                                                      │\n│                                                                                                                                      │\n│                                                                                                                                      │\n│                                                                                                                                      │\n│                       ┌───┐┌───┐┌───┐┌───┐┌────┐┌────┐┌────┐┌────┐┌────┐┌────┐                                                       │\n│                       │   ││   ││   ││   ││    ││\x1b[48;5;57m    \x1b[0m││    ││    ││    ││    │                                                       │\n│                       │ X ││ S ││ C ││ T ││ RP ││\x1b[48;5;57m \x1b[0m\x1b[48;5;57mLP\x1b[0m\x1b[48;5;57m \x1b[0m││ R2 ││ L2 ││ SH ││ OP │                                                       │\n│                       │   ││   ││   ││   ││    ││\x1b[48;5;57m    \x1b[0m││    ││    ││    ││    │                                                       │\n│                       └───┘└───┘└───┘└───┘└────┘└────┘└────┘└────┘└────┘└────┘                                                       │\n│                       ┌────┐┌────┐┌────┐┌────┐┌────┐┌────┐┌────┐┌────┐┌───┐┌────┐                                                    │\n│                       │    ││    ││    ││    ││    ││\x1b[48;5;57m    \x1b[0m││    ││    ││   ││\x1b[48;5;57m    \x1b[0m│                                                    │\n│                       │ R3 ││ L3 ││ 1t ││ 2d ││ 3d ││\x1b[48;5;57m \x1b[0m\x1b[48;5;57m4h\x1b[0m\x1b[48;5;57m \x1b[0m││ 5h ││ 6h ││ R ││\x1b[48;5;57m \x1b[0m\x1b[48;5;57mPl\x1b[0m\x1b[48;5;57m \x1b[0m│                                                    │\n│                       │    ││    ││    ││    ││    ││\x1b[48;5;57m    \x1b[0m││    ││    ││   ││\x1b[48;5;57m    \x1b[0m│                                                    │\n│                       └────┘└────┘└────┘└────┘└────┘└────┘└────┘└────┘└───┘└────┘                                                    │\n│                       ┌────┐┌────┐┌────┐┌────┐┌────┐┌────┐┌────┐┌────┐┌────┐                                                         │\n│                       │    ││    ││    ││    ││    ││    ││    ││    ││    │                                                         │\n│                       │ Mi ││ RR ││ RL ││ RB ││ PS ││ DU ││ DL ││ DD ││ DR │                                                         │\n│                       │    ││    ││    ││    ││    ││    ││    ││    ││    │                                                         │\n│                       └────┘└────┘└────┘└────┘└────┘└────┘└────┘└────┘└────┘                                                         │\n│                       \x1b[97m│\x1b[0m clutch(138) \x1b[97m│\x1b[0m\x1b[97m│\x1b[0m break(242)  \x1b[97m│\x1b[0m\x1b[97m│\x1b[0mthrottle(157)\x1b[97m│\x1b[0m                                                                  │\n│                       \x1b[97m├─────────────┤\x1b[0m\x1b[97m├─────────────┤\x1b[0m\x1b[97m├─────────────┤\x1b[0m                                                                  │\n│                       \x1b[97m│\x1b[0m\x1b[38;5;57m             \x1b[0m\x1b[97m│\x1b[0m\x1b[97m│\x1b[0m\x1b[38;5;57m             \x1b[0m\x1b[97m│\x1b[0m\x1b[97m│\x1b[0m\x1b[38;5;57m             \x1b[0m\x1b[97m│\x1b[0m                                                                  │\n│                       \x1b[97m│\x1b[0m\x1b[38;5;57m             \x1b[0m\x1b[97m│\x1b[0m\x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m\x1b[97m│\x1b[0m\x1b[38;5;57m             \x1b[0m\x1b[97m│\x1b[0m                                                                  │\n│                       \x1b[97m│\x1b[0m\x1b[38;5;57m             \x1b[0m\x1b[97m│\x1b[0m\x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m\x1b[97m│\x1b[0m\x1b[38;5;57m             \x1b[0m\x1b[97m│\x1b[0m                                                                  │\n│                       \x1b[97m│\x1b[0m\x1b[38;5;57m             \x1b[0m\x1b[97m│\x1b[0m\x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m\x1b[97m│\x1b[0m\x1b[38;5;57m             \x1b[0m\x1b[97m│\x1b[0m                                                                  │\n│                       \x1b[97m│\x1b[0m\x1b[38;5;57m             \x1b[0m\x1b[97m│\x1b[0m\x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m\x1b[97m│\x1b[0m\x1b[38;5;57m             \x1b[0m\x1b[97m│\x1b[0m                                                                  │\n│                       \x1b[97m│\x1b[0m\x1b[38;5;57m             \x1b[0m\x1b[97m│\x1b[0m\x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m\x1b[97m│\x1b[0m\x1b[38;5;57m             \x1b[0m\x1b[97m│\x1b[0m                                                                  │\n│                       \x1b[97m│\x1b[0m\x1b[38;5;57m             \x1b[0m\x1b[97m│\x1b[0m\x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m\x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m                                                                  │\n│                       \x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m\x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m\x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m                                                                  │\n│                       \x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m\x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m\x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m                                                                  │\n│                       \x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m\x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m\x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m                                                                  │\n│                       \x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m\x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m\x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m                                                                  │\n│                       \x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m\x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m\x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m                                                                  │\n│                       \x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m\x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m\x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m                                                                  │\n│                       \x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m\x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m\x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m                                                                  │\n│                       \x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m\x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m\x1b[97m│\x1b[0m\x1b[38;5;57m█████████████\x1b[0m\x1b[97m│\x1b[0m                                                                  │\n│                       \x1b[97m└─────────────┘\x1b[0m\x1b[97m└─────────────┘\x1b[0m\x1b[97m└─────────────┘\x1b[0m                                                                  │\n└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘",
		Border(NormalBorder(), s),
	)
}
