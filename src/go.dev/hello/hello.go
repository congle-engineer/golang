package main

import (
	"fmt"
	"rsc.io/quote"
	"congle/greetings"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println(quote.Go())

	message := greetings.Hello("CongLe")
	fmt.Println(message)
}
