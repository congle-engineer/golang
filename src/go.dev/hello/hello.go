package main

import (
	"fmt"
	// "rsc.io/quote"
	"congle/greetings"

	"log"
)

func main() {
	// fmt.Println("Hello World!")
	// fmt.Println(quote.Go())

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"CongLe", "PhuongHuynh", "KhoiLe"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
