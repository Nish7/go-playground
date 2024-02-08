package main

import "fmt"

const englishHelloPrefix = "Hello, "
const SpanishHelloPrefix = "Hola, "
const FrenchHelloPrefix = "Bonjour, "

func hello_world(name string, language string) string {

	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) { // the return value is named variable which is accessible in func scope
	switch language {
	case "Spanish":
		prefix = SpanishHelloPrefix
	case "French":
		prefix = FrenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(hello_world("World", englishHelloPrefix))
}
