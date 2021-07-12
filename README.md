# ezreflect
Easy Reflect tools make reflect more easier to use.

## install
```
go get github.com/ka1hung/ezreflect
```

## sample 1 
list struct field names
```go
package main

import (
	"fmt"

	"github.com/ka1hung/ezreflect"
)

type structA struct {
	AAA string  `tag:"aaa"`
	BBB int     `tag:"bbb"`
	CCC float64 `tag:"ccc"`
	DDD bool    `tag:"ddd"`
}

func main() {
	a := structA{
		AAA: "hello",
		BBB: 123,
		CCC: 0.123,
		DDD: true,
	}
	//list struct field names
	fmt.Println(ezreflect.GetFieldNames(a))

	//list struct field types
	fmt.Println(ezreflect.GetFieldTypes(a))

	//list struct field values
	fmt.Println(ezreflect.GetFieldDataString(a))

	//list struct field tags
	fmt.Println(ezreflect.GetFieldTags(a))
}


```

Hope you like it.