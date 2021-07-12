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

import(
    github.com/ka1hung/ezreflect
)

type structA struct{
    AAA string
    BBB int
    CCC float64
    DDD bool
}

func main(){
    a:=structA{}
    ns:=ezreflect.GetFieldNames(&a)
    fmt.Println(ns)
}

```

Hope you like it.