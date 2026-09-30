package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {

		got := Hello("Ilya")
		want := "Hello, Ilya"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to world", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) { // функция принимает тестируемую функцию(?), как мы её запускаем(что передаём в неё),и что мы хотим получить

	t.Helper() // необходим что бы сообщить тестовому набору, что этот метод является вспомогательным. При воздникновении ошибки номер строки будет находиться в нашем вызове функции, а не внутри нашей вспомогательной функции теста
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
