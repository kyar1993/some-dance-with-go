// Структура applePhone - Телефон Apple
package electronic

// private struct applePhone
type applePhone struct {
	brand     string
	model     string
	phoneType string
	os        string
}

// public Construct for applePhone struct
func ApplePhoneConstruct(model string) applePhone {
	return applePhone{brand: "Apple", model: model, phoneType: "smartphone", os: "Ios"}
}

func (p applePhone) Brand() string { return p.brand }

func (p applePhone) Model() string { return p.model }

func (p applePhone) Type() string { return p.phoneType }

func (p applePhone) OS() string { return p.os }
