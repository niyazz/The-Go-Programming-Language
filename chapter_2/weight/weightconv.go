package weightconv

import "fmt"

type Kilo float64
type Pund float64

func (kg Kilo) String() string{
	return fmt.Sprintf("%g` Kg ", kg)
}
func (p Pund) String() string{
	return fmt.Sprintf("%g lbs ", p)
}

const PundMulp = 2.2

func KgToP(kg Kilo) Pund {
	return Pund(kg * PundMulp);
}

func PToKg(p Pund) Kilo{
	return Kilo(p / PundMulp)
}