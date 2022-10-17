package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Holla, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix
	switch lang {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", "English"))
}
