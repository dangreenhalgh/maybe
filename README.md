# Maybe
Maybe types for Go

## Usage

```go
package main

import (
	"encoding/json"
	"fmt"
	"math"

    "github.com/dangreenhalgh/maybe"
)


func main() {
    someInt := maybe.Int{Some: 4}
    noneInt := maybe.Int{None: true}

	someIntAsJSON, _ := json.Marshal(someInt)
	noneIntAsJSON, _ := json.Marshal(noneInt)

	fmt.Printf("%v\n", string(someIntAsJSON))
	fmt.Printf("%v\n", string(noneIntAsJSON))

	someFloat64 := maybe.NewFloat64(123.456)
	noneFloat64 := maybe.NewFloat64(math.NaN())

	someFloat64AsJSON, _ := json.Marshal(someFloat64)
	noneFloat64AsJSON, _ := json.Marshal(noneFloat64)

	fmt.Printf("%v\n", string(someFloat64AsJSON))
	fmt.Printf("%v\n", string(noneFloat64AsJSON))
}
```

```
4
null
123.456
null
```