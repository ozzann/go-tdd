package main

import "fmt"

const english = "English"
const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello"
const spanishHelloPrefix = "Hola"
const frenchHelloPrefix = "Bonjour"

func Hello(name, lang string) string {

	if name == "" {
		name = "World"
	}
	return greetingPrefix(lang) + ", " + name
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	name := "world"
	fmt.Println(Hello(name, "English"))
}
