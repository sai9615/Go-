package main

import (
	"example.com/myModule"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("myModules: ")
	log.SetFlags(0)

	// message, err := myModule.Hello("Sam")
	names := []string{"Sai", "Sam", "Sol"}
	messages, err := myModule.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Print(message)
	for _, name := range names {
		fmt.Println(messages[name])
	}
}
