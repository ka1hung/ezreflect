# ezreflect
Easy Reflect tools make reflect more easier to use.

## install
```
go get github.com/ka1hung/ezreflect
```

## sample1: list struct field infomations
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

### sample2: operate the struct fields
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

type structB struct {
	AAA string `tag:"aaa"`
	BBB int    `tag:"bbb"`
}

func main() {
	a := structA{
		AAA: "hello",
		BBB: 123,
		CCC: 0.123,
		DDD: true,
	}

	//struct field copy 
	b := structB{}
	ezreflect.FieldCopy(&a, &b)
	fmt.Printf("%+v\n", b) //{AAA:hello BBB:123}

	//struct field copy by names
	bb := structB{}
	ezreflect.FieldCopyByNames(&a, &bb, []string{"AAA"})
	fmt.Printf("%+v\n", bb) //{AAA:hello BBB:0}

	//struct field parse by map data
	m := map[string]string{
		"AAA": "hahaha",
		"BBB": "999",
		"CCC": "-1.002",
		"DDD": "false",
	}
	ezreflect.FieldParseFromString(&a, m)
	fmt.Printf("%+v\n", a) //{AAA:hahaha BBB:999 CCC:-1.002 DDD:false}
}
```

Hope you like it.