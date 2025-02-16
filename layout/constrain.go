package layout

type ConstrainType int

const (
	Min ConstrainType = iota
	Max
	Lenght
	Percent
)

type Constrain struct {
	t ConstrainType
	v uint16
}

func NewConstrain(t ConstrainType, v uint16) Constrain {
	return Constrain{
		t: t,
		v: v,
	}
}
