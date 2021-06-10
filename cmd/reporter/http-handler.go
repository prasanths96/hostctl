package reporter

import (
	"fmt"

	httpclient "github.com/ddliu/go-httpclient"
)

func post(url string, body []byte) {
	client := httpclient.Defaults(httpclient.Map{
		httpclient.OPT_USERAGENT: "hostctl http client",
		"Accept-Language":        "en-us",
	})

	resp, err := client.Post(url, string(body))
	if err != nil {
		panic(err.Error()) // Handle
	}

	outB, err := resp.ReadAll()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Final HTTP resp:", string(outB))

}
