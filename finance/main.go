package main

import (
	"errors"
	"fmt"
)

const lim float64 = 50000

type expenses []float64

func (e expenses) Sum() float64 { // Считаем сумму
	var sum float64
	for i := range e {
		sum += e[i]
	}
	return sum
}

func BalanceLimit(s float64) float64 { // Считаем остаток лимита
	return lim - s
}

func AddExpenses(e *expenses) {
	var amount float64
	fmt.Println("Введите сумму:")
	fmt.Scanln(&amount)
	*e = append(*e, amount)
}

func DescriptionAll(eVTB, eOzon expenses, my float64) {

	fmt.Println()
	fmt.Println()
	fmt.Printf("Сумма расходов по кредитной карте ВТБ: %0.f рублей\n", eVTB.Sum())
	fmt.Printf("Остаток лимита по кредитной карте ВТБ: %0.f рублей\n", BalanceLimit(eVTB.Sum()))
	fmt.Println()
	fmt.Printf("Сумма моих расходов по рассрочке Ozon: %0.f рублей\n", eOzon.Sum())
	fmt.Printf("Остаток лимита Ozon рассрочке без вычета обязательного платежа: %0.f рублей\n", BalanceLimit(eVTB.Sum()))
	fmt.Println()
	fmt.Printf("Сумма всех задолженностей: %0.f\n", eVTB.Sum()+eOzon.Sum())
	fmt.Printf("Мои собственные денежные средства: %0.f\n", my)
	fmt.Println()
	fmt.Println()
}
func main() {
	var expCreditVTB expenses = expenses{50000} //это список моих расходов по кредитке
	var expInstallmentOzon expenses = expenses{50000}
	var ownFounds float64 = 147.37
	DescriptionAll(expCreditVTB, expInstallmentOzon, ownFounds)
	var yn string
	fmt.Println("1 - Добавить расходы?")
	fmt.Println("2 - Добавить планируемые поступления?")
	fmt.Scanln(&yn)
	switch yn {
	case "1":
		fmt.Println("Выберите счёт:")
		fmt.Println(" 1 - Кредитная карта ВТБ")
		fmt.Println(" 2 - перевод со счёта Ozon рассрочки")
		fmt.Scanln(&yn)
		switch yn {
		case "1":
			AddExpenses(&expCreditVTB)
			DescriptionAll(expCreditVTB, expInstallmentOzon, ownFounds)
		case "2":
			AddExpenses(&expInstallmentOzon)
			DescriptionAll(expCreditVTB, expInstallmentOzon, ownFounds)
		default:
			fmt.Println(errors.New("Неправильный ввод"))
		}
	case "2":
		fmt.Println()

	default:
		fmt.Println(errors.New("Неправильный ввод (y/n)"))
	}
}
