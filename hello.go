package main

import "fmt"

const hello = "Hello, "

func Hello(name string) string {
	if len(name) == 0 {
		name = "World"
	}
	return hello + name
}

func main() {
	fmt.Println(Hello(""))
}
