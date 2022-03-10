package main

import (
	"fmt"
	"github.com/PeterYangs/request/v2"
	"strings"
)

func main() {

	client := request.NewClient()

	content, err := client.R().Body(strings.NewReader("name=123&age=18")).PostToContent("http://list.com/demo/post.php")

	if err != nil {

		fmt.Println(err)

		return
	}

	fmt.Println(content.ToString())

}
