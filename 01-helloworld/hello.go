package main

import "fmt"

func Hello(msg string) string {
	if msg == "" {
		msg = "Hello, World"
	}
	return msg
}

func main() {
	fmt.Println(Hello("Hello, World"))
}
