package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	message := fmt.Sprint(englishHelloPrefix, name)
	return message
}
func main() {
	fmt.Println(Hello("world"))

}
