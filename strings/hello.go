package strings

import (
	"fmt"
)

const enPrefix = "Hello, "
const jpPrefix = "Konnichiwa, "
const krPrefix = "Annyeonghaseyo, "

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case "korean":
		prefix = krPrefix
	case "japanese":
		prefix = jpPrefix
	default:
		prefix = enPrefix
	}

	return
}
func main() {
	fmt.Println(Hello("", ""))
}
