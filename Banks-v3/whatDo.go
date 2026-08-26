package main

import (
	"fmt"
  "os"
)





func Choise()  {
      	var c int
  repeat:
	fmt.Println("Что хотите сделать?")
	fmt.Println("Получить кредит - 1")
	fmt.Println("Н Сделать перевод - 2")
	fmt.Println("Страхование - 3")
	fmt.Println("Получить дебетовую карту - 4")
	fmt.Println("Получить кредитную карту - 5")
	fmt.Println("Назад- 6")
	fmt.Println("Выйти - 7")
  fmt.Println("Узнать балланс - 8")
  fmt.Scanln(&c)

	switch c {
	case 1, 2, 3, 4, 5,6,7,8:

		switch c {
		case 1:

//		GetNewProduct()
	case 2:
	 MakeTransfer()
	case 3:
	// TopUpYourAccount()
	case 4:
//	View()
	case 5:
 case 6:
      fmt.Println("ошибка выбора")
      goto repeat
  case 7:
      os.Exit(3)
    case 8:
      fmt.Println(balance)
		}
	}
}
