# Japanese MyNumber validator

```go
package main

import (
	"fmt"

	"github.com/koron/go-mynumber"
)

func main() {
	// true
	fmt.Println(mynumber.ValidateStr("123456789018"))
	// false
	fmt.Println(mynumber.ValidateStr("123456789010"))
}
```

refer: [Ruby - マイナンバーのチェックデジットを計算する](http://qiita.com/qube81/items/fa6ef94d3c8615b0ce64)
