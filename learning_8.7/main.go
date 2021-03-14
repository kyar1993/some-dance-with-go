/*
 * Модуль 8.7
 */
package main

import (
	"fmt"
	"learning_8.7/electronic"
)

func main() {
	androidPhone := electronic.AndroidPhoneConstruct("Xiaomi", "Note 10")
	applePhone := electronic.ApplePhoneConstruct("11 Pro Max")
	radioPhone := electronic.RadioPhoneConstruct("Panasonic", "KX-TG1612", 20)

	printCharacteristics(androidPhone)
	printCharacteristics(applePhone)
	printCharacteristics(radioPhone)
}

// Вывод бренд, модель, тип телефона
func printCharacteristics(p electronic.Phone) {
	fmt.Printf("Бренд: %s\n", p.Brand())
	fmt.Printf("Модель: %s\n", p.Model())
	fmt.Printf("Тип: %s\n", p.Type())

	switch p := p.(type) {
	case electronic.StationPhone:
		fmt.Printf("Количество кнопок: %d\n", p.ButtonsCount())

		break
	case electronic.Smartphone:
		fmt.Printf("Операционная система: %s\n", p.OS())
		break
	}

	fmt.Print("\n")
}
