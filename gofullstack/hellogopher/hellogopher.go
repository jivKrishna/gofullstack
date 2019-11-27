package main

import "fmt"

func main() {
	fmt.Println("Hello Gopher!")
	const (
		_ = 2010 + iota*4
		hello
		bye
		goodmorning
	)

	fmt.Println(hello, bye, goodmorning)
	fmt.Printf("%-18s %-18v\n", "hello", hello)
	fmt.Printf("%-18s %-18v\n", "bye", bye)
	fmt.Printf("%-18s %-18v\n", "goodmorning", goodmorning)
}
