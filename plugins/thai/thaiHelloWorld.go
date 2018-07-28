package main

import "fmt"

type thai string

func (t *thai) Print(name string) {
	fmt.Println("สวัสดีครับ", name)
}

//export as symbol
var HelloWorld thai
