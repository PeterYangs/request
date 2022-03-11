package request

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type content struct {
	content  []byte
	response *response
}

func NewContent(rsp *response) (*content, error) {

	defer rsp.response.Body.Close()

	var err error

	var read io.ReadCloser

	read = rsp.response.Body

	if rsp.response.Header.Get("Content-Encoding") == "gzip" {

		read, err = gzip.NewReader(rsp.response.Body)

		if err != nil {

			return &content{
				content:  []byte{},
				response: rsp,
			}, err
		}

	}

	bb, err := ioutil.ReadAll(read)

	if err != nil {

		return &content{
			content:  []byte{},
			response: rsp,
		}, err
	}

	return &content{
		content:  bb,
		response: rsp,
	}, err

}

func (c *content) ToString() string {

	return string(c.content)
}

// ToJsonMap 转map
func (c *content) ToJsonMap() (map[string]interface{}, error) {

	var jsons map[string]interface{}

	err := json.Unmarshal(c.content, &jsons)

	if err != nil {

		return map[string]interface{}{}, err
	}

	return jsons, nil

}

// ToJsonStruct 转结构体
func (c *content) ToJsonStruct(st interface{}) error {

	err := json.Unmarshal(c.content, st)

	if err != nil {

		return err
	}

	return nil
}

func (c *content) Header() http.Header {

	return c.response.Header()
}

func (c *content) StatusCode() int {

	return c.response.response.StatusCode
}

// Proto http协议,如HTTP/1.1
func (c *content) Proto() string {

	return c.response.response.Proto
}

//// Time 获取响应时间
//func (c content) Time() time.Duration {
//
//	return c.request.responseTime
//}
