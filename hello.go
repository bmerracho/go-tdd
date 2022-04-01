package main

import (
	"fmt"
)

const enPrefix = "Hello, "
const jpPrefix = "Konnichiwa, "
const krPrefix = "Annyeonghaseyo, "

func Hello(name, lang string) string {
	if len(name) == 0 {
		name = "World"
	}

	prefix := ""

	switch lang {
	case "korean":
		prefix = krPrefix
	case "japanese":
		prefix = jpPrefix
	default:
		prefix = enPrefix
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("", ""))
}
