// Структура androidPhone - Телефон Apple
package electronic

// private struct androidPhone
type androidPhone struct {
	brand     string
	model     string
	phoneType string
	os        string
}

// public Construct for androidPhone structAndroidPhoneConstruct
func AndroidPhoneConstruct(brand, model string) androidPhone {
	return androidPhone{brand: brand, model: model, phoneType: "smartphone", os: "Android"}
}

func (p androidPhone) Brand() string { return p.brand }

func (p androidPhone) Model() string { return p.model }

func (p androidPhone) Type() string { return p.phoneType }

func (p androidPhone) OS() string { return p.os }
