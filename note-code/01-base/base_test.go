package main

import (
	"fmt"
	"testing"
)

func TestExist(t *testing.T) {
	arr := []int{1, 4, 6, 27, 33, 66, 97}
	target := 31
	fmt.Println(exist(arr, target))
}
