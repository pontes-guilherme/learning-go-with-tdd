package main

import "fmt"

func main() {
	fmt.Println(Hello("Go", ""))
}

func Hello(name, language string) string {
	var greeting string

	if name == "" {
		switch language {
		case "Portuguese":
			name = "Mundo"
		case "English":
			name = "World"
		default:
			name = "World"
		}
	}

	switch language {
	case "Portuguese":
		greeting = "Olá"
	case "English":
		greeting = "Hello"
	default:
		greeting = "Hello"
	}

	return fmt.Sprintf("%s, %s", greeting, name)
}
