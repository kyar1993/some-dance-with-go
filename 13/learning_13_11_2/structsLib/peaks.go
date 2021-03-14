package structsLib

// Вершина
type Peak struct {
	Id         int  // id вершины
	Proccessed bool // обработана ли вершина
}

// вершины
// вес первой (начальной) = 0
var FirstPeak = &Peak{Id: 1}
var SecondPeak = &Peak{Id: 2}
var ThirdPeak = &Peak{Id: 3}
var FourthPeak = &Peak{Id: 4}
var FifthPeak = &Peak{Id: 5}
var SixthPeak = &Peak{Id: 6}
var SeventhPeak = &Peak{Id: 7}
var EighthPeak = &Peak{Id: 8}
var NinthPeak = &Peak{Id: 9}
var TenthPeak = &Peak{Id: 10}

func GetPeakById(id int) *Peak {
	switch id {
	case 1:
		return FirstPeak
	case 2:
		return SecondPeak
	case 3:
		return ThirdPeak
	case 4:
		return FourthPeak
	case 5:
		return FifthPeak
	case 6:
		return SixthPeak
	case 7:
		return SeventhPeak
	case 8:
		return EighthPeak
	case 9:
		return NinthPeak
	// 10
	default:
		return TenthPeak
	}
}
