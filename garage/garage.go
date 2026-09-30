package main

import (
	"fmt"
	"math"
)

type Material struct {
	width,
	lenght,
	S float64
}
type Roof struct {
	width,
	lenght,
	S float64
}

func numbersOfRolls(r Roof, f Material) (numbersCeil, quantity float64) {
	overlap := 0.1
	// количество кусков
	numbers := (math.Ceil(r.lenght/(f.width-overlap)) * r.width / f.lenght)
	numbersCeil = math.Ceil(numbers)
	quantity = numbersCeil - numbers

	return
}

// Рассчёт ширины
func CalculateWidth(f *Material) {
	f.width = f.lenght / f.S
}
func S(r *Roof) float64 {
	return r.lenght * r.width
}

func Description(f Material, r Roof) {
	fmt.Println("Для материала длиной", f.lenght, "и площадью", f.S, "метров")
	fmt.Println()
	fmt.Println("Ширина рулона будет", f.width, "метр")
	n, q := numbersOfRolls(r, f)
	fmt.Println("Необходимое количество рулонов на крышу:", n, "штуки")
	fmt.Printf("Остаток: %.2f метра\n", q)
}
func main() {
	var felt10 Material = Material{lenght: 10, S: 10}
	var felt15 Material = Material{lenght: 15, S: 15}
	var roof Roof = Roof{lenght: 7.3, width: 4.1}
	CalculateWidth(&felt10)
	CalculateWidth(&felt15)
	fmt.Println("Площаль крыши:", S(&roof))
	fmt.Println("__________________________________________")
	Description(felt10, roof)
	fmt.Println("==========================================")
	Description(felt15, roof)
}
