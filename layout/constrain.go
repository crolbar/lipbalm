package layout

type ConstrainType int

const (
	None ConstrainType = iota

	// space is taken in this order going top to bottom
	//
	// Length, Min
	// Percent
	// Max
	// Min

	// Takes exactly this amount of space
	// before everything else
	Length

	// Takes this exact amount of space before others
	// and takes the last remainder of space after others
	Min

	// Takes % of remainder of space after Length & Min
	Percent

	// Takes exact amount of remainder space after Length &
	// before Min takes its remainder
	Max
)

type Constrain struct {
	t ConstrainType
	v uint16
}

func NewConstrain(t ConstrainType, v uint16) Constrain {
	return Constrain{t: t, v: v}
}
