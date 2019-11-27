package greetingspackage

import "fmt"

var MagicNumber int

func GopherGreetings() {
	fmt.Println("I'm printing that msg from GopherGreetings() func")
	printGreetingUnexported()
}

func init() {
	MagicNumber = 108
}
