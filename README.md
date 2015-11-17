# Japanese MyNumber validator

```go
package main

import "github.com/koron/go-mynumber"

func main() {
	// true
	println(mynumber.ValidateStr("123456789018"))
	// false
	println(mynumber.ValidateStr("123456789010"))
}
```

refer: [Ruby - マイナンバーのチェックデジットを計算する](http://qiita.com/qube81/items/fa6ef94d3c8615b0ce64)
