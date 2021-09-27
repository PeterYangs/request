# request

go的工具集

**安装**

```shell
go get github.com/PeterYangs/tools
```

**快速开始**

```go
package main

import (
	"fmt"
	"github.com/PeterYangs/request"
)

func main() {

	client := request.NewClient()

	content, err := client.R().GetToContent("https://www.baidu.com")

	if err != nil {

		fmt.Println(err)

		return

	}

	fmt.Println(content.ToString())

}

```