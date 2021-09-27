package request

import (
	"github.com/PeterYangs/tools"
	"io/ioutil"
	"net/http"
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

func (c *Client) R() *request {

	return &request{client: c.client}
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

func (r *request) dealParams(p map[string]interface{}, form string, parentName []string) string {

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

	return r.client.Do(r2)

}

// Get get请求
func (r *request) Get(url string) (*response, error) {

	r.method = "GET"

	url = r.dealParams(r.params, url, []string{})

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {

		return nil, err
	}

	rsp, err := r.Do(req)

	return &response{response: rsp}, nil

}

func (r *request) GetToContent(url string) (content, error) {

	r.method = "GET"

	url = r.dealParams(r.params, url, []string{})

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

	p := r.dealParams(r.params, "", []string{})

	req, err := http.NewRequest("POST", url, strings.NewReader(p))

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {

		return nil, err
	}

	rsp, err := r.Do(req)

	return &response{response: rsp}, nil

}

func (r *request) PostToContent(url string) (content, error) {

	r.method = "POST"

	p := r.dealParams(r.params, "", []string{})

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
