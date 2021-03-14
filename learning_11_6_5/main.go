// Эмуляция маршрута метро оранжевой ветки Спб.
//
// P.s Реализация Стека на слайсах
package main

import (
	"fmt"
	entities "learning_11_6_5/linesEntities"
	lib "learning_11_6_5/structsLib"
	"time"
)

// ErrEmptyStack Ошибка при извлечении данных из пустого стека
var ErrEmptyStack error = fmt.Errorf("%s", "Стек пуст")

// смежные станции (переход на другие станции)
func adjacentList(station *lib.Station) string {
	var res string

	list := entities.AdjacentStationsList[station.Id]

	// нет соседних станций
	if len(list) == 0 {
		return ""
	}

	for _, adjacentStation := range list {
		data := *adjacentStation
		line := *data.MetroLine
		res += fmt.Sprintf("Переход на станцию %s, к поездам %s.\n", data.Name, line.Name)
	}

	return res
}

// строим маршрут
func buildARoute(stack *[]*lib.Station) {
	//push(stack, entities.Depo2222)
	push(stack, entities.Spasskaya)
	push(stack, entities.Dostoyevskaya)
	push(stack, entities.LigovskiyProspect)
	push(stack, entities.PloschadAlexandraNevskogo2)
	push(stack, entities.Novocherkasskaya)
	push(stack, entities.Ladozhskaya)
	push(stack, entities.ProspectBolshevikov)
	push(stack, entities.UlitsaDybenko)
	//push(stack, entities.Depo1111)
}

func bagsAttention() {
	fmt.Printf("Будьте внимательны и не забывайте Ваши вещи в вагонах метро!\n\n")
}

// объявить название станции
func announceStationName(station *lib.Station) {
	// проверить смежные станции
	fmt.Printf("Уважаемые пассажиры, поезд прибыл на станцию %s\n", station.Name)
	bagsAttention()
}

// закрытие дверей
func doorClosingAnnounce() {
	fmt.Println("Осторожно, двери закрываются.")
}

// объявить название следующей станции
func nextStationAnnounce(station *lib.Station) {
	fmt.Printf("Следующая станция %s\n", station.Name)

	// список соседних станций
	adjacents := adjacentList(station)

	if len(adjacents) > 0 {
		fmt.Println(adjacents)
	} else {
		fmt.Println()
	}
}

// сообщение о посадке/высадке пассажиров
func boardingPassengersNotification() {
	fmt.Println("Производится посадка / высадка пассажиров...\n")
}

// объявить название следующей станции
func routeEndAnnounce() {
	fmt.Printf("Уважаемые пассажиры, поезд прибыл в конечный пункт назначения.\n")
	fmt.Printf("Просьба всех покинуть вагоны.\n")
}

func fetchNextStation(route *[]*lib.Station) (*lib.Station, error) {
	// получаем элемент с вершины стека
	station, err := pop(route)

	return station, err
}

func main() {
	// создаём стек
	route := &[]*lib.Station{}

	// строим маршрут движения поезда (добавляем в стек)
	buildARoute(route)

	var goToStation, nextStation *lib.Station
	var err error

	for {
		// поезд только начал движение  - получаем 1 станцию из стека
		if goToStation == nil {
			goToStation, err = fetchNextStation(route)
			continue
		}

		// иначе сообщаем о прибытии на станцию
		goToStationData := *goToStation

		if goToStationData.Mute == false {
			announceStationName(goToStation)
		}

		nextStation, err = fetchNextStation(route)

		// ошибка
		if err != nil {
			routeEndAnnounce()
			break
		}

		goToStation = nextStation

		// оповещаем о посадке
		boardingPassengersNotification()
		// ожидаем
		time.Sleep(3 * time.Second)

		// оповещаем о закрытии дверей
		doorClosingAnnounce()

		// оповещаем о следующей станции
		nextStationAnnounce(nextStation)

		// ожидаем
		time.Sleep(2 * time.Second)

		// поезд едет к следующей станции
		tripInProgress(5)
	}
}

// размер стека
func size(stack *[]*lib.Station) int {
	return len(*stack)
}

// добавление в стек
func push(stack *[]*lib.Station, station *lib.Station) {
	*stack = append(*stack, station)
}

// получение крайнего элемента (вершины стека)
func pop(stack *[]*lib.Station) (*lib.Station, error) {
	size := size(stack)

	// стек пуст
	if size == 0 {
		return &lib.Station{}, ErrEmptyStack
	}

	el := (*stack)[size-1]
	*stack = (*stack)[:size-1]

	return el, nil
}

// фон для поездки
func tripInProgress(repeatCount int) {
	for i := 1; i <= repeatCount; i++ {
		fmt.Println("Чу-чух Чу-чух\n")
		time.Sleep(1 * time.Second)
	}
}
