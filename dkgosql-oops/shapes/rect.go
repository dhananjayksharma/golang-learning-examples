package shapes

type Shapers interface { // abstract
	Area(length, width float64) float64
	GetValue() float64
}

func NewShapers() rectangle { // instance
	return rectangle{}
}

type rectangle struct { // encaps
	length, width float64
}

func (r *rectangle) Area(length, width float64) float64 {
	r.length = length
	r.width = width

	return r.length * r.width
}

func (r rectangle) GetValue() float64 {
	return r.length * r.width
}
