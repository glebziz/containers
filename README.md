# Containers

[![Test](https://github.com/glebziz/containers/actions/workflows/test.yml/badge.svg)](https://github.com/glebziz/containers/actions/workflows/test.yml)
[![Coverage](https://codecov.io/gh/glebziz/containers/branch/master/graph/badge.svg?token=4FRAIEVYAL)](https://codecov.io/gh/glebziz/containers/)
[![Go Reference](https://pkg.go.dev/badge/github.com/glebziz/containers.svg)](https://pkg.go.dev/github.com/glebziz/containers)

Containers are a library of data structures that use a pool of nodes to minimize memory allocation during operation.

## Install

```shell
go get -u github.com/glebziz/containers
```

## Structures

All structures in the library use a pool of nodes.
For more information see [node](https://github.com/glebziz/containers/internal/node) package.

### List

The doubly linked list with a pool of nodes and thread safety.
Supports index get and insert operations with `O(n)` complexity.

```go
package main

import (
	"fmt"
	
	"github.com/glebziz/containers/list"
)

func main() {
	l := list.NewPresized[string](3)

	l.PushBack("World")
	l.PushFront("Hello")
	l.PushAfter(1, "!")

	for it := l.Iter(); it.Next(); {
		fmt.Print(it.Val(), " ")
	}
}
```

#### Benchmarks

Benchmarks for a node pooled list versus a standard `list.List`.

```
PushBack/std_list                       100000000           73.68 ns/op         56 B/op         1 allocs/op
PushBack/list_with_pool                 100000000           17.48 ns/op         32 B/op         0 allocs/op
PushBack/list_with_presized_pool        100000000           16.11 ns/op         24 B/op         0 allocs/op

PopBack/std_list                        100000000           17.49 ns/op          0 B/op         0 allocs/op
PopBack/list_with_pool                  100000000           11.43 ns/op          0 B/op         0 allocs/op
PopBack/list_with_presized_pool         100000000           11.28 ns/op          0 B/op         0 allocs/op

Iter/std_list                           100000000           2.584 ns/op          0 B/op         0 allocs/op
Iter/list_with_pool                     100000000           1.240 ns/op          0 B/op         0 allocs/op
Iter/list_with_presized_pool            100000000           1.237 ns/op          0 B/op         0 allocs/op
```

### Ordered map

An ordered map with doubly linked list for order.
Supports thread safety storage, loading and deletion operations with `O(1)` complexity.

```go
package main

import (
	"fmt"
	
	"github.com/glebziz/containers/omap"
)

func main() {
	l := omap.NewPresized[int, string](3)

	l.Store(1, "Hello")
	l.Store(2, "World")
	l.Store(3, "!")

	for it := l.Iter(); it.Next(); {
		fmt.Print(it.Val(), " ")
	}
}
```

#### Benchmarks

Benchmarks for ordered map.

```
Store/ordered_map                       100000000           213.5 ns/op         87 B/op         0 allocs/op
Store/presized_ordered_map              100000000           165.3 ns/op         51 B/op         0 allocs/op

Delete/ordered_map                      100000000           181.4 ns/op          0 B/op         0 allocs/op
Delete/presized_ordered_map             100000000           188.0 ns/op          0 B/op         0 allocs/op

Load/ordered_map                        100000000           98.08 ns/op          0 B/op         0 allocs/op
Load/presized_ordered_map               100000000           88.05 ns/op          0 B/op         0 allocs/op

Iter/ordered_map                        100000000           1.231 ns/op          0 B/op         0 allocs/op
Iter/presized_ordered_map               100000000           1.273 ns/op          0 B/op         0 allocs/op
```

## License

[MIT](https://choosealicense.com/licenses/mit/)