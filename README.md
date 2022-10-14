# request

http请求库

**安装**

```shell
go get github.com/PeterYangs/request/v2
```

**快速开始**

```go
package main

import (
	"fmt"
	"github.com/PeterYangs/request/v2"
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
	"github.com/PeterYangs/request/v2"
)

func main() {

	client := request.NewClient()

	r, err := client.R().Get("https://www.baidu.com")

	if err != nil {

		fmt.Println(err)

		return
	}

	content, err := r.Content()

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
	"github.com/PeterYangs/request/v2"
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

Params作为body参数，Query作为url参数

```go
package main

import (
	"fmt"
	"github.com/PeterYangs/request/v2"
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

	content, err := r.Content()

	if err != nil {

		fmt.Println(err)

		return
	}

	fmt.Println(content.ToString())

}
```

**post Multipart提交**

```go
package main

import (
	"fmt"
	"github.com/PeterYangs/request/v2"
	"os"
)

func main() {

	f, e := os.OpenFile("README.md", os.O_CREATE|os.O_RDWR, 0644)

	if e != nil {

		return
	}

	client := request.NewClient()

	ct, err := client.R().Params(map[string]interface{}{"name": "yy", "file": f}).PostMultipartToContent("http://www.yycms.com/post_test")

	if err != nil {

		fmt.Println(err)

		return
	}

	fmt.Println(ct.ToString())

}
```


**header**

```go
package main

import (
	"fmt"
	"github.com/PeterYangs/request/v2"
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
	"github.com/PeterYangs/request/v2"
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
	"github.com/PeterYangs/request/v2"
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
	"github.com/PeterYangs/request/v2"
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
	"github.com/PeterYangs/request/v2"
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

**自定义body**

```go
package main

import (
	"fmt"
	"github.com/PeterYangs/request/v2"
	"strings"
)

func main() {

	client := request.NewClient()
	//设置body后，Params将失效
	content, err := client.R().Body(strings.NewReader("name=123&age=18")).PostToContent("http://list.com/demo/post.php")

	if err != nil {

		fmt.Println(err)

		return
	}

	fmt.Println(content.ToString())

}
```



