package main

import (
	"fmt"
	"github.com/emirpasic/gods/sets/treeset"
)

func main() {
	set := treeset.NewWithIntComparator()
	set.Add()
	set.Add(1)
	set.Add(2)
	set.Add(2, 3)
	set.Add()
	set.Add(6)
	set.Add(4)
	if actualValue := set.Empty(); actualValue != false {
		fmt.Printf("Got %v expected %v", actualValue, false)
	}
	if actualValue := set.Size(); actualValue != 3 {
		fmt.Printf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, expectedValue := fmt.Sprintf("%d%d%d", set.Values()...), "12346"; actualValue != expectedValue {
		fmt.Printf("Got %v expected %v", actualValue, expectedValue)
	}

	fmt.Println(set.Values()...)
}
