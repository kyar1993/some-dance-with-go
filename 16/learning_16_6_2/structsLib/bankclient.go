// банковский клиент
package structsLib

import (
	"errors"
	"fmt"
)

var balanceError = errors.New("На балансе клиента недостаточно средств!!!")

//// это трюк для проверки типа: до тех пор пока InMemoryCache не будет
//// реализовывать интерфейс Cache, программа не запустится
var _ BankClient = &BankClientItem{}

type BankClient interface {
	// Deposit deposits given amount to clients account
	Deposit(amount int)

	// Withdrawal withdraws given amount from clients account.
	// return error if clients balance less the withdrawal amount
	Withdrawal(amount int) error

	// Balance returns clients balance
	Balance() int
}

// клиент банка
type BankClientItem struct {
	clientId int
	amount   int
}

// конструктор для клиента банка
func NewBankClientItem(clientId int, amount int) *BankClientItem {
	return &BankClientItem{clientId, amount}
}

// внесение на счёт клиент
func (c *BankClientItem) Deposit(amount int) {
	c.amount += amount
}

// списание со счёта клиента
func (c *BankClientItem) Withdrawal(amount int) error {
	// проверить, что сумма на счету достаточная, если нет - кидаем ошибку
	balance := c.Balance()

	if balance < amount {
		return balanceError
	}

	c.amount -= amount

	return nil
}

// текущий баланс клиента
func (c *BankClientItem) Balance() int {
	return c.amount
}

// обёртка над списанием, с выводом ошибок
func (c *BankClientItem) WithdrawalWrapper(amount int) {
	err := c.Withdrawal(amount)

	if err != nil {
		c.DisplayError(err)
	}
}

// отображение баланса
func (c *BankClientItem) DisplayClientBalance() {
	fmt.Printf("Баланс вашего счёта составляет: %d coin \n", c.Balance())
}

// вывод ошибок
func (c *BankClientItem) DisplayError(err error) {
	//fmt.Println(err)
	fmt.Printf("Ошибка!!! %s\n", err)
}
