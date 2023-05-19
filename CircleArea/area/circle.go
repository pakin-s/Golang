package area

type CircleArea struct {
	Ra float64
}

func (c CircleArea) AreaCal() float64 {
	const PI float64 = 3.14
	return PI * c.Ra * c.Ra
}