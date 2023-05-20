# typa
typa is a Go library for type assertion of interface{}. You can assert type of interface{}, converting type in some cases.

## Instllation
You can install with go get.
```
go get github.com/tomatod/typa
```

## Import and Example
```
package main

import (
  "github.com/tomatod/typa"
  "fmt"
)

func main() {
  var origin interface{}  = 12345
  result, ok := typa.Float64(origin)
  if !ok {
    fmt.Println("Convert error")
    return
  }
  fmt.Println("Value of float64 type:", result)
}
```

## Functions
### Convert to specific type
You can safely convert interface{} to specific type. If argument cannot be convertd, second return value is false.

For example, if you give 100 internally with int type to the bellow Int8(), you can get 100 with int8 type.
```
func Int8(i interface{}) (int8, bool)
```

```
func Int16(i interface{}) (int16, bool)
```

```
func Int32(i interface{}) (int32, bool)
```

```
func Int64(i interface{}) (int64, bool)
```

```
func Int(i interface{}) (int, bool)
```

```
func Float32(i interface{}) (float32, bool)
```

```
func Float64(i interface{}) (float64, bool)
```

```
func Uint8(i interface{}) (uint8, bool)
```

```
func Uint16(i interface{}) (uint16, bool)
```

```
func Uint32(i interface{}) (uint32, bool)
```

```
func Uint64(i interface{}) (uint64, bool)
```

```
func Uint(i interface{}) (uint, bool)
```

```
func String(i interface{}) (string, bool)
```

```
func Bool(i interface{}) (bool, bool)
```

### Convert to other type
```
func InternalTypeConvert(i interface{}, t reflect.Type) (interface{}, bool)
```
You can convert to types that aren't in the above list. Similarly, If argument cannot be convertd, second return value is false.

For example converting to byte type,
```
package main

import (
  "github.com/tomatod/typa"
  "fmt"
  "reflect"
)

func main() {
  var origin interface{}  = 123
  result, _ := typa.InternalTypeConvert(origin, reflect.TypeOf((byte)(0)))
  var wanted byte = result.(byte)
  fmt.Println("Value of byte type:", wanted)
}
```
