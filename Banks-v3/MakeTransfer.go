package main

import (
  "fmt"
 _ "errors"
  )

  var sumTransfer float64 = 5000



func Transfer(b *float64) {
  *b-=sumTransfer
	fmt.Println("Перевод осуществлён как будто")
}

var obj param_error

type param_error struct{}

func(object_error param_error)Error(){
  return "no money!"
}

  func CheckBalance(b, t float64)error{
    if t<b{
      return obj.Error
    }else{
      return nil
    }
  }

func MakeTransfer()  {
  
  
	fmt.Println("Кому перевести")

  if CheckBalance !=nil{
      Transfer(&balance)
  }else{
    fmt.Println("пополните счет")
  }
}
