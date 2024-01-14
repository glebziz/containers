package omap_test

import (
	"fmt"

	"github.com/glebziz/containers/omap"
)

func ExampleOMap_Iter() {
	l := omap.New[int, string]()

	l.Store(1, "Hello")
	l.Store(2, "World")
	l.Store(3, "!")

	for it := l.Iter(); it.Next(); {
		fmt.Print(it.Val(), " ")
	}

	// Output: Hello World !
}
