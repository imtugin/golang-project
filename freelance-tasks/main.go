package main

// Нужно подправить кое что что дипсик говорит

import (
	"fmt"
)

type Task interface {
	Execute() string
	Price() int
	Difficulty() string
}
type DesignTask struct {
	ProjectName string
	Pages       int
	isComplete  bool
}
type CodeTask struct {
	Language   string
	Hours      int
	isComplete bool
}

type TextTask struct {
	Topic      string
	Words      int
	isComplete bool
}

type DesinTaskSlice []DesignTask

// DesignTask
func (d DesignTask) Execute() string {
	return fmt.Sprintf("Дизайн для %s: %d стр.\n", d.ProjectName, d.Pages)
}

func (d DesignTask) Price() int {
	return d.Pages * 500

}

func (d DesignTask) Difficulty() string {
	if d.Pages <= 3 {
		return "Лёгкая"
	} else if d.Pages <= 8 {
		return "Средняя"
	} else {
		return "Сложная"
	}
}

// CodeTas
func (c CodeTask) Execute() string {
	return fmt.Sprintf("Код на %s: %d часов\n", c.Language, c.Hours)
}

func (c CodeTask) Price() int {
	return c.Hours * 1500
}

func (c CodeTask) Difficulty() string {
	if c.Hours <= 5 {
		return "Лёгкая"
	} else if c.Hours <= 20 {
		return "Средняя"
	} else {
		return "Сложная"
	}
}

// TextTask
func (t TextTask) Execute() string {
	return fmt.Sprintf("Текст про %s: %d слов\n", t.Topic, t.Words)
}

func (t TextTask) Price() int {
	return t.Words * 3
}

func (t TextTask) Difficulty() string {
	if t.Words <= 1000 {
		return "Лёгкая"
	} else if t.Words <= 5000 {
		return "Средняя"
	} else {
		return "Сложная"
	}

}

// Суммарная стоимость всех задач

func TotalPrice(tasks []Task) int {
	var totalPrice int
	for i := 0; i < len(tasks); i++ {
		totalPrice += tasks[i].Price()
	}
	return totalPrice
}

// // Возвращает срез задач с указанной сложностью
func FilterByDifficulty(tasks []Task, diff string) []Task { //Здесь я изменил принимаемые параметры функции с (tasks []Task, diff string) на (tasks []Task). Кажется так легче. Если нет, объясни почему. Хотя моет если я не буду эту функцию на прямую в main вызвать, а использовать её в каком-то методе... Но тогда бы здесь принимался не срез а каждый элемент по отдельности
	var outputTasks []Task
	for i := range tasks {
		if diff == tasks[i].Difficulty() {
			outputTasks = append(outputTasks, tasks[i])
		}
	}
	return outputTasks
}

// Принимает Task, делает type switch и выводит подробное описание для каждого типа (включая название проекта/язык/тему и цену)
func Describe(t Task) {

	fmt.Println()
	fmt.Println()
	fmt.Println("Подробное описание задачи")
	fmt.Println(t.Execute())
	fmt.Printf("Цена: %d рублей.\n", t.Price())
	fmt.Printf("Уровень: %s\n", t.Difficulty())

}

// находит самую дешёвую задачу
func FindCheapest(tasks []Task) Task {
	var cheapestTask Task = tasks[0]
	for i := range tasks {
		if tasks[i].Price() <= cheapestTask.Price() {
			cheapestTask = tasks[i]
		}

	}
	return cheapestTask
}

func (d *DesinTaskSlice) MarkDone() {
	for i := range *d {
		if (*d)[i].isComplete == false {
			(*d)[i].isComplete = true
		}
	}
}
func main() {
	var input int
	var inputStr string
	alltasks := []Task{
		DesignTask{"дизайн альбома", 20, true},
		DesignTask{"дизайн сайта", 10, false},
		DesignTask{"дизайн офиса", 30, false},
		CodeTask{"JS", 3, false},
		CodeTask{"TS", 154, true},
		TextTask{"Важность китайско-кыргызско-узбекской железной дороги для стран Центральной Азии", 4, false},
		TextTask{"Искусственный интеллект - когда судныый день", 2, false},
	}

	var designTasks DesinTaskSlice = DesinTaskSlice{
		DesignTask{"дизайн альбома", 20, true},
		DesignTask{"дизайн сайта", 10, false},
		DesignTask{"дизайн офиса", 30, false},
	}
	fmt.Println("Задачи по дизайну до выполнения:", designTasks)
	designTasks.MarkDone()
	fmt.Println("Задачи по дизайну выполненные:", designTasks)

	for i := range alltasks {
		fmt.Println(i+1, ".", alltasks[i].Execute(), alltasks[i].Difficulty(), "Цена:", alltasks[i].Price(), "рублей")
	}
	fmt.Println()
	fmt.Println()
	fmt.Println("Суммараня стоимость всех задач:", TotalPrice(alltasks))
	fmt.Println()
	fmt.Println()
	fmt.Println("====================================")
	fmt.Println("Вот вам лёгкие задачи")
	for i := range alltasks {
		if alltasks[i].Difficulty() == "Лёгкая" {
			fmt.Println()
			fmt.Println(alltasks[i].Execute(), alltasks[i].Difficulty(), "Цена:", alltasks[i].Price(), "рублей")
			fmt.Println()
		}
	}
	fmt.Println("-------------------------------------")
	fmt.Println("Выбор задач указанной сложности")
	fmt.Println("Какую сложность выбрать? 1 - Лёгкая. 2 - Средняя. 3 - Сложная")
	fmt.Scanln(&input)
	switch input {
	case 1:
		inputStr = "Лёгкая"
	case 2:
		inputStr = "Средняя" // Проблема где-то. Именно задачи средней сложности не выводятся
	case 3:
		inputStr = "Сложная"
	}
	fmt.Println("Вы выбрали", input, ":", inputStr)
	fmt.Println(FilterByDifficulty(alltasks, inputStr))
	fmt.Println()
	fmt.Println()
	fmt.Println("А вот самая дешёвая задача:", FindCheapest(alltasks))
	fmt.Println()
	fmt.Println()
	fmt.Println("Введите задачу о которой вы хотите узнать из списка из списка")
	fmt.Scanln(&input)
	Describe(alltasks[input-1])
}
