package exrepo

type volumer interface {
	Volume() float64
}

type cube struct {
	edge float64
} // edge x edge x edge

type triangularPrism struct {
	base     float64
	attitude float64
	height   float64
} // 0.5 x base x attitude x height

func VolumeOf(v volumer) float64 {
	return v.Volume()
}

func (c cube) Volume() float64 {
	return c.edge * c.edge * c.edge
}

func (t triangularPrism) Volume() float64 {
	return 0.5 * t.base * t.attitude * t.height
}
