package main

import (
	"fmt"
	"github.com/PeterYangs/request"
	"time"
)

func main() {

	client := request.NewClient()

	content, err := client.R().Timeout(6 * time.Second).ReTry(1).GetToContent("http://list.com/demo/get.php")

	if err != nil {

		fmt.Println(err)

		return
	}

	fmt.Println(content.ToString())

}
