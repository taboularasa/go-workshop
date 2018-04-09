package main

import "fmt"

func main() {

	eric := Name("Eric")
	david := Name("David")

	// Normal usage of receiver methods
	eric.Print()
	david.Print()

	// You can also call a receiver method directly
	Name.Print(eric)
	Name.Print(david)

	// Since its just a method, you can reference it and pass it around
	printName := Name.Print
	printName(eric)
	printName(david)
}

type Name string

func (n Name) Print() {
	fmt.Println(n)
}
