package request

import (
	"context"
	"errors"
	"github.com/PeterYangs/tools"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type request struct {
	request    *http.Response
	client     *Client
	params     map[string]interface{}
	query      map[string]interface{}
	method     string
	header     map[string]string
	retryTimes int
	timeout    time.Duration
	url        string
	body       io.Reader
}

func newRequest(c *Client) *request {

	r := &request{
		client:     c,
		header:     c.header,
		timeout:    c.timeout,
		retryTimes: c.retryTimes,
	}

	return r
}

// Params 设置参数(body设置后，Params会失效)
func (r *request) Params(p map[string]interface{}) *request {

	r.params = p

	return r

}

func (r *request) Query(p map[string]interface{}) *request {

	r.query = p

	return r
}

func (r *request) Header(header map[string]string) *request {

	r.header = header

	return r
}

// ReTry 重试次数
func (r *request) ReTry(times int) *request {

	r.retryTimes = times

	return r
}

func (r *request) Timeout(timeout time.Duration) *request {

	r.timeout = timeout

	return r

}

// Body body设置后，Params会失效
func (r *request) Body(b io.Reader) *request {

	r.body = b

	return r
}

func (r *request) Request(method string, url string) (*response, error) {

	req, err := r.dealRequest()

	if err != nil {

		return nil, err
	}

	rsp, err := r.do(req)

	return &response{response: rsp}, err

}

// Get get请求
func (r *request) Get(url string) (*response, error) {

	r.method = "GET"

	r.url = url

	return r.Request(r.method, r.url)

}

func (r *request) GetToContent(url string) (content, error) {

	r.method = "GET"

	r.url = url

	rsp, err := r.Request(r.method, r.url)

	if err != nil {

		return content{content: []byte{}}, err

	}

	b := r.toBody(rsp.response)

	return b.Content()

}

func (r *request) GetToContentWithHeader(url string) (content, http.Header, error) {

	r.method = "GET"

	r.url = url

	rsp, err := r.Request(r.method, r.url)

	if err != nil {

		return content{content: []byte{}}, map[string][]string{}, err

	}

	b := r.toBody(rsp.response)

	header := rsp.response.Header

	c, err := b.Content()

	return c, header, err
}

func (r *request) Post(url string) (*response, error) {

	r.method = "POST"

	r.url = url

	rsp, err := r.Request(r.method, r.url)

	return &response{response: rsp.response}, err

}

func (r *request) PostToContent(url string) (content, error) {

	r.method = "POST"

	r.url = url

	rsp, err := r.Request(r.method, r.url)

	if err != nil {

		return content{content: []byte{}}, err

	}

	b := r.toBody(rsp.response)

	return b.Content()
}

func (r *request) Put(url string) (*response, error) {

	r.method = "PUT"

	r.url = url

	rsp, err := r.Request(r.method, r.url)

	return &response{response: rsp.response}, err

}

func (r *request) PutToContent(url string) (content, error) {

	r.method = "PUT"

	r.url = url

	rsp, err := r.Request(r.method, r.url)

	if err != nil {

		return content{content: []byte{}}, err

	}

	b := r.toBody(rsp.response)

	return b.Content()

}

func (r *request) Delete(url string) (*response, error) {

	r.method = "DELETE"

	r.url = url

	rsp, err := r.Request(r.method, r.url)

	return &response{response: rsp.response}, err

}

func (r *request) DeleteToContent(url string) (content, error) {

	r.method = "DELETE"

	r.url = url

	rsp, err := r.Request(r.method, r.url)

	if err != nil {

		return content{content: []byte{}}, err

	}

	b := r.toBody(rsp.response)

	return b.Content()
}

// Download 下载文件
/**
url 文件链接
savePath 保存路径
*/
func (r *request) Download(url string, savePath string) error {

	r.method = "GET"

	r.url = url

	rsp, err := r.Request(r.method, r.url)

	if err != nil {

		return err
	}

	defer rsp.response.Body.Close()

	f, err := os.Create(savePath + ".temp")

	if err != nil {

		return err
	}

	defer func() {

		os.Remove(savePath + ".temp")

	}()

	_, err = io.Copy(f, rsp.response.Body)

	if err != nil {

		f.Close()

		return err
	}

	f.Close()

	if err = os.Rename(savePath+".temp", savePath); err != nil {

		return err
	}

	return nil
}

// Upload 文件上传
func (r *request) Upload(url string, filePath ...string) (content, error) {

	r.method = "POST"

	boundary := "285fa365bd76e6378f91f09f4eae20877246bbba4d31370d3c87b752d350" //可以自己设定，需要比较复杂的字符串作边界

	picData := ""

	for i, s := range filePath {

		var data []byte

		f, err := os.Open(s)

		if err != nil {

			return content{content: []byte{}}, err
		}

		data, err = ioutil.ReadAll(f)

		if err != nil {

			return content{content: []byte{}}, err
		}

		picData += "--" + boundary + "\n"
		picData += "Content-Disposition: form-data; name=\"f" + strconv.Itoa(i) + "\"; filename=" + f.Name() + "\n"
		picData += "Content-Type: application/octet-stream\n\n"
		picData += string(data) + "\n"
		picData += "--" + boundary + "\n"

	}

	req, err := http.NewRequest("POST", url, strings.NewReader(picData))

	req.Header.Set("Content-Type", "multipart/form-data; boundary="+boundary)

	rsp, err := r.do(req)

	if err != nil {

		return content{content: []byte{}}, err
	}

	defer rsp.Body.Close()

	bb, err := ioutil.ReadAll(rsp.Body)

	if err != nil {

		return content{content: []byte{}}, err
	}

	return content{content: bb}, err

}

//------------------------------------------------------------------------------------------------------------
//解析参数拼接参数字符串
func (r *request) resolveInterface(p map[string]interface{}, form string, parentName []string) string {

	if len(p) > 0 {

		for i, v := range p {

			switch key := v.(type) {

			case string:

				form += r.getKey(parentName, i) + "=" + key + "&"

			case int:

				form += r.getKey(parentName, i) + "=" + strconv.Itoa(key) + "&"

			case []string:

				for _, vv := range key {

					form += r.getKey(parentName, i) + "[]=" + vv + "&"

				}

			case []int:

				for _, vv := range key {

					form += r.getKey(parentName, i) + "[]=" + strconv.Itoa(vv) + "&"

				}

			case map[string]interface{}:

				t := append(parentName, i)

				form = r.resolveInterface(key, form, t)

			}

		}

	}

	return form

}

func (r *request) dealParamsAndQuery() (string, string) {

	query := ""

	params := ""

	if len(r.query) > 0 {

		query += "?"

		query = r.resolveInterface(r.query, query, []string{})

		if tools.SubStr(query, -1, -1) == "&" {

			query = tools.SubStr(query, 0, -2)
		}

	}

	if len(r.params) > 0 {

		params = r.resolveInterface(r.params, params, []string{})

		if tools.SubStr(params, -1, -1) == "&" {

			params = tools.SubStr(params, 0, -2)
		}

	}

	return query, params
}

func (r *request) getKey(parentName []string, ii string) string {

	f := ""

	for i, s := range parentName {

		if i == 0 {

			f += s

		} else {

			f += "[" + s + "]"

		}

	}

	if len(parentName) > 0 {

		f += "[" + ii + "]"

	} else {

		f += ii

	}

	return f

}

func (r *request) do(r2 *http.Request) (*http.Response, error) {

	for s, s2 := range r.header {

		r2.Header.Set(s, s2)

	}

	var err error

	var rsp *http.Response

	//错误重试
	for i := 0; i < r.retryTimes+1; i++ {

		rsp, err = r.work(r2)

		if err != nil {

			continue

		}

		if rsp.StatusCode != 200 {

			msg := ""

			if r.client.debug {

				b, _ := ioutil.ReadAll(rsp.Body)

				msg = "\n" + string(b)

			}

			rsp.Body.Close()

			err = errors.New("status code :" + strconv.Itoa(rsp.StatusCode) + "," + r2.URL.String() + msg)

			continue
		}

		return rsp, err

	}

	return rsp, err

}

func (r *request) work(r2 *http.Request) (*http.Response, error) {

	t := r.timeout
	//
	//默认超时时间
	if t == 0 {

		t = 30 * time.Second
	}

	cxt, _ := context.WithTimeout(context.Background(), t)

	r2 = r2.WithContext(cxt)

	return r.client.client.Do(r2)

}

func (r *request) toBody(rsp *http.Response) body {

	return body{
		body:   rsp.Body,
		header: rsp.Header,
	}

}

func (r *request) dealRequest() (*http.Request, error) {

	query, params := r.dealParamsAndQuery()

	var req *http.Request

	var err error

	if r.body != nil {

		req, err = http.NewRequest(r.method, r.url+query, r.body)

	} else {

		req, err = http.NewRequest(r.method, r.url+query, strings.NewReader(params))
	}

	if err != nil {

		return nil, err
	}

	switch r.method {

	case "POST":

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	case "PUT":

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	}

	return req, nil

}
