package list_test

import (
	"fmt"

	"github.com/glebziz/containers/list"
)

func ExampleNewPresized() {
	l := list.NewPresized[int](10)
	l.PushBack(1)
	fmt.Println(l.Back())
	l.PopFront()
	fmt.Println(l.Front())

	// Output:
	// 1
	// 0
}

func ExampleList_Iter() {
	l := list.New[string]()

	l.PushBack("World")
	l.PushFront("Hello")
	l.PushAfter(1, "!")

	for it := l.Iter(); it.Next(); {
		fmt.Print(it.Val(), " ")
	}

	// Output: Hello World !
}
