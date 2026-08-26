package main

import "fmt"

func RepeatGetNewProduct() {
	fmt.Println("Попробуйте ещё раз!")
}
func (c int)GetNewProduct()  {

	switch choice {
	case 1, 2, 3, 4, 5, 6:
		switch choice {
		case 1:
			return GetCredit
		case 2:
			return Save
		case 3:
			return Inshurance
		case 4:
			return GetDebitCard
		case 5:
			return GetCreditCard
		case 6:
			return Back
		}
	default:
		return RepeatGetNewProduct, parem_error{}
	}
}
