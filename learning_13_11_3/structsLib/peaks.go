package structsLib

import (
	"math"
)

// Вершина
type Peak struct {
	Id     int   // id вершины
	Weight int   // уточнённый вес вершины
	Prev   *Peak // для обратного пути
}

// вершины
// вес первой (начальной) = 0
var FirstPeak = &Peak{Id: 1, Weight: 0, Prev: nil}
var SecondPeak = &Peak{Id: 2, Weight: math.MaxInt32, Prev: nil}
var ThirdPeak = &Peak{Id: 3, Weight: math.MaxInt32, Prev: nil}
var FourthPeak = &Peak{Id: 4, Weight: math.MaxInt32, Prev: nil}

func GetPeakById(id int) *Peak {
	switch id {
	case 1:
		return FirstPeak
	case 2:
		return SecondPeak
	case 3:
		return ThirdPeak
	// 4
	default:
		return FourthPeak
	}
}
