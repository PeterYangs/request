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
