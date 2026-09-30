package main

import "fmt"

//здесь я посчитаю размеры ящика под духовой шкаф

// ширина ящика
// нужно из ширины проёмма вычесть 127*2 (зазор на направляющие)
func Width(wP int) int { // Функция принимает ширину проёма
	n := 127 // зазор на направляющие
	return wP - 2*n
}

func Depth(oD int) int {
	n := 3 // зазор сзади
	return oD - n
} // Глдубина ящика. Функция принимает глубину проёма
func main() {
	openingWidth := 570 // ширина проёма под ящик
	openingDepth := 532 // глубина проёма под ящик
	fmt.Println("Габариты ящика")
	fmt.Println("по ширине:", Width(openingWidth))
	fmt.Println("по глубине:", Depth(openingDepth))
	fmt.Println("по высоте:", 130)
	fmt.Println()
	fmt.Println("Детали")
	fmt.Println("Две боковины")
	fmt.Println(Depth(openingDepth), "на", 130)
	fmt.Println("Дно и шифлятки зависят от толщины ДСП. Ящик должен получиться по заданным габаритам")
}
