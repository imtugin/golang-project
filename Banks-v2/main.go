package main

import (
	"fmt"
)

var choice int
var d, asset Product
var amount float64
var err error

type Product interface {
	Income() float64
	IncomeMonth() float64
}

type param_error struct{}
type Deposit struct {
	Amount float64
	Term   int
	Annual float64
}
type SavingAccount struct {
	DailyBalance bool
	Amount       float64
	Annual       float64
	Term         int
}
type DigitalAssets struct {
	Amount float64
	Term   int
	Annual float64
	Tax    float64
}

func (object_error param_error) Error() string {
	return "invalid parameter"
}

func (s SavingAccount) Income() float64 {
	return s.Amount * ((s.Annual / 100) / 12)
}

func (d Deposit) Income() float64 {
	d.Amount *= (d.Annual / 100) / 12
	d.Amount *= float64(d.Term)
	return d.Amount
}

func (d DigitalAssets) Income() float64 {
	d.Amount *= (d.Annual / 100) / 12
	d.Amount *= float64(d.Term)
	d.Amount -= d.Amount * (d.Tax / 100)
	return d.Amount
}

func (s SavingAccount) IncomeMonth() float64 {
	return s.Amount * ((s.Annual / 100) / 12)
}

func (d Deposit) IncomeMonth() float64 {
	return d.Income() / float64(d.Term)
}

func (d DigitalAssets) IncomeMonth() float64 {
	return d.Income() / float64(d.Term)
}

func Description(a Product, err error) {
	if err != nil {
		fmt.Println(param_error{})
	} else {
		fmt.Println(a)
		fmt.Printf("Доход за весь период: %.f\n", a.Income())
		fmt.Printf("Доходность за месяц:%.f\n", a.IncomeMonth())
	}
}

// Нужно добавить чтобы при выводе цифрового актива мне предлагался вклад для сравнения на тот же срок
func ForComparsion(ac, rAn float64) {
	fmt.Println()
	fmt.Println("Вот для сравнения информация по вкладу на аналогичный срок:")
	// Для сравнения нужно в первую очередь достать отсюда процентную ставку
	value := OpenDeposit(ac).(Deposit)
	diff := rAn - value.Annual
	fmt.Printf("Процентная ставка: %.2f\n (разница - %.2f) Доход за весь период: %.f (%.f в месяц)\n", value.Annual, diff, value.Income(), value.IncomeMonth())
}

func OpenDeposit(a float64) Product { // Функция открывает депозит
	fmt.Println("На какой срок? ")
	fmt.Scanln(&choice)
	if choice < 3 {
		d = Deposit{
			Amount: a,
			Term:   choice,
			Annual: 11,
		}
	} else if 6 < choice && choice < 10 {
		d = Deposit{
			Amount: a,
			Term:   choice,
			Annual: 12.2,
		}
	} else if 9 < choice && choice < 12 {
		d = Deposit{
			Amount: a,
			Term:   choice,
			Annual: 12.1,
		}
	} else {
		switch choice {
		case 4:
			d = Deposit{
				Amount: a,
				Term:   choice,
				Annual: 13.6,
			}
		case 5:
			d = Deposit{
				Amount: a,
				Term:   choice,
				Annual: 12.3,
			}
		case 6:
			d = Deposit{
				Amount: a,
				Term:   choice,
				Annual: 13,
			}
		case 12:
			d = Deposit{
				Amount: a,
				Term:   choice,
				Annual: 13,
			}
		case 24:
			d = Deposit{
				Amount: a,
				Term:   choice,
				Annual: 11.8,
			}
		}

	}
	return d
}

func OpenDigitalAsset(amount float64) Product {
	fmt.Println("На какой срок? Доступно 1, 3, 6, 9 или 12 месяцев")
	fmt.Scanln(&choice)
	switch choice {

	case 1:
		d = DigitalAssets{
			Amount: amount,
			Term:   choice,
			Annual: 15.2,
			Tax:    13,
		}

	case 3:
		d = DigitalAssets{
			Amount: amount,
			Term:   choice,
			Annual: 15.7,
			Tax:    13,
		}

	case 6:
		d = DigitalAssets{
			Amount: amount,
			Term:   choice,
			Annual: 17,
			Tax:    13,
		}

	case 9:
		d = DigitalAssets{
			Amount: amount,
			Term:   choice,
			Annual: 17.5,
			Tax:    13,
		}

	case 12:
		d = DigitalAssets{
			Amount: amount,
			Term:   choice,
			Annual: 18,
			Tax:    13,
		}
	default:

		fmt.Println("ввод не верный")

	}
	return d
}

func OpenSavingAccount(a float64) Product {
	fmt.Println("Начисление на минимальную сумму - 1\n На ежедневный остаток - 2")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		d = SavingAccount{
			Amount:       a,
			DailyBalance: false,
			Annual:       12,
			Term:         1,
		}
	case 2:
		d = SavingAccount{
			Amount:       a,
			DailyBalance: true,
			Annual:       11.5,
			Term:         1,
		}
	}
	return d
}

// Функция выбора операции
// И здесь вывод ошибки ещё
// IncomeMonth чё то дурит теперь
// Добавил указатель. Мне нужны оригиналы а не копии, иначе в main будут работать функции с дефолтными значениями типа Product. Посмотрим как сработает
func Choice() (Product, error, func(Product, error)) {

	fmt.Println("Введите сумму")
	fmt.Scanln(&amount)
	fmt.Printf("Цифровой актив - 1\n Вклад - 2\n Накопительный счет - 3\n")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		asset = OpenDigitalAsset(amount)
		//		realAnnual := (asset.IncomeMonth() * 1200.0) / amount
		//		fmt.Printf("Реальная годовая процентная ставка после налогового вычета: %.2f\n", realAnnual)
		//		ForComparsion(amount, realAnnual, choice)
	case 2:
		asset = OpenDeposit(amount)

	case 3:
		asset = OpenSavingAccount(amount)
	}
	if choice > 3 {
		return asset, param_error{}
	} else {
		return asset, nil, Description(asset)
	}

}
func main() {
	var myProducts []Product
	//Здорово, что я вынес это. Но если вдруг ввод будет не верный, то эти функции всё равно попытаются выполниться и будет ошибочка. Так что божно выбор этот тоже вынести в отдельную функцию и добавить сюда вывод ошибки. И эти функции пусть выполняются если ошибка равна nil
	fmt.Printf("Что хотите сделать?\n 1 - получить новый продукт\n")
	fmt.Scanln(&choice)
	switch choice {
	case 1:

		asset, err = Choice()
		myProducts = append(myProducts, asset)
	default:
		fmt.Println(param_error{})
	}
	fmt.Println(Choice())

}
