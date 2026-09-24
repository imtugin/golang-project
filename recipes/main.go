package main

import (
	"fmt"
	"strconv"
)

type Recipes interface {
	List()
	Change()
}

type Ingredients map[string]float64

// изучить мапы, и возможно они больше подойдут здесь чем простые пременные
// Изучить запись файл. Создать файл, чтение из файла (или файлы с рецептами)

func List(r Ingredients) {
	fmt.Println()
	fmt.Println("Список ингридиентов")
	fmt.Println("-------------------")
	for i := range r {
		fmt.Printf("%s: %.f\n", i, r[i])
	}
	fmt.Println()
	fmt.Println()
}

func Change(r *Ingredients) {
	var inputIngredients, inputGram string
	fmt.Println("Введите название ингридиента, который нужно изменить")
	fmt.Scanln(&inputIngredients)
	slice := *r
	for i := range slice {
		if i == inputIngredients {
			fmt.Println("Введите нужное количество грамм:")
			fmt.Scanln(&inputGram)
			val, err := strconv.ParseFloat(inputGram, 64)
			if err != nil {
				fmt.Println("Не удалось конвенртировать ввёдённое количество", err)
				return
			} else if val == 0 {
				fmt.Println("Вы ввели 0", err)
			}
			ratio := val / slice[i]
			for i := range slice {
				slice[i] *= ratio
			}

		}
	}

}
func main() {
	recipes := Ingredients{
		"Творог":  180,
		"Огурец":  120,
		"Редис":   60,
		"Зелень":  15,
		"Сметана": 20,
	}

	List(recipes)
	Change(&recipes)
	fmt.Println("После")
	List(recipes)

}
