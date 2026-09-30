package main

import (
	"fmt"
)

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, language string) string {

	if name == "" {
		name = "world"
	}
	return greetingLanguage(language) + name
}

func greetingLanguage(lang string) string {
	switch lang {
	case spanish:
		return spanishHelloPrefix
	case french:
		return frenchHelloPrefix
	default:
		return englishHelloPrefix
	}

}
func main() {
	fmt.Println(Hello("world", ""))

}
