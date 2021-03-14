// Структура radioPhone - Стационарный Телефон
package electronic

// private struct radioPhone
type radioPhone struct {
	brand        string
	model        string
	phoneType    string
	buttonsCount int
}

// public Construct for radioPhone struct
func RadioPhoneConstruct(brand, model string, buttonsCount int) radioPhone {
	return radioPhone{brand: brand, model: model, phoneType: "station", buttonsCount: buttonsCount}
}

func (p radioPhone) Brand() string { return p.brand }

func (p radioPhone) Model() string { return p.model }

func (p radioPhone) Type() string { return p.phoneType }

func (p radioPhone) ButtonsCount() int { return p.buttonsCount }
