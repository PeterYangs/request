package main

import (
	"fmt"
	"github.com/PeterYangs/request"
)

func main() {

	client := request.NewClient()

	r, err := client.R().Get("https://www.925g.com")

	if err != nil {

		fmt.Println(err)

		return
	}

	content := r.Body().Content().ToString()

	fmt.Println(content)

}
