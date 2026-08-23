package main

import(
  "fmt"
_"errors"
  )
//ощибки
type params_error struct{}

func (error_object params_error)Error()string{
  return "invalid parametr"
}
func divide (a,b int)(int, error){
  if b <=0{
    return 0, params_error{}
  }else{
    return a/b, nil
  }
}

func main (){
//  obj := params_error{}
//  fmt.Println(obj.Error())
  //fmt.Println(params_error{})
  x, y := 10,5
  fmt.Println(divide(x,y))
  y =0
    fmt.Println(divide(x,y))

}