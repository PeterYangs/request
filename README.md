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

	req, err := client.R().Get("https://www.925g.com")

	if err != nil {

		fmt.Println(err)

		return
	}

	fmt.Println(req.Body().Content().ToString())

}
```