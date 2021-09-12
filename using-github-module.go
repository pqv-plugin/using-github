package main

import (
	"fmt"

	hello "github.com/pqv-plugin/using-modules-go"
)

func main() {
	fmt.Println("Hello, playground")
	fmt.Println(hello.Hello())
	fmt.Println(hello.Proverb())
	fmt.Println(hello.Mala())
}
