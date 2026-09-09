package main

import (
	"fmt"
	"time"
)

type month struct {
	Month      time.Month
	numberDays int
	hours      int
	Shift      shifts
}

type shifts struct {
	First  int
	Second int
	Third  int
}

type salaryHour struct {
	standart             float64
	increasedBy20Percent float64
}

func Salary(m month, o float64) (float64, float64, float64, float64, float64, float64) {

}
func Description(m month) {
	standartH, incr20, first, second, third, sal = Salary(m)
	fmt.Printf("За стандартный час:%0.f, За ночной час: %0.f\n", salaryHour.standart, salaryHour.increasedBy20Percent)
	fmt.Println("По сменам:")
	fmt.Printf("Первая: %0.f, Вторая: %0.f, Третья: %0.f\n", salaryFirst, salarySecond, salaryThird)
	fmt.Printf("В общем: %0.f\n", salary)
}
func main() {
	var septermber month
	septermber.Month = time.September
	septermber.Shift = shifts{
		Second: 17,
		Third:  5,
	}
	septermber.numberDays = septermber.Shift.First + septermber.Shift.Second + septermber.Shift.Third

	var oklad float64 = 74000.0
	septermber.hours = 8 * septermber.numberDays
	salaryHour.standart = oklad / float64(septermber.hours)
	salaryHour.increasedBy20Percent = salaryHour.standart + salaryHour.standart*0.2
	var salaryFirst float64 = float64(septermber.Shift.First) * salaryHour.standart * 8
	var salarySecond = float64(septermber.Shift.Second)*salaryHour.standart*1.5 + float64(septermber.Shift.Second)*salaryHour.increasedBy20Percent*7.5
	var salaryThird = float64(septermber.Shift.Third)*salaryHour.standart*1 + float64(septermber.Shift.Third)*salaryHour.increasedBy20Percent*7
	var salary = salaryFirst + salarySecond + salaryThird

}
