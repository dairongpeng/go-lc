package main

import (
	"fmt"
	"sync"
)

var once sync.Once
var A string

func main() {
	for i := 0; i < 10; i++ {
		NewABC()
	}
	fmt.Println(A)
}

func NewABC() {
	once.Do(func() {
		fmt.Println("Init...")
		A = "abc"
	})
}
