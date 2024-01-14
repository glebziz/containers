package list_test

import (
	"fmt"

	"github.com/glebziz/containers/list"
)

func ExampleNew() {
	l := list.New[string]()

	l.PushBack("World")
	l.PushFront("Hello")
	l.PushAfter(1, "!")

	for it := l.Iter(); it.Next(); {
		fmt.Print(it.Val(), " ")
	}

	// Output: Hello World !
}
