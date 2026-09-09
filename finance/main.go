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
	var expCreditVTB expenses = expenses{165.43, 4961, 2486.12, 190.98, 484.67, 150, 60, 1298.27, 159.99, 2517, 79.90, 506.45, 343.29, 485.95, 640} //это список моих расходов по кредитке
	var expInstallmentOzon expenses = expenses{10000, 469, 2591, 131, 2576, 929, 1294}
	var ownFounds float64 = 5047.37 //
	DescriptionAll(expCreditVTB, expInstallmentOzon, ownFounds)
	var yn byte
	fmt.Println("Добавить расходы?y/n")
	fmt.Scanln(&yn)
	if yn == 'y' {
		fmt.Println("Выберите счёт:")
		fmt.Println(" 1 - Кредитная карта ВТБ")
		fmt.Println(" 2 - перевод со счёта Ozon рассрочки")
		fmt.Scanln(&yn)
		if yn == 1 {
			AddExpenses(&expCreditVTB)
			DescriptionAll(expCreditVTB, expInstallmentOzon, ownFounds)
		} else if yn == 2 {
			AddExpenses(&expInstallmentOzon)
			DescriptionAll(expCreditVTB, expInstallmentOzon, ownFounds)
		} else {
			fmt.Println(errors.New("Неправильный ввод"))
		}
	} else if yn != 'n' {
		fmt.Println(errors.New("Неправильный ввод (y/n)"))
	}
}
