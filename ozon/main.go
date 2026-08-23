package main


import(
  "fmt"
)
func percentYear (m int, a, es float64) float64{
return(((( a-es)/float64(m))*12)/a)*100
}

func main(){
  Expenditure := 29672.0 //расход
  Canceled := 23015.0 //погашено
  EstimateLiability :=  Expenditure-Canceled// рассчетный долг
  AllocatedLimit := 50000.0 //предоставленный лимит
  ActualDebt := 7959.81
  fmt.Println("Остаток лимита должен быть:", AllocatedLimit-EstimateLiability)
  fmt.Println("Остаток долга должен быть:", EstimateLiability)
  fmt.Println("По факту же долг:", ActualDebt)
  fmt.Println("Разница между ними:", ActualDebt-EstimateLiability)
  // это переплата за 3 месяца или 4 пользования 6657₽.
  // проценты
  // на 3 месяца
  fmt.Println("Какая это годовая процентная ставка?")
fmt.Println("Если за 3 месяца:", percentYear(3, ActualDebt, EstimateLiability))
  // на 4 месяца
  fmt.Println("Если за 4 месяца:", percentYear(4, ActualDebt, EstimateLiability))
// добавлб изменения для гита
}