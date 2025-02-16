package layout

type Direction int

const (
	Vertical Direction = iota
	Horizontal
)

type Layout struct {
	direction  Direction
	constrains []Constrain
}

func DefaultLayout() Layout {
	return Layout{}
}

func (l Layout) Vercital() Layout {
	l.direction = Vertical
	return l
}

func (l Layout) Horizontal() Layout {
	l.direction = Horizontal
	return l
}

func (l Layout) Constrains(constrains []Constrain) Layout {
	l.constrains = constrains
	return l
}
