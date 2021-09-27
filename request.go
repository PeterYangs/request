package request

import (
	"errors"
	"github.com/PeterYangs/tools"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type request struct {
	request *http.Response
	client  *http.Client
	params  map[string]interface{}
	method  string
	header  map[string]string
}

// Params 设置参数
func (r *request) Params(p map[string]interface{}) *request {

	r.params = p

	return r

}

func (r *request) Header(header map[string]string) *request {

	r.header = header

	return r
}

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

func (r *request) dealParams(form string) string {

	if len(r.params) > 0 {

		if r.method == "GET" {

			form += "?"
		}

		form = r.resolveInterface(r.params, form, []string{})

		if tools.SubStr(form, -1, -1) == "&" {

			form = tools.SubStr(form, 0, -2)
		}

		return form

	}

	return form
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

func (r *request) Do(r2 *http.Request) (*http.Response, error) {

	for s, s2 := range r.header {

		r2.Header.Set(s, s2)

	}

	rsp, err := r.client.Do(r2)

	if err != nil {

		return rsp, err
	}

	if rsp.StatusCode != 200 {

		rsp.Body.Close()

		return rsp, errors.New("status code :" + strconv.Itoa(rsp.StatusCode))
	}

	return rsp, err

}

// Get get请求
func (r *request) Get(url string) (*response, error) {

	r.method = "GET"

	url = r.dealParams(url)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {

		return nil, err
	}

	rsp, err := r.Do(req)

	return &response{response: rsp}, err

}

func (r *request) GetToContent(url string) (content, error) {

	r.method = "GET"

	url = r.dealParams(url)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {

		return content{content: []byte{}}, err
	}

	rsp, err := r.Do(req)

	if err != nil {

		return content{content: []byte{}}, err

	}

	defer rsp.Body.Close()

	bb, err := ioutil.ReadAll(rsp.Body)

	if err != nil {

		return content{content: []byte{}}, err

	}

	return content{content: bb}, nil

}

func (r *request) Post(url string) (*response, error) {

	r.method = "POST"

	p := r.dealParams("")

	req, err := http.NewRequest("POST", url, strings.NewReader(p))

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {

		return nil, err
	}

	rsp, err := r.Do(req)

	return &response{response: rsp}, err

}

func (r *request) PostToContent(url string) (content, error) {

	r.method = "POST"

	p := r.dealParams("")

	req, err := http.NewRequest("POST", url, strings.NewReader(p))

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {

		return content{content: []byte{}}, err
	}

	rsp, err := r.Do(req)

	if err != nil {

		return content{content: []byte{}}, err

	}

	defer rsp.Body.Close()

	bb, err := ioutil.ReadAll(rsp.Body)

	if err != nil {

		return content{content: []byte{}}, err

	}

	return content{content: bb}, nil
}

// Download 下载文件
/**
url 文件链接
savePath 保存路径
*/
func (r *request) Download(url string, savePath string) error {

	r.method = "GET"

	url = r.dealParams(url)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {

		return err
	}

	rsp, err := r.Do(req)

	if err != nil {

		return err
	}

	defer rsp.Body.Close()

	f, err := os.Create(savePath + ".temp")

	if err != nil {

		return err
	}

	defer func() {

		os.Remove(savePath + ".temp")

	}()

	_, err = io.Copy(f, rsp.Body)

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
