package utils

import (
	"math"
)


func Round(val float64,precision int) float64 {
	if precision==0{
		return math.Round(val)
	}
	p:=math.Pow10(precision)
	if precision<0 {
		return math.Floor(val*p+0.5)*math.Pow10(-precision)
	}
	return math.Floor(val*p+0.5)/p
}

