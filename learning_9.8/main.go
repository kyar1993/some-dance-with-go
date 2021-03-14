/*
 * Модуль 9.8
 */
package main

import "fmt"

const manGender = "Male"
const womanGender = "Female"

var mostDangerousMan Man

type Man struct {
	Name     string
	LastName string
	Age      int
	Gender   string
	Crimes   int
}

func main() {
	// сезон охоты на паники - открыт!
	defer errorsCatch()

	// создаём список подозреваемых (словарь)
	// (ключ - имя человека, значением - экземпляр структуры Man)
	people := fetchSuspectsList()

	//	проверяем наличие подозреваемых в "БД"
	//	нет - выдаём ошибку
	if len(people) == 0 {
		panic("В базе данных нет информации по запрошенным подозреваемым")
	}

	// создаём слайс — suspects, в котором перечислены имена подозреваемых
	// длина (length) = 0
	// вместимость (capacity) = количеству записей
	// P.s не нужно тратить ресурсы на расширение
	suspects := make([]string, 0, len(people))

	// обрабатываем список
	for _, man := range people {
		// добавляем имена в слайс
		suspects = append(suspects, man.Name)

		// вычисляем подозреваемого с наибольшим количество совершённых преступлений
		if mostDangerousMan.Crimes < man.Crimes {
			mostDangerousMan = man
		}
	}

	// подозреваемые раньше не совершали преступлений - выдаём ошибку
	if mostDangerousMan.Crimes == 0 {
		panic("Подозреваемые не совершали преступлений!!! Надежда только на Вас!!!")
	}

	// выводим список имён подозреваемых
	displaySuspectsNamesList(suspects)

	// выводим данные особо опасного подозреваемого
	displayMostDangerousSuspect(&mostDangerousMan)
}

// обработка ошибок (паник)
func errorsCatch() {
	if r := recover(); r != nil {
		fmt.Println("Attention! Внимание! Achtung!")
		fmt.Printf("%s\n", r)
	}
}

// получаем список подозреваемых из "БД"
func fetchSuspectsList() map[string]Man {
	return map[string]Man{
		"Rick":   Man{Name: "Rick", LastName: "Sanchez", Age: 70, Gender: manGender, Crimes: 323},
		"Morty":  Man{Name: "Morty", LastName: "Smith", Age: 14, Gender: manGender, Crimes: 23},
		"Peter":  Man{Name: "Peter", LastName: "Griffin", Age: 45, Gender: manGender, Crimes: 54},
		"Brian":  Man{Name: "Brian", LastName: "Griffin", Age: 6, Gender: manGender, Crimes: 14},
		"Stewie": Man{Name: "Stewie", LastName: "Griffin", Age: 2, Gender: manGender, Crimes: 127},
		"Homer":  Man{Name: "Homer", LastName: "Simpson", Age: 38, Gender: manGender, Crimes: 100500},
		"Marge":  Man{Name: "Marge", LastName: "Simpson", Age: 34, Gender: womanGender, Crimes: 4},
		"Bart":   Man{Name: "Bart", LastName: "Simpson", Age: 10, Gender: manGender, Crimes: 219},
		"Lisa":   Man{Name: "Lisa", LastName: "Simpson", Age: 8, Gender: womanGender, Crimes: 9},
		"Maggie": Man{Name: "Maggie", LastName: "Simpson", Age: 1, Gender: womanGender, Crimes: 1},
	}
}

// выводим список имён всех подозреваемых
func displaySuspectsNamesList(list []string) {
	fmt.Println("Список имён всех подозреваемых:")
	fmt.Printf("%v\n", list)
}

// выводим данные самого опасного подозреваемого
func displayMostDangerousSuspect(mostDangerousMan *Man) {
	fmt.Println()
	fmt.Println("Обратите внимание!!! Особо опасен!!!")
	fmt.Println("Имя:", mostDangerousMan.Name)
	fmt.Println("Фамилия:", mostDangerousMan.LastName)
	fmt.Println("Возраст:", mostDangerousMan.Age)
	fmt.Println("Пол:", mostDangerousMan.Gender)
	fmt.Println("Число преступлений:", mostDangerousMan.Crimes)
}
