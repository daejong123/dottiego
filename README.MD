## 测试go的新特性 
> go mod

```go
export GO111MODULE=on

go mod init xxxx
```

## 具体库的使用
```go
1. 下载并编译库
go get -u github.com/daejong123/dottiego

2. 示例代码
package main

import (
	"github.com/daejong123/dottiego"
	"fmt"
)

func main() {
	tip := dottiego.SayHi("dottie")
	fmt.Println(tip)
}
```