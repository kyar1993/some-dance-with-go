package main

import "fmt"

type CoffeeMachine interface {
	Type() string
	Brand() string
	Model() string
}
type baseCoffeeMachine struct {
	brand string
	model string
}

func (m *baseCoffeeMachine) Brand() string {
	return m.brand
}
func (m *baseCoffeeMachine) Model() string {
	return m.model
}

//капельная кофемашина
type DripCoffeeMachine struct {
	baseCoffeeMachine
}

func (m *DripCoffeeMachine) Type() string { return "drip" }

//капсульная кофемашина
type CapsuleCoffeeMachine struct {
	baseCoffeeMachine
}

func (m *CapsuleCoffeeMachine) Type() string { return "capsule" }

//рожковая кофемашина
type CarobCoffeeMachine struct {
	baseCoffeeMachine
}

func (m *CarobCoffeeMachine) Type() string { return "carob" }

func main() {
	var coffeeMachine CoffeeMachine = &CapsuleCoffeeMachine{}

	fmt.Println(coffeeMachine.Type())
}

//var (
//	veryExpensiveCarErr = errors.New("слишком дорогая машина")
//	carNotExistErr      = errors.New("машина не существует")
//)
//
//type car struct {
//	brand string
//	model string
//}
//
//func main() {
//	car, err := buyCar("BMW", "X6", 3000000)
//	switch err {
//	case veryExpensiveCarErr:
//		fmt.Println("Нужно поднакопить денег")
//	case carNotExistErr:
//		fmt.Println("Выберите другую машину")
//	case nil:
//		fmt.Printf("Поздравляем с покупкой %s %s", car.brand, car.model)
//	}
//}
//
//func buyCar(desiredBrand, desiredModel string, moneyAmount int) (car, error) {
//	carPrice, err := getCarPrice(desiredBrand, desiredModel)
//	if err != nil {
//		return car{}, err
//	}
//	if carPrice > moneyAmount {
//		return car{}, veryExpensiveCarErr
//	}
//	return car{brand: desiredBrand, model: desiredModel}, nil
//}
//
//func getCarPrice(brand, model string) (int, error) {
//	if brand == "BMW" && model == "X6" {
//		return 6000000, nil
//	}
//	if brand == "Audi" && model == "A7" {
//		return 4000000, nil
//	}
//	if brand == "Mazda" && model == "6" {
//		return 2000000, nil
//	}
//	return 0, carNotExistErr
//}

//func main() {
//	result, err := joinStringsWithSpace("John", "Smith")
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(result)
//
//	result, err = joinStringsWithSpace("", "apple")
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(result)
//}
//
//func joinStringsWithSpace(str1, str2 string) (string, error) {
//	if str1 == "" {
//		return "", errors.New("первая строка пустая")
//	}
//	if str2 == "" {
//		return "", errors.New("вторая строка пустая")
//	}
//
//	return str1 + " " + str2, nil
//}

//func main() {
//	fmt.Println(recursive(5))
//}
//
//func recursive(number int) int {
//	if number > 1 {
//		return number * recursive(number-1)
//	}
//	return 1
//}

//func main() {
//	//fmt.Println("Функция вернула: " + some_lib.Test())
//	//"fmt"
//	//"fake.com/test/some_lib"
//
//	j := 1
//
//	for i := 1; i < 100; i *= 2 {
//		i += 4
//
//		if i%2 == 0 {
//			println("чётное")
//			j++
//		}
//	}
//
//	println(j)
//}
