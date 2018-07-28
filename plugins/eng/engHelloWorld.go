package main

import "fmt"

type eng string

func (t *eng) Print(name string) {
	fmt.Println("Hello", name)
}

//export as symbol
var HelloWorld eng
