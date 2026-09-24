package main

import "fmt"

// потом изучить мапы, и возможно они больше подойдут здесь чем простые пременные

func main() {
	// Ингридиенты в граммах по рецепту
	var cottageCheese float64 = 180 // творог
	var cucumber float64 = 120      // огурец
	var radish float64 = 60         // редис
	var herbs float64 = 15          // зелень
	var sourCream float64 = 20      // сметана
	var ingredients []float64 = []float64{
		cottageCheese,
		cucumber,
		radish,
		herbs,
		sourCream,
	}

	fmt.Println("Граммовка до изменения")
	fmt.Println()

	// Вот тут чтобы самостоятельно не расписывать я мог бы сделать название - ключ, значиние - граммовка. И пройтись циклом по массиву
	fmt.Printf("Творог: %.f\n", ingredients[0])
	fmt.Printf("Огурец: %.f\n", ingredients[1])
	fmt.Printf("Редис: %.f\n", ingredients[2])
	fmt.Printf("Зелень: %.f\n", ingredients[3])
	fmt.Printf("Сметана: %.f\n", ingredients[4])

	// Ингридинты, граммовку которых я изменяю
	var modified float64 = 470

	ratio := modified / cottageCheese

	for i := range ingredients {
		ingredients[i] *= ratio
	}

	fmt.Println()
	fmt.Println()
	fmt.Println("Граммовка после изменения")
	fmt.Println()
	fmt.Printf("Творог: %.f\n", ingredients[0])
	fmt.Printf("Огурец: %.f\n", ingredients[1])
	fmt.Printf("Редис: %.f\n", ingredients[2])
	fmt.Printf("Зелень: %.f\n", ingredients[3])
	fmt.Printf("Сметана: %.f\n", ingredients[4])
}
