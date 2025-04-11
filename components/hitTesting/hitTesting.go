package hitTesting

// minimal hit-testing package

import lbl "github.com/crolbar/lipbalm/layout"

type HitTesting struct {
	HitTriggers []func(any)

	// arg passed to the func on hit
	Argument any
}

// offset the rect from the side to make the hit area smaller
// useful when you don't want click on a border
var (
	TopOff    = 0
	RightOff  = 0
	BottomOff = 0
	LeftOff   = 0
)

// size = number of triggers = number of rectangles
func InitHT(size int) HitTesting {
	return HitTesting{
		HitTriggers: make([]func(any), size),
	}
}
func InitHTA(size int, arg any) HitTesting {
	return HitTesting{
		HitTriggers: make([]func(any), size),
		Argument:    arg,
	}
}

// assumes i is < than the size given in the init
//
// make sure i is equal to the idx of rectangle (in the rects
// list you give in CheckHit) you want to map this trigger to.
func (ht *HitTesting) SetTrigger(i int, c func(any)) {
	ht.HitTriggers[i] = c
}

// appends a trigger to the end of hit triggers
// THIS WILL GROW THE SIZE! USE SetTrigger IF YOU DON'T WANT THAT
func (ht *HitTesting) AppendRect(c func(any)) {
	ht.HitTriggers = append(ht.HitTriggers, c)
}

// check up until it find a hit
func (ht HitTesting) CheckHit(x, y int, rects []lbl.Rect) {
	for i := 0; i < len(ht.HitTriggers); i++ {
		if HitTest(x, y, rects[i]) {
			ht.HitTriggers[i](ht.Argument)
			break
		}
	}
}

func HitTest(x, y int, box lbl.Rect) bool {
	// out to the left
	if x < int(box.X)+LeftOff {
		return false
	}
	// out to the right
	if x > int(box.X+box.Width-1)-RightOff {
		return false
	}
	// out to the top
	if y < int(box.Y)+TopOff {
		return false
	}
	// out to the bottom
	if y > int(box.Y+box.Height-1)-BottomOff {
		return false
	}

	return true
}
