# request

http请求库

**安装**

```shell
go get github.com/PeterYangs/request
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


**get**
```go
package main

import (
	"fmt"
	"github.com/PeterYangs/request"
)

func main() {

	client := request.NewClient()

	r, err := client.R().Get("https://www.baidu.com")

	if err != nil {

		fmt.Println(err)

		return
	}

	
	content, err := r.Body().Content()

	if err != nil {

		fmt.Println(err)

		return
	}

	fmt.Println(content.ToString())

}
```
或者直接读取结果
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

**post**
```go
package main

import (
	"fmt"
	"github.com/PeterYangs/request"
)

func main() {

	client := request.NewClient()

	r, err := client.R().Params(map[string]interface{}{
		"list":  "123",
		"array": []string{"123", "456"},
		"form": map[string]interface{}{
			"name":  "yy",
			"email": "904801074@qq.com",
			"age":   10,
			"order": []string{"1", "2", "3"},
		},
	}).Post("https://www.baidu.com")

	if err != nil {

		fmt.Println(err)

		return
	}

	content, err := r.Body().Content()

	if err != nil {

		fmt.Println(err)

		return
	}

	fmt.Println(content.ToString())

}
```
get也可以设置Params,会添加到url地址上


**header**
```go
package main

import (
	"fmt"
	"github.com/PeterYangs/request"
)

func main() {

	//全局设置header
	client := request.NewClient().Header(map[string]string{"test": "demo"})

	//请求设置header
	content, err := client.R().Header(map[string]string{"test": "demo2"}).GetToContent("http://list.com/demo/header.php")

	if err != nil {

		fmt.Println(err)

		return
	}

	fmt.Println(content.ToString())

}
```

**proxy**
```go
package main

import (
	"fmt"
	"github.com/PeterYangs/request"
)

func main() {

	client := request.NewClient().Proxy("http://127.0.0.1:4780")

	content, err := client.R().GetToContent("https://www.google.com.hk/")

	if err != nil {

		fmt.Println(err)

		return
	}

	fmt.Println(content.ToString())

}
```


**download**

下载
```go
package main

import (
	"fmt"
	"github.com/PeterYangs/request"
)

func main() {

	url := "http://list.com/demo/demo.zip"

	client := request.NewClient()

	err := client.R().Download(url, "123.zip")

	if err != nil {

		fmt.Println(err)
	}
}
```

**reTry**

重试
```go
package main

import (
	"fmt"
	"github.com/PeterYangs/request"
)

func main() {

	client := request.NewClient()

	content, err := client.R().ReTry(1).GetToContent("http://list.com/asdjk.php")

	if err != nil {

		fmt.Println(err)

		return
	}

	fmt.Println(content.ToString())

}

```

**upload**
```go
package main

import (
	"fmt"
	"github.com/PeterYangs/request"
)

func main() {

	client := request.NewClient()

	content, err := client.R().Upload("http://list.com/demo/get.php", "README.md")

	if err != nil {

		fmt.Println(err)

		return

	}

	fmt.Println(content.ToString())

}
```




