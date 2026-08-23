package main

import (
	"fmt"
)

type Product interface {
	Income() float64
	IncomeMonth() float64
}
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

func Description(a Product) {
	fmt.Println(a)
	fmt.Println("Доход за весь период:", a.Income())
	fmt.Println("Доодность за месяц:", a.IncomeMonth())
}

// Нужно добавить чтобы при выводе цифрового актива мне предлагался вклад для сравнения на тот же срок
func ForComparsion(c int) {
	fmt.Println("Вот для сравнения информация по вкладу на аналогичный срок")
}

func OpenDeposit(a float64, choice int) Deposit {

	if choice < 3 {
		return Deposit{
			Amount: a,
			Term:   choice,
			Annual: 11,
		}
	} else if 6 < choice && choice < 10 {
		return Deposit{
			Amount: a,
			Term:   choice,
			Annual: 12.2,
		}
	} else if 9 < choice && choice < 12 {
		return Deposit{
			Amount: a,
			Term:   choice,
			Annual: 12.1,
		}
	} else {
		switch choice {
		case 4:
			return Deposit{
				Amount: a,
				Term:   choice,
				Annual: 13.6,
			}
		case 5:
			return Deposit{
				Amount: a,
				Term:   choice,
				Annual: 12.3,
			}
		case 6:
			return Deposit{
				Amount: a,
				Term:   choice,
				Annual: 13,
			}
		case 12:
			return Deposit{
				Amount: a,
				Term:   choice,
				Annual: 13,
			}
		case 24:
			return Deposit{
				Amount: a,
				Term:   choice,
				Annual: 11.8,
			}
		default:
			fmt.Println("Будет ошибка")
		}

	}
}

func main() {
	var myProducts []Product
	var choice int
	var amount float64
metka:
	fmt.Printf("Что хотите сделать?\n 1 - получить новый продукт\n")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		fmt.Println("Введите сумму")
		fmt.Scanln(&amount)
	metka1:
		fmt.Printf("Цифровой актив - 1\n Вклад - 2\n Накопительный счет - 3\n")
		// Можно добавить еще вариант с кредитами но не буду усложнять
		fmt.Scanln(&choice)

		switch choice {
		case 1:

			myProducts = append(myProducts, DigitalAssets{
				Amount: amount,
				Term:   choice,
				Annual: 15.2,
				Tax:    13,
			})
		metka4:
			fmt.Println("На какой срок? Доступно 1, 3, 6, 9 или 12 месяцев")
			fmt.Scanln(&choice)
			switch choice {

			case 1:
				myProducts = append(myProducts, DigitalAssets{
					Amount: amount,
					Term:   choice,
					Annual: 15.2,
					Tax:    13,
				})

			case 3:
				myProducts = append(myProducts, DigitalAssets{
					Amount: amount,
					Term:   choice,
					Annual: 15.7,
					Tax:    13,
				})

			case 6:
				myProducts = append(myProducts, DigitalAssets{
					Amount: amount,
					Term:   choice,
					Annual: 17,
					Tax:    13,
				})

			case 9:
				myProducts = append(myProducts, DigitalAssets{
					Amount: amount,
					Term:   choice,
					Annual: 17.5,
					Tax:    13,
				})

			case 12:
				myProducts = append(myProducts, DigitalAssets{
					Amount: amount,
					Term:   choice,
					Annual: 18,
					Tax:    13,
				})
			default:

				fmt.Println("ввод не верный")
				goto metka4

			}

			asset := myProducts[len(myProducts)-1]

			Description(asset)
			fmt.Println("Реальная годовая процентная ставка после налогового вычета:", (asset.IncomeMonth()*1200.0)/amount)
			ForComparsion(choice)
		case 2:
			fmt.Println("На какой срок? Доступно  от 1 до 12, или 24 месяца")
			fmt.Scanln(&choice)

			asset := myProducts[len(myProducts)-1]

			Description(asset)

		case 3:
		metka2:
			fmt.Println("Начисление на минимальную сумму - 1\n На ежедневный остаток - 2")
			fmt.Scanln(&choice)
			switch choice {
			case 1:
				myProducts = append(myProducts, SavingAccount{
					Amount:       amount,
					DailyBalance: false,
					Annual:       12,
					Term:         1,
				})
			case 2:
				myProducts = append(myProducts, SavingAccount{
					Amount:       amount,
					DailyBalance: true,
					Annual:       11.5,
					Term:         1,
				})

			default:

				fmt.Println("ввод не верный")
				goto metka2

			}
			asset := myProducts[len(myProducts)-1]

			fmt.Println("Доход за месяц:", asset.Income())
		default:

			fmt.Println("ввод не верный")
			goto metka1

		}
	default:
		fmt.Println("ввод не верный")
		goto metka
	}

}
