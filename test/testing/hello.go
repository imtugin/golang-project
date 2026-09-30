package main

import "fmt"

func Hello(name string) string {
	message := fmt.Sprint("Hello, ", name)
	return message
}
func main() {
	fmt.Println(Hello("world"))

}
