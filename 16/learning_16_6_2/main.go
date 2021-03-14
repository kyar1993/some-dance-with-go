// 16.6.2 Напишите структуру, которая будет реализовывать клиент для клиента банковского приложения. Этот клиент должен реализовывать следующий интерфейс:Используя этот клиент, создайте консольное приложение, которое:
//
// В момент старта запускает 10 горутин, каждая из которых с промежутком от
// 0.5 секунд до 1 секунды зачисляет на счёт клиента случайную сумму от 1 до 10.
// Так же запускается 5 горутин, которые с промежутком 0.5 секунд до 1 секунды
// снимают с клиента случайную сумму от 1 до 5. Если снятие невозможно, в консоль
// выводится сообщение об ошибке, и приложение продолжает работу.
// Если пользователь введет в консоль слово balance — приложение выведет в консоль текущий баланс клиента.
// deposit — запрашивается сумма (целое число) — которая добавляется на счёт пользователя
// withdrawal — запрашивается сумма (целое число) — которая выводится со счёта пользователя.
// Если пользователь введет слова exit — приложение завершит работу.
// При вводе любой другой команды приложение выведет сообщение
// “Unsupported command. You can use commands: balance, deposit, withdrawal, quit.”
package main

import (
	"errors"
	"fmt"
	"learning_16_6_2/structsLib"
	"math/rand"
	"strings"
	"sync"
	"time"
)

// команда из списка balance, deposit, withdrawal, quit
var command string

// сумма списания / начисления, введённая клиентом
var amount int

// ошибка
var err error

var unsupportedCommandError = errors.New("Unsupported command. You can use commands: balance, deposit, withdrawal, quit")

// создаём клиента
var client = structsLib.NewBankClientItem(1, 0)

var m sync.RWMutex
var rs, ws int

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	fmt.Println("Добро пожаловать в приложение нашего банка, Уважаемый Клиент!!!")

	// В момент старта запускает 10 горутин, каждая из которых с промежутком
	// от 0.5 секунд до 1 секунды зачисляет на счёт клиента случайную сумму от 1 до 10.
	// Так же запускается 5 горутин, которые с промежутком 0.5 секунд до 1
	// секунды снимают с клиента случайную сумму от 1 до 5. Если снятие
	// невозможно, в консоль выводится сообщение об ошибке, и приложение продолжает работу.
	rsCh := make(chan int)
	wsCh := make(chan int)

	go func() {
		for {
			select {
			case n := <-rsCh:
				rs += n
			case n := <-wsCh:
				ws += n
			}
			fmt.Printf("%s%s\n", strings.Repeat("R", rs),
				strings.Repeat("W", ws))
		}
	}()

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go withdrawalMutexWrapper(rsCh, &m, &wg, client)
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go depositMutexWrapper(wsCh, &m, &wg, client)
	}

	wg.Wait()

	// получение команд
	dataReading()
}

// обёртка для списания с баланса
func withdrawalMutexWrapper(c chan int, m *sync.RWMutex, wg *sync.WaitGroup, client *structsLib.BankClientItem) {
	sleep()
	m.RLock()

	client.WithdrawalWrapper(randInt(1, 5))

	c <- 1
	sleep()
	c <- -1
	m.RUnlock()
	wg.Done()
}

// обёртка для пополнения баланса
func depositMutexWrapper(c chan int, m *sync.RWMutex, wg *sync.WaitGroup, client *structsLib.BankClientItem) {
	sleep()
	m.Lock()

	client.Deposit(randInt(1, 10))

	c <- 1
	sleep()
	c <- -1
	m.Unlock()
	wg.Done()
}

// обработчик команд
func commandsHandler() {
	switch command {
	// баланс
	case "balance":
		client.DisplayClientBalance()
		break
		// пополнение
	case "deposit":
		// запрос ввода значения
		depositReading()

		// сохрание
		client.Deposit(amount)
		// списание
	case "withdrawal":
		// запрос ввода значения
		withdrawalReading()

		// сохрание
		client.WithdrawalWrapper(amount)
		// выход из программы
	case "quit":
		return
	// ошибка - неверная команда
	default:
		// ошибка неверная команда
		fmt.Printf("%s\n", unsupportedCommandError)
	}

	// ожидание новой команды
	dataReading()
}

// программа должна выполняться n секунд и после этого закрываться.
func dataReading() {
	fmt.Println("\nУкажите тип операции, который желаете выполнить. Доступны следующие варианты: balance, deposit, withdrawal, quit")
	_, err = fmt.Scanln(&command)

	if err != nil {
		fmt.Printf("%s\n", unsupportedCommandError)
		dataReading()
		return
	}

	commandsHandler()

	return
}

// считывание суммы
func amountReading() {
	_, err = fmt.Scanln(&amount)

	if err != nil {
		fmt.Println("Допустимые значения - целые числа")
		amountReading()
		return
	}

	return
}

// считывание суммы пополнения
func depositReading() {
	fmt.Println("Укажите сумму пополнения:")
	amountReading()
}

// считывание суммы списания
func withdrawalReading() {
	fmt.Println("Укажите сумму списания:")
	amountReading()
}

// ожидание от 0.5 секунд до 1 секунды
func sleep() {
	time.Sleep(time.Duration(randInt(500, 1000)) * time.Millisecond)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
