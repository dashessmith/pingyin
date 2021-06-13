# pinyin

内存中的汉字和拼音对照

使用 PH, HP 两个变量
PH : 拼音 -> 汉字
HP : 汉字 -> 拼音

```go
package main

import (
	"fmt"

	"github.com/dashessmith/pinyin"
)

func main() {
	fmt.Printf("%v\n", pinyin.HP["汉"])
}
```
