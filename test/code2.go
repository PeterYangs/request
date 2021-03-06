package main

import (
	"fmt"
	"github.com/PeterYangs/request/v2"
	"time"
)

func main() {

	//gzip.

	client := request.NewClient()

	rsp, err := client.R().Header(map[string]string{
		"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
		"Accept-Encoding": "gzip, deflate, br",
		"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
		"Cookie":          "BIDUPSID=DB17BDBD53EF4578A2AF0F348C53571D; PSTM=1590714377; sug=0; sugstore=0; ORIGIN=0; bdime=20100; MCITY=-218%3A; BDUSS=3FJbkMxLWdtWEhlOUNsd1ZCMTBBaTRkfmdTT2hFfkt2SX50a1dmaEJ4cFJYRzVnRVFBQUFBJCQAAAAAAAAAAAEAAABJzpi8wrfNvjdBTHhnAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFHPRmBRz0Zgc; BDUSS_BFESS=3FJbkMxLWdtWEhlOUNsd1ZCMTBBaTRkfmdTT2hFfkt2SX50a1dmaEJ4cFJYRzVnRVFBQUFBJCQAAAAAAAAAAAEAAABJzpi8wrfNvjdBTHhnAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAFHPRmBRz0Zgc; __yjs_duid=1_7451f69d888d4264ca837eed9efbda1d1620353058581; COOKIE_SESSION=88534_0_9_7_13_2_1_0_7_2_0_0_591342_0_1_0_1626408036_0_1626408035%7C9%2316839059_374_1610939555%7C9; BD_UPN=12314753; BAIDUID=240A7CFB87A518D8B44E180DA80D7195:FG=1; H_WISE_SIDS=107312_110085_127969_168389_176398_177371_177945_178384_178529_178616_179347_181106_181133_181135_181399_181589_181713_182000_182273_182529_182848_183030_183330_183536_183611_183983_184042_184246_184267_184320_184440_184560_184583_184794_184810_184838_184893_185029_185036_185141_185363_185517_185652_185747_186312_186318_186411_186596_186635_186649_186743_186840_187023_187040_187042_187067_187181_187195_187206_187287_187325_187380_187433_187450_187488_187495_187533_187678_187719_187726_187912_187928_187991_8000088_8000100_8000133_8000142_8000143_8000144_8000150_8000157_8000162_8000173_8000186; plus_lsv=e9e1d7eaf5c62da9; plus_cv=1::m:7.94e+147; Hm_lvt_12423ecbc0e2ca965d84259063d35238=1632391216; BAIDUID_BFESS=240A7CFB87A518D8B44E180DA80D7195:FG=1; BDRCVFR[k2U9xfnuVt6]=mk3SLVN4HKm; BD_HOME=1; H_PS_PSSID=26350; BA_HECTOR=a12l8h8k8005ag25a71gl5f2b0r",
		"Host":            "www.baidu.com",
		"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.61 Safari/537.36 Edg/94.0.992.31",
	}).Timeout(10 * time.Second).Get("https://www.baidu.com/")

	if err != nil {

		fmt.Println(err)

		return
	}

	c, err := rsp.Content()

	if err != nil {

		fmt.Println(err)

		return
	}

	fmt.Println(c.ToString())

}
