package main

import (
	"fmt"
	"github.com/PeterYangs/request/v2"
)

func main() {

	c := request.NewClient()

	c.Header(map[string]string{
		//"Cookie":          "BAIDUID=FF6DF67FD4291F501DA92CEACF0459C8:FG=1; BIDUPSID=FF6DF67FD4291F501DA92CEACF0459C8; PSTM=1626333985; __yjs_duid=1_63d707a7a111f60675e79c7ed8bf1c5b1626334002625; BD_UPN=19314753; BAIDUID_BFESS=FF6DF67FD4291F501DA92CEACF0459C8:FG=1; BDUSS=WRqUGhpNDNIQllTRk5DRm9IejRURXBDQldGM1AyQlhGMnFMRzVxeXFheFRVV05pRUFBQUFBJCQAAAAAAAAAAAEAAADBx1PysOvi-bexu6pVAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFPEO2JTxDtiR; BDUSS_BFESS=WRqUGhpNDNIQllTRk5DRm9IejRURXBDQldGM1AyQlhGMnFMRzVxeXFheFRVV05pRUFBQUFBJCQAAAAAAAAAAAEAAADBx1PysOvi-bexu6pVAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFPEO2JTxDtiR; H_PS_645EC=501aF16mq1mtBWc4SZsY4ZzDfFm7LQVWTKzTZlwVSB%2B%2FsPVWS7k7TiFtXqID%2F7LtPPOVoO5SG%2Fi9; BDORZ=FFFB88E999055A3F8A630C64834BD6D0; COOKIE_SESSION=437479_0_9_9_10_0_1_0_9_0_0_0_437479_0_2_0_1648521486_0_1648521484%7C9%230_0_1648521484%7C1; ab_sr=1.0.1_NjZiZTBjZjEwNDhmODQxODE3NDQ0OTZhZjVjNjBmNDMxNTYyN2EyN2JmNGRjZmY0NGI3M2MxZGVjNzVmODMyNzc2ZjQ0MGVjNmZmZjRjNjFlNzYzODhkNWZkMDIxY2Y4OTEzMDU4NzcxOWIzMDc0NzE3NTU5NTljYzVjMWU5YzA0NTRhNzVkODVkYjc5MzU3MjVkNTk4MTk3ZGZlNDRhOTA2OGMxYzkxY2YwMWJlZjI0ZTRhZWFlZjYwNWVmOWMx; BDRCVFR[usl9kXwRIYc]=mk3SLVN4HKm; BD_HOME=1; H_PS_PSSID=; BA_HECTOR=2k8k040ga10h2hal6d1h44sd50q",
		"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
		"Accept-Encoding": "gzip, deflate, br",
		"Accept-Language": "zh-CN,zh;q=0.9",
		"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.87 Safari/537.36 SE 2.X MetaSr 1.0",
	})

	ct, err := c.R().GetToContent("https://www.baidu.com/s?wd=王者荣耀")

	if err != nil {

		fmt.Println(err)

		return
	}

	fmt.Println(ct.ToString())
	fmt.Println(ct.Header())

}
