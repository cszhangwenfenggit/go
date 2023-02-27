package extend

type Kilo float64
type Pound float64

func KToP(k Kilo) Pound {
	return Pound(2.20462 * k)
}

func PTok(p Pound) Kilo {
	return Kilo(0.453592 * p)
}
