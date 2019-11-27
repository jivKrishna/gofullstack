package greetingspackage

import f "fmt"

func PrintGreetings() {
	f.Println("I'm printing that msg from PrintGreetings() func")
}

func printGreetingUnexported() {
	f.Println("I'm printing that msg from func printGreetingUnexported() func")
}
